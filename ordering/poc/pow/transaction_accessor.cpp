#include "ordering/poc/pow/transaction_accessor.h"

#include <glog/logging.h>
#include "common/utils/utils.h"

namespace XXXX {

TransactionAccessor::TransactionAccessor(const XDBPoCConfig& config,
                                         bool auto_start)
    : config_(config) {
  stop_ = false;
  max_received_seq_ = 0;
  next_consume_ = 1;
  if (auto_start) {
    fetching_thread_ =
        std::thread(&TransactionAccessor::TransactionFetching, this);
  }
  
  std::vector<ReplicaInfo> replicas = config.GetBFTConfig()->GetReplicaInfos();
  replica_client_ = std::make_unique<XDBReplicaClient>(replicas);
	//txn_client_ = GetXDBTxnClient();

  //prometheus_handler_ = Stats::GetGlobalPrometheus();
}

TransactionAccessor::~TransactionAccessor() {
  stop_ = true;
  if (fetching_thread_.joinable()) {
    fetching_thread_.join();
  }
}

void TransactionAccessor::Start() {
  fetching_thread_ =
      std::thread(&TransactionAccessor::TransactionFetching, this);
}

void TransactionAccessor::TransactionFetching() {
  std::unique_ptr<XDBTxnClient> client = GetXDBTxnClient();
  assert(client != nullptr);

  while (!stop_) {
    if( max_received_seq_ > next_consume_ && max_received_seq_ - next_consume_ > 2 * config_.BatchTransactionNum() ){
      LOG(ERROR)<<"max received :"<<max_received_seq_<<" next consume:"<<next_consume_<<" skip";
      sleep(1);
      continue;
    }
    uint64_t cur_seq = max_received_seq_ + 1;
    TxnQueryRequest request;
    request.set_min_seq(cur_seq);
    request.set_max_seq(cur_seq + 1000);

    std::string str;
    request.SerializeToString(&str);

    auto ret = client->GetCustomQuery(str);
    if (!ret.ok()){
      LOG(ERROR) << "get txn fail:" << cur_seq;
      sleep(1);
      continue;
    }

    TxnQueryResponse response;
    if(!response.ParseFromString(*ret)){
      LOG(ERROR)<<"parse from response fail";
      continue;
    }

    if(response.data().empty()){
      sleep(1);
      continue;
    }

    for(int i = 0; i < response.data_size(); ++i){
      std::unique_ptr<ClientTransactions> client_txn =
        std::make_unique<ClientTransactions>();
      client_txn->set_transaction_data(response.data(i));
      client_txn->set_seq(response.seq(i));
      client_txn->set_create_time(GetCurrentTime());
      queue_.Push(std::move(client_txn));
      cur_seq = std::max(cur_seq, response.seq(i));
      /*
       BatchClientRequest req;
       if(!req.ParseFromString(response.data(i))){
        LOG(ERROR)<<"parse fail";
       }
       LOG(ERROR)<<"get req txn:"<<req.DebugString();
       */

      //LOG(ERROR)<<"obtain txn seq:"<<response.seq(i);
    }

    max_received_seq_ = cur_seq;

    std::lock_guard<std::mutex> lk(mutex_);
    cv_.notify_all();
  }
  return;
}

std::unique_ptr<XDBTxnClient> TransactionAccessor::GetXDBTxnClient() {
  return std::make_unique<XDBTxnClient>(*config_.GetBFTConfig());
}

// obtain [seq, seq+batch_num-1] transactions
std::unique_ptr<BatchClientTransactions>
TransactionAccessor::ConsumeTransactions(uint64_t seq) {
  LOG(ERROR) << "consume transaction:" << seq
             << " batch:" << config_.BatchTransactionNum()
             << " received max seq:" << max_received_seq_;
  if (seq + config_.BatchTransactionNum() > max_received_seq_ + 1) {
	  std::unique_lock<std::mutex> lk(mutex_);
	  cv_.wait_for(lk, std::chrono::seconds(1),
			  [&] { return seq + config_.BatchTransactionNum() <= max_received_seq_ + 1; });

    return nullptr;
  }
  /*
  while(seq > next_consume_){
    queue_.Pop();
    next_consume_++;
  }
  if (seq != next_consume_) {
    LOG(ERROR) << "next should consume:" << next_consume_;
    return nullptr;
  }
  */

  std::unique_ptr<BatchClientTransactions> batch_transactions =
      std::make_unique<BatchClientTransactions>();
  for (uint32_t i = 0; i < config_.BatchTransactionNum(); ++i) {
    auto ptr = queue_.Pop();
    if(ptr == nullptr){
      LOG(ERROR)<<"get null";
      i--;
      continue;
    }
    int64_t seq = ptr->seq();
    //LOG(ERROR)<<"get seq:"<<seq<<" next consume:"<<next_consume_;
    if(seq < next_consume_){
      i--;
      continue;
    }
    batch_transactions->set_max_seq(seq);
    *batch_transactions->add_transactions() = *ptr;
  }
  batch_transactions->set_min_seq(seq);
  //batch_transactions->set_max_seq(seq + config_.BatchTransactionNum() - 1);
  next_consume_ = batch_transactions->max_seq()+1;
  LOG(ERROR)<<"get batch from "<<batch_transactions->max_seq()<<" to "<<batch_transactions->min_seq();
  return batch_transactions;
}
			
void TransactionAccessor::SendMiningResult(const BlockMiningInfo& mining_result) {
	Request request;
	request.set_type(Request::TYPE_GEO_MINING_RESULT);
	request.mutable_region_info()->set_region_id(
			config_.GetConfigData().self_region_id());
	mining_result.SerializeToString(request.mutable_data());
	replica_client_->SendMessage(request);
}

absl::StatusOr<BlockMiningInfo> TransactionAccessor::FetchingResult(uint64_t seq) {
	auto txn_client = GetXDBTxnClient();
  LOG(ERROR)<<"fetch result:"<<seq;

  TxnQueryRequest request;
  request.set_min_seq(seq);
  request.set_max_seq(seq);
  request.set_is_query_results(true);

	std::string str;
	request.SerializeToString(&str);
	auto ret = txn_client->GetCustomQuery(str);
	if(ret.ok()){
    TxnQueryResponse response;
    if(!response.ParseFromString(*ret)){
      LOG(ERROR)<<"parse from response fail";
      return absl::InternalError("recv data fail.");
    }
    if(response.data_size() == 0){
      return absl::InternalError("recv data fail.");
    }
		BlockMiningInfo resp;
    if(resp.ParseFromString(response.data(0))){
      LOG(ERROR)<<"obtain mining info:"<<resp.DebugString();
      return resp;
    }
    return absl::InternalError("recv data fail.");
	}
  else {
    LOG(ERROR)<<"get custom query fail:";
  }
	return ret.status();
}

}  // namespace XXXX
