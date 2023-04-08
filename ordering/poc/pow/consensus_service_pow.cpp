#include "ordering/poc/pow/consensus_service_pow.h"

#include "common/utils/utils.h"
#include "glog/logging.h"

namespace resdb {
using poc::PoCTransactionManager;

ConsensusServicePoW::ConsensusServicePoW(const ResDBPoCConfig& config, PoCTransactionManager * manager)
    : ConsensusService(config), txn_manager_(manager) {
  miner_manager_ = std::make_unique<MinerManager>(config);
  pow_manager_ = std::make_unique<PoWManager>(config,GetBroadCastClient(), txn_manager_);
}

void ConsensusServicePoW::Start() {
  ConsensusService::Start();
  pow_manager_->Start();
}

ConsensusServicePoW::~ConsensusServicePoW() {
}

std::vector<ReplicaInfo> ConsensusServicePoW::GetReplicas() {
  return miner_manager_->GetReplicas();
}

int ConsensusServicePoW::ConsensusCommit(std::unique_ptr<Context> context,
                                         std::unique_ptr<Request> request) {
  LOG(ERROR) << "recv impl type:" << request->type() << " "
            << request->client_info().DebugString()
           << "sender id:" << request->sender_id();
  switch (request->type()) {
    case PoWRequest::TYPE_COMMITTED_BLOCK: {
      std::unique_ptr<Block> block = std::make_unique<Block>();
      if (block->ParseFromString(request->data())) {
        pow_manager_->Commit(std::move(block));
      }
      break;
    }
    case PoWRequest::TYPE_SHIFT_MSG: {
      std::unique_ptr<SliceInfo> slice_info = std::make_unique<SliceInfo>();
      if (slice_info->ParseFromString(request->data())) {
        pow_manager_->AddShiftMsg(*slice_info);
      }
      else {
	      LOG(ERROR)<<"parse info fail";
      }
      break;
    }
    case PoWRequest::TYPE_MINING_RESULTS: {
      std::unique_ptr<BlockMiningInfo> info = std::make_unique<BlockMiningInfo>();
      if (info->ParseFromString(request->data())) {
        pow_manager_->ReceiveMiningResult(std::move(info));
      }
      else {
	      LOG(ERROR)<<"parse info fail";
      }
      break;
    }
    case Request::TYPE_CLIENT_REQUEST:
      return ClientQuery(std::move(context), std::move(request));
  }

  return 0;
}

int ConsensusServicePoW::ClientQuery(std::unique_ptr<Context> context,
                                         std::unique_ptr<Request> request) {
      std::unique_ptr<std::string> resp = txn_manager_->ClientQuery(request->data());
      return context->client->SendRawMessageData(*resp);
}

}  // namespace resdb

