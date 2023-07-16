#include "ordering/poc/pow/block_manager.h"

#include <glog/logging.h>

#include "common/utils/utils.h"
#include "crypto/signature_verifier.h"
#include "ordering/poc/pow/merkle.h"
#include "proto/resdb.pb.h"
#include "ordering/poc/pow/miner_utils.h"
#include "statistic/stats.h"

namespace resdb {

BlockManager::BlockManager(const XDBPoCConfig& config,
    TransactionExecutorImpl* executor) : config_(config), executor_(executor){
  miner_ = std::make_unique<Miner>(config);
  global_stats_ = Stats::GetGlobalStats();
  last_update_time_ = 0;
  //prometheus_handler_ = Stats::GetGlobalPrometheus();
}

void BlockManager::SaveClientTransactions(std::unique_ptr<BatchClientTransactions> client_request){
  for(const ClientTransactions& client_tx : client_request->transactions()){
    *request_candidate_.add_transactions() = client_tx;
  }
  if(request_candidate_.min_seq() == 0){
	  request_candidate_.set_min_seq(client_request->min_seq());
  }
  else {
	  request_candidate_.set_min_seq(std::min(request_candidate_.min_seq(), client_request->min_seq()));
  }
  request_candidate_.set_max_seq(std::max(request_candidate_.max_seq(), client_request->max_seq()));
}

int BlockManager::SetNewMiningBlock(
    std::unique_ptr<BatchClientTransactions> client_request) {
SaveClientTransactions(std::move(client_request));

  std::unique_ptr<Block> new_block = std::make_unique<Block>();
  if (request_candidate_.min_seq() != GetLastSeq() + 1) {
    LOG(ERROR) << "seq invalid:" << request_candidate_.min_seq()
               << " last seq:" << GetLastSeq();
    return -2;
  }

  int64_t new_time = 0;
  for(ClientTransactions& client_tx : *request_candidate_.mutable_transactions()){
	  new_time = client_tx.create_time();
	  if(new_time>0){
		  create_time_[client_tx.seq()] = new_time;
		  client_tx.clear_create_time();
	  }
  }

  new_block->mutable_header()->set_height(GetCurrentHeight() + 1);
  *new_block->mutable_header()->mutable_pre_hash() =
      GetPreviousBlcokHash();  // set the hash value of the parent block.
  request_candidate_.SerializeToString(new_block->mutable_transaction_data());
  *new_block->mutable_header()->mutable_merkle_hash() =
      Merkle::MakeHash(request_candidate_);
  new_block->mutable_header()->set_min_seq(request_candidate_.min_seq());
  new_block->mutable_header()->set_max_seq(request_candidate_.max_seq());
  new_block->set_miner(config_.GetSelfInfo().id());
  new_block->set_block_time(GetCurrentTime());

  LOG(ERROR) << "create new block:" << request_candidate_.transactions(0).create_time()<<"["<<new_block->header().min_seq()<<","<<new_block->header().max_seq()<<"]" << " miner:"<<config_.GetSelfInfo().id()<<" time:"<<new_block->block_time()<<" delay:"<<(new_block->block_time() - new_time)/1000000.0 << " current:"<<GetCurrentTime();
  new_mining_block_ = std::move(new_block);
  miner_->SetSliceIdx(0);
  return 0;
}

Block* BlockManager::GetNewMiningBlock() {
  return new_mining_block_ == nullptr ? nullptr : new_mining_block_.get();
}

// Mine the nonce.
absl::Status BlockManager::Mine() {
  LOG(ERROR)<<"mine??";
  if (new_mining_block_ == nullptr) {
    LOG(ERROR) << "don't contain mining block.";
    return absl::InvalidArgumentError("height invalid");
  }

  if (new_mining_block_->header().height() != GetCurrentHeight() + 1) {
    // a new block has been committed.
    LOG(ERROR) << "new block height:" << new_mining_block_->header().height()
               << " current height:" << GetCurrentHeight() << " not equal";
    return absl::InvalidArgumentError("height invalid");
  }

  return miner_->Mine(new_mining_block_.get());
}

int32_t BlockManager::GetSliceIdx() { return miner_->GetSliceIdx(); }

int BlockManager::SetSliceIdx(const SliceInfo& slice_info) {
  if (new_mining_block_ == nullptr) {
    return 0;
  }
  if (slice_info.height() != new_mining_block_->header().height()) {
    return -2;
  }
  size_t f = config_.GetMaxMaliciousReplicaNum();
  if(slice_info.shift_idx() > f ){
	  // reset
	  LOG(ERROR)<<"reset slice";
	miner_->SetSliceIdx(0);
	  return 1;
  }
  miner_->SetSliceIdx(slice_info.shift_idx());
  return 0;
}

int BlockManager::Commit() {
	  request_candidate_.Clear();
  uint64_t mining_time = GetCurrentTime() - new_mining_block_->block_time();
  new_mining_block_->set_mining_time(mining_time);
  int ret = AddNewBlock(std::move(new_mining_block_));
  new_mining_block_ = nullptr;
  return ret;
}

// =============== Mining Related End ==========================
int BlockManager::Commit(std::unique_ptr<Block> new_block) {
  return AddNewBlock(std::move(new_block));
}

int BlockManager::AddNewBlock(std::unique_ptr<Block> new_block) {
  {
    std::unique_lock<std::mutex> lck(mtx_);
    if (new_block->header().height() != GetCurrentHeightNoLock() + 1) {
      // a new block has been committed.
      LOG(ERROR) << "new block height:" << new_block->header().height()
                 << " current height:" << GetCurrentHeightNoLock()
                 << " not equal";
      return -2;
    }
    LOG(INFO) << "============= commit new block:"
              << new_block->header().height()
              << " current height:" << GetCurrentHeightNoLock();
LOG(ERROR)<<"commit:"<<new_block->header().height()<<" from:"<<new_block->miner()<<" mining time:"<<new_block->mining_time()/1000000.0;
    //prometheus_handler_->SetValue(MINING_TIME, new_block->mining_time()/1000000.0);
    //prometheus_handler_->Inc(MINING_LATENCY, new_block->mining_time()/1000000000.0);
    miner_->Terminate();
    request_candidate_.Clear();

    block_list_.push_back(std::move(new_block));
    Execute(*block_list_.back());
  }
  return 0;
}

void BlockManager::Execute(const Block& block){

BatchClientTransactions batch_client_request;
  if(!batch_client_request.ParseFromString(block.transaction_data())){
	  LOG(ERROR)<<"parse client transaction fail";
  }

  LOG(ERROR)<<" execute seq:["<<batch_client_request.min_seq()<<","<<batch_client_request.max_seq()<<"]";	
  for(const ClientTransactions& client_tx : batch_client_request.transactions()){
	  BatchClientRequest batch_request;
	  if (!batch_request.ParseFromString(client_tx.transaction_data())) {
		  LOG(ERROR) << "parse data fail";
	  }
	  global_stats_->IncTotalRequest(batch_request.client_requests_size()); 
    if(executor_){
      executor_->ExecuteBatch(batch_request);
    }
  }
}

bool BlockManager::VerifyBlock(const BlockMiningInfo &block_info) {
	if(new_mining_block_ == nullptr){
		LOG(ERROR)<<"no mining block";
		return false;
	}

	if(new_mining_block_->header().min_seq() != block_info.header().min_seq()){
		return false;
	}
	if(new_mining_block_->header().max_seq() != block_info.header().max_seq()){
		return false;
	}
	Block new_block;
	*new_block.mutable_header() = block_info.header();
	new_block.set_transaction_data(new_mining_block_->transaction_data());
	*new_block.mutable_hash() = block_info.hash();
	return VerifyBlock(&new_block);
}

bool BlockManager::VerifyBlock(const Block* block) {
  if (!miner_->IsValidHash(block)) {
    LOG(ERROR) << "hash not valid:" << block->hash().DebugString();
    return false;
  }

  if(block->transaction_data().empty()){
	  	LOG(ERROR)<<"no data";
		return false;
  }

  BatchClientTransactions client_request;
  if (!client_request.ParseFromString(block->transaction_data())) {
    LOG(ERROR) << "parse transaction fail";
    return false;
  }
  return Merkle::MakeHash(client_request) == block->header().merkle_hash();
}

uint64_t BlockManager::GetCurrentHeight() {
  std::unique_lock<std::mutex> lck(mtx_);
  return GetCurrentHeightNoLock();
}

uint64_t BlockManager::GetLastSeq() {
  std::unique_lock<std::mutex> lck(mtx_);
  return block_list_.empty() ? 0 : block_list_.back()->header().max_seq();
}

uint64_t BlockManager::GetLastCandidateSeq() {
  return request_candidate_.max_seq();
}

uint64_t BlockManager::GetCurrentHeightNoLock() {
  return block_list_.empty() ? 0 : block_list_.back()->header().height();
}

HashValue BlockManager::GetPreviousBlcokHash() {
  std::unique_lock<std::mutex> lck(mtx_);
  return block_list_.empty() ? HashValue() : block_list_.back()->hash();
}


Block* BlockManager::GetBlockByHeight(uint64_t height) {
  std::unique_lock<std::mutex> lck(mtx_);
  if (block_list_.empty() || block_list_.back()->header().height() < height) {
    return nullptr;
  }
  return block_list_[height - 1].get();
}

Block * BlockManager::GetCurrentBlock() {
	if(new_mining_block_ == nullptr) return nullptr;
	return new_mining_block_.get();
}

void BlockManager::SetTargetValue(const HashValue& target_value) {
  miner_->SetTargetValue(target_value);
}

std::unique_ptr<BlockMiningInfo> BlockManager::GetPendingBlockInfo() {
	if(new_mining_block_== nullptr){
		LOG(ERROR)<<"no mining block";
		return nullptr;
	}
	LOG(ERROR)<<"????";
	std::unique_ptr<BlockMiningInfo> block_info = std::make_unique<BlockMiningInfo>();
	*block_info->mutable_header() = new_mining_block_->header();
	block_info->set_miner(config_.GetSelfInfo().id());
	*block_info->mutable_hash() = new_mining_block_->hash();
	return block_info;
}

int BlockManager::Confirm(const BlockMiningInfo& block_info) {
  if(new_mining_block_ == nullptr){
    LOG(ERROR)<<"block not exist";
    return -1;
  }

  if(new_mining_block_->header().min_seq() != block_info.header().min_seq()){
    LOG(ERROR)<<"block not exist";
    return -1;
  }
  
LOG(ERROR)<<"confirm block:"<<block_info.DebugString()<<" data size:"<<new_mining_block_->transaction_data().size();
	std::unique_ptr<Block> block = std::make_unique<Block>();
	*block->mutable_header() = block_info.header();
	block->set_transaction_data(new_mining_block_->transaction_data());
  *block->mutable_hash() = block_info.hash();

	if(!VerifyBlock(block.get())){
		LOG(ERROR)<<"varify block fail";
		return -2;
	}
	int ret = Commit(std::move(block));
  if(ret==0){
    new_mining_block_ = nullptr;
  }
  return ret;
}

}  // namespace resdb
