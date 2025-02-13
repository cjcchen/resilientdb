#pragma once

#include "config/xxxx_poc_config.h"
#include "ordering/poc/pow/pow_manager.h"
#include "ordering/poc/pow/miner_manager.h"
#include "server/consensus_service.h"
#include "execution/poc_transaction_manager.h"

namespace XXXX {

class ConsensusServicePoW : public ConsensusService {
 public:
  ConsensusServicePoW(const XDBPoCConfig& config, poc::PoCTransactionManager* txn_manager);
  virtual ~ConsensusServicePoW();

  // Start the service.
  void Start() override;

  int ConsensusCommit(std::unique_ptr<Context> context,
                      std::unique_ptr<Request> request) override;

  std::vector<ReplicaInfo> GetReplicas() override;

private:
  int ClientQuery(std::unique_ptr<Context> context, std::unique_ptr<Request> request);

 protected:
  std::unique_ptr<PoWManager> pow_manager_;
  std::unique_ptr<MinerManager> miner_manager_;
  poc::PoCTransactionManager * txn_manager_;
};

}  // namespace XXXX
