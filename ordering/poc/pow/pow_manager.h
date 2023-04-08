#pragma once

#include "server/resdb_replica_client.h"
#include "common/queue/blocking_queue.h"
#include "config/resdb_poc_config.h"
#include "ordering/poc/pow/block_manager.h"
#include "ordering/poc/pow/miner_manager.h"
#include "ordering/poc/pow/shift_manager.h"
#include "ordering/poc/pow/transaction_accessor.h"
#include "execution/transaction_executor_impl.h"
#include "ordering/poc/proto/pow.pb.h"

namespace resdb {

class PoWManager {
 public:
  PoWManager(const ResDBPoCConfig& config, ResDBReplicaClient * bc_client, TransactionExecutorImpl * executor);
  virtual ~PoWManager();

  void Start();
  void Stop();
  bool IsRunning();
  void Reset();
  void Commit(std::unique_ptr<Block> block);
  void AddShiftMsg(const SliceInfo& slice_info);
  void ReceiveMiningResult(std::unique_ptr<BlockMiningInfo> info);

  enum MiningStatus {
	  OK = 0,
	  TIMEOUT = 1,
	  FAIL = 2,
  };

  enum BlockStatus {
	  GENERATE_NEW = 0,
	  NEXT_NEWBLOCK = 1,
  };

  enum MiningType {
	  NEWBLOCK = 0,
	  SHIFTING = 1,
  };

 protected:
  virtual std::unique_ptr<TransactionAccessor> GetTransactionAccessor(const ResDBPoCConfig& config);
  virtual std::unique_ptr<ShiftManager> GetShiftManager(const ResDBPoCConfig& config);
  virtual std::unique_ptr<BlockManager> GetBlockManager(const ResDBPoCConfig& config);

  virtual MiningStatus Wait();
  virtual void NotifyBroadCast();
  virtual int GetShiftMsg(const SliceInfo& slice_info);

  int GetMiningTxn(MiningType type);
  void NotifyNextBlock();
  absl::Status WaitBlockDone();
  BlockStatus GetBlockStatus();

  void SendShiftMsg();
  void MiningProcess();
  int BroadCastNewBlock(const Block& block);
  int BroadCastShiftMsg(const SliceInfo& slice_info);
  int BroadCastMiningResult(const BlockMiningInfo& info);

  void MiningResultsProcess();
  void ResultsVerifyProcess();


uint64_t GetFirstResultSeq();
std::unique_ptr<BlockMiningInfo> GetResultInfo(uint64_t seq);
void RemoveResult(uint64_t seq);

bool IsSent(uint64_t seq);
void SetSent(uint64_t seq);

 private:
  ResDBPoCConfig config_;
  uint32_t self_id_;
  std::unique_ptr<BlockManager> block_manager_;
  std::unique_ptr<ShiftManager> shift_manager_;
  std::unique_ptr<TransactionAccessor> transaction_accessor_;
  std::thread miner_thread_, fetch_thread_, verify_thread_;
  std::atomic<bool> is_stop_;

  std::mutex broad_cast_mtx_, mutex_, tx_mutex_, block_result_mutex_;
  std::condition_variable broad_cast_cv_,cv_;
  std::atomic<BlockStatus> current_status_ = BlockStatus::GENERATE_NEW;
  ResDBReplicaClient* bc_client_;
  SliceInfo need_slice_info_;
  PrometheusHandler * prometheus_handler_;

  LockFreeQueue<Request> pending_blocks_results_;
  std::unordered_multimap<uint64_t, std::unique_ptr<BlockMiningInfo>> result_candidate_;
	
  std::atomic<uint64_t> last_done_seq_;
  std::set<uint64_t> sent_list_;
  TransactionExecutorImpl * executor_;
};

}  // namespace resdb
