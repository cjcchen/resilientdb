#include "ordering/poc/pow/pow_manager.h"

#include "common/utils/utils.h"
#include "glog/logging.h"

namespace resdb {
namespace {

std::unique_ptr<Request> NewRequest(PoWRequest type,
                                    const ::google::protobuf::Message& message,
                                    int32_t sender) {
  auto new_request = std::make_unique<Request>();
  new_request->set_type(type);
  new_request->set_sender_id(sender);
  message.SerializeToString(new_request->mutable_data());
  return new_request;
}

}  // namespace


PoWManager::PoWManager(const ResDBPoCConfig& config,
		ResDBReplicaClient* client
		) : config_(config),bc_client_(client){
Reset();
self_id_ = config_.GetSelfInfo().id();
  is_stop_ = false;
  last_done_seq_ = 0;
  //prometheus_handler_ = Stats::GetGlobalPrometheus();
}

PoWManager::~PoWManager() {
  Stop();
  if (miner_thread_.joinable()) {
    miner_thread_.join();
  }
  if (fetch_thread_.joinable()) {
    fetch_thread_.join();
  }
}

std::unique_ptr<TransactionAccessor> PoWManager::GetTransactionAccessor(const ResDBPoCConfig& config){
  return std::make_unique<TransactionAccessor>(config);
}

std::unique_ptr<ShiftManager> PoWManager::GetShiftManager(const ResDBPoCConfig& config){
  return std::make_unique<ShiftManager>(config);
}

std::unique_ptr<BlockManager> PoWManager::GetBlockManager(const ResDBPoCConfig& config){
  return std::make_unique<BlockManager>(config);
}

void PoWManager::Commit(std::unique_ptr<Block> block){
       std::unique_lock<std::mutex> lck(tx_mutex_);
	if(block_manager_->Commit(std::move(block))==0){
		LOG(ERROR)<<"commit block succ";
		NotifyNextBlock();
	}
	LOG(ERROR)<<"commit block done";
}

void PoWManager::NotifyNextBlock(){
	std::unique_lock<std::mutex> lk(mutex_);
	cv_.notify_one();
	if(current_status_ == GENERATE_NEW){
		current_status_ = NEXT_NEWBLOCK;
	}
	LOG(ERROR)<<"notify block:"<<current_status_;
}

PoWManager::BlockStatus PoWManager::GetBlockStatus() {
	return current_status_;
}

absl::Status PoWManager::WaitBlockDone(){
	LOG(ERROR)<<"wait block:"<<current_status_;
	//int timeout_ms = config_.GetMiningTime();
	//60000000;
	int timeout_ms = 300000000;
	std::unique_lock<std::mutex> lk(mutex_);
	cv_.wait_for(lk, std::chrono::microseconds(timeout_ms), [&] {
			return current_status_ == NEXT_NEWBLOCK;
			});
	LOG(ERROR)<<"wait block done:" << current_status_<<" next block?:"<<(current_status_ == NEXT_NEWBLOCK);
	if (current_status_ == NEXT_NEWBLOCK){
		return absl::OkStatus();
	}
	return absl::NotFoundError("No new transaction.");
}

void PoWManager::Reset(){
  transaction_accessor_ = GetTransactionAccessor(config_);
  shift_manager_ = GetShiftManager(config_);
  block_manager_ = GetBlockManager(config_);
  LOG(ERROR)<<"reset:"<<transaction_accessor_.get();
}

void PoWManager::Start() {
  miner_thread_ = std::thread(&PoWManager::MiningProcess, this);
  fetch_thread_ = std::thread(&PoWManager::MiningResultsProcess, this);
  verify_thread_ = std::thread(&PoWManager::ResultsVerifyProcess, this);
  is_stop_ = false;
}

void PoWManager::Stop() {
	is_stop_ = true;
}

bool PoWManager::IsRunning() {
	return !is_stop_;
}

void PoWManager::AddShiftMsg(const SliceInfo& slice_info) {
	shift_manager_->AddSliceInfo(slice_info);
}
        
void PoWManager::ReceiveMiningResult(std::unique_ptr<BlockMiningInfo> info) {
	LOG(ERROR)<<"receive mining result from:"<<info->miner()<<" seq:"<<info->header().min_seq();
       std::unique_lock<std::mutex> lck(block_result_mutex_);
	result_candidate_.insert(std::make_pair(info->header().min_seq(), std::move(info)));
}
	  
uint64_t PoWManager::GetFirstResultSeq(){
       std::unique_lock<std::mutex> lck(block_result_mutex_);
       if(result_candidate_.empty()){
	       return 0;
       }
	return result_candidate_.begin()->first;
}

std::unique_ptr<BlockMiningInfo> PoWManager::GetResultInfo(uint64_t seq) {
	std::unique_lock<std::mutex> lck(block_result_mutex_);
       auto it = result_candidate_.find(seq);
       if(it == result_candidate_.end()){
	       return nullptr;
       }
       auto ret = std::move(it->second);
       result_candidate_.erase(it);
       return ret;
}

void PoWManager::RemoveResult(uint64_t seq){
       std::unique_lock<std::mutex> lck(block_result_mutex_);
       if(result_candidate_.empty()){
	       return ;
       }
       auto it = result_candidate_.find(seq);
       if(it == result_candidate_.end()){
	       return;
       }
       result_candidate_.erase(it);
}
	  

bool PoWManager::IsSent(uint64_t seq){
	if(seq <= last_done_seq_){
		return true;
	}
	return sent_list_.find(seq) != sent_list_.end();
}

void PoWManager::SetSent(uint64_t seq){
	sent_list_.insert(seq);
}


int PoWManager::GetShiftMsg(const SliceInfo& slice_info) {
	LOG(ERROR)<<"check shift msg:"<<slice_info.DebugString();
	if(!shift_manager_->Check(slice_info)){
	  return -1;
	}
	LOG(ERROR)<<"slice info is ok:"<<slice_info.DebugString();
        if(block_manager_->SetSliceIdx(slice_info)==1){
		return 1;
	}
	return 0;
}

int PoWManager::GetMiningTxn(MiningType type){
       std::unique_lock<std::mutex> lck(tx_mutex_);
       //LOG(ERROR)<<"get mining txn status:"<<current_status_;
	if(current_status_ == NEXT_NEWBLOCK){
		type = MiningType::NEWBLOCK;
	}
	if(type == NEWBLOCK){
		uint64_t max_seq = std::max(block_manager_->GetLastSeq(), block_manager_->GetLastCandidateSeq());
		//LOG(ERROR)<<"get block last max:"<<block_manager_->GetLastSeq()<<" "<<block_manager_->GetLastCandidateSeq();
		auto client_tx = transaction_accessor_->ConsumeTransactions(max_seq+1);
		if(client_tx == nullptr){
		  return -2;
		}
		block_manager_->SetNewMiningBlock(std::move(client_tx));
	}
	else {
    		//prometheus_handler_->Inc(SHIFT_MSG, 1);
		int ret = GetShiftMsg(need_slice_info_);
		//LOG(ERROR)<<"get shift msg ret:"<<ret;
		if(ret==1){
			// no solution after enought shift.
			return 1;
		}
		if (ret !=0){
		  BroadCastShiftMsg(need_slice_info_);
		  return -2;
		}
		LOG(ERROR)<<"get shift msg:"<<need_slice_info_.DebugString();
	}
	return 0;
}

PoWManager::MiningStatus PoWManager::Wait(){
	current_status_ = GENERATE_NEW;
	LOG(ERROR)<<"wait mining";
	auto mining_thread = std::thread([&](){
		LOG(ERROR)<<"start to mine";
		absl::Status status = block_manager_->Mine();
		if(status.ok()){
			LOG(ERROR)<<"mine done get info";
			std::unique_ptr<BlockMiningInfo> info= block_manager_->GetPendingBlockInfo();
			if(info !=nullptr){
				LOG(ERROR)<<"send to pbft:"<<info->header().min_seq();
				transaction_accessor_->SendMiningResult(*info);
				BroadCastMiningResult(*info);
			}
			LOG(ERROR)<<"done mine";
		}
		LOG(ERROR)<<"mine:"<<status.ok();
	});

	auto status = WaitBlockDone();
	if(mining_thread.joinable()){
		mining_thread.join();
	}
	LOG(ERROR)<<"success:"<<status.ok()<<" status:"<<current_status_;
	if(status.ok()||current_status_ == NEXT_NEWBLOCK){
		return MiningStatus::OK;
	}
	return MiningStatus::TIMEOUT;
}

// receive a block before send and after need send
void PoWManager::SendShiftMsg(){
	LOG(ERROR)<<"send shift";
	SliceInfo slice_info;
	slice_info.set_shift_idx(block_manager_->GetSliceIdx()+1);
	slice_info.set_height(
			block_manager_->GetNewMiningBlock()->header().height());
	slice_info.set_sender(config_.GetSelfInfo().id());
	BroadCastShiftMsg(slice_info);
	
	need_slice_info_ = slice_info;
	LOG(ERROR)<<"send shift msg";
}

int PoWManager::BroadCastShiftMsg(const SliceInfo& slice_info) {
  auto request = NewRequest(PoWRequest::TYPE_SHIFT_MSG, slice_info,
                            config_.GetSelfInfo().id());
  bc_client_->BroadCast(*request);
  return 0;
}

int PoWManager::BroadCastMiningResult(const BlockMiningInfo& info) {
  auto request = NewRequest(PoWRequest::TYPE_MINING_RESULTS, info,
                            config_.GetSelfInfo().id());
  LOG(ERROR)<<"bc result:"<<info.header().min_seq();
  bc_client_->BroadCast(*request);
  return 0;
}

// Broadcast the new block if once it is committed.
void PoWManager::MiningResultsProcess() {
  uint64_t current_seq = 0;
  while (IsRunning()) {
	  uint64_t last_seq = block_manager_->GetLastSeq();
    if(last_seq+1 <= current_seq){
      sleep(1);
      continue;
    }
	  while(IsRunning()){
		  absl::StatusOr<BlockMiningInfo> resp = transaction_accessor_->FetchingResult(last_seq+1);
      if(resp.ok()){
        while(IsRunning()){
          int ret = block_manager_->Confirm(*resp);
          if(ret == 0) {
            NotifyNextBlock();
            current_seq = last_seq+1;
            break;
          }
          else if(ret==-1){
            LOG(ERROR)<<"block not ready";
            sleep(1);
            continue;
          }
          else {
            break;
          }
        }
        break;
		  }
		  else {
			  //LOG(ERROR)<<"no result, continue query";
			  sleep(1);
		  }
	  }

  }
}

void PoWManager::NotifyBroadCast() {
  std::unique_lock<std::mutex> lck(broad_cast_mtx_);
  broad_cast_cv_.notify_all();
}

void PoWManager::ResultsVerifyProcess() {
  while (IsRunning()) {
	  uint64_t seq = GetFirstResultSeq();
	  if(seq == 0){
		  sleep(1);
		  continue;
	  }
	  uint64_t current_seq = block_manager_->GetLastSeq();
	  if(seq <= current_seq){
		  RemoveResult(seq);
		  LOG(ERROR)<<"get result seq invalid:"<<seq;
		  continue;
	  }
	  Block * block = block_manager_->GetCurrentBlock();
	  if(block == nullptr){
		  LOG(ERROR)<<"no mining block";
		  sleep(1);
		  continue;
	  }
	  if(block->header().min_seq() != seq){
		  LOG(ERROR)<<"result not equal to the mining block:"<<seq<<" block:"<<block->header().min_seq();
		  sleep(1);
		  continue;
	  }
	  std::unique_ptr<BlockMiningInfo> block_info = GetResultInfo(seq);
	  LOG(ERROR)<<"block info miner:"<<block_info->miner();
	  if(block_info == nullptr){
		  continue;
	  }

	  if(block_info->miner() == self_id_){
		  sleep(1);
		  continue;
	  }
	  
	  if(!block_manager_->VerifyBlock(*block_info)){
		  LOG(ERROR)<<"varify block fail";
		  continue;
	  }

	  LOG(ERROR)<<"send to pbft";
	  transaction_accessor_->SendMiningResult(*block_info);
  }
}


// Mining the new blocks got from PBFT cluster.
void PoWManager::MiningProcess() {
	if(config_.GetSelfInfo().id() > 8 && config_.GetSelfInfo().id() < 13){
		//return;
	}
	LOG(ERROR)<<"start";
	MiningType type = MiningType::NEWBLOCK;
	while (IsRunning()) {
		int ret = GetMiningTxn(type);
		if(ret<0){
		      usleep(10000);
		continue;
		}
		else if(ret>0){
			type = MiningType::NEWBLOCK;
			continue;
		}

		auto mining_status = Wait();
		if(mining_status == MiningStatus::TIMEOUT){
			type = MiningType::SHIFTING;
			SendShiftMsg();
		}
		else {
			type = MiningType::NEWBLOCK;
			//LOG(ERROR)<<"done";
		      	//NotifyBroadCast();
			//SendResults();
		}
	}
}

}
