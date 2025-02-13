/*
 * Copyright (c) 2019-2022 XXXX, XXXX
 *
 * Permission is hereby granted, free of charge, to any person
 * obtaining a copy of this software and associated documentation
 * files (the "Software"), to deal in the Software without
 * restriction, including without limitation the rights to use,
 * copy, modify, merge, publish, distribute, sublicense, and/or
 * sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be
 * included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
 * OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
 * HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
 * WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
 * DEALINGS IN THE SOFTWARE.
 *
 */

#pragma once

#include <stdint.h>

#include <map>
#include <memory>
#include <queue>
#include <set>

#include "common/queue/lock_free_queue.h"
#include "config/xxxx_config.h"
#include "database/txn_memory_db.h"
#include "execution/transaction_executor_impl.h"
#include "ordering/pbft/checkpoint_manager.h"
#include "ordering/pbft/lock_free_collector_pool.h"
#include "ordering/pbft/transaction_collector.h"
#include "ordering/pbft/transaction_utils.h"
#include "proto/checkpoint_info.pb.h"
#include "proto/XXXX.pb.h"
#include "server/server_comm.h"
#include "statistic/stats.h"

namespace XXXX {

class TransactionManager {
 public:
  TransactionManager(const XDBConfig& config,
                     std::unique_ptr<TransactionExecutorImpl> data_impl,
                     CheckPointManager* checkpoint_manager,
                     SystemInfo* system_info);
  ~TransactionManager();

  absl::StatusOr<uint64_t> AssignNextSeq();

  int64_t GetCurrentPrimary() const;
  uint64_t GetMinExecutCandidateSeq();
  void SetNextSeq(uint64_t seq);

  // Add commit messages and return the number of messages have been received.
  // The commit messages only include post(pre-prepare), prepare and commit
  // messages. Messages are handled by state (PREPARE,COMMIT,READY_EXECUTE).

  // If there are enough messages and the state is changed after adding the
  // message, return 1, otherwise return 0. Return -2 if the request is not
  // valid.
  CollectorResultCode AddConsensusMsg(const SignatureInfo& signature,
                                      std::unique_ptr<Request> request);

  // Obtain the request that has been executed from Executor.
  // The messages that have been executed from Executor will save inside
  // Message Manager. Consensus Service can obtain the message then send back
  // to the client proxy.
  std::unique_ptr<BatchClientResponse> GetResponseMsg();

  // Get committed messages with 2f+1 proof in [min_seq, max_seq].
  RequestSet GetRequestSet(uint64_t min_seq, uint64_t max_seq);

  // Get the transactions that have been execuited.
  Request* GetRequest(uint64_t seq);

  // Get the proof info containing the request and signatures
  // if the request has been prepared, having received 2f+1
  // pre-prepare messages.
  std::vector<RequestInfo> GetPreparedProof(uint64_t seq);
  TransactionStatue GetTransactionState(uint64_t seq);

  // =============  System information ========
  // Obtain the current replica list.
  std::vector<ReplicaInfo> GetReplicas();

  uint64_t GetCurrentView() const;

  // Replica State
  int GetReplicaState(ReplicaState* state);
  std::unique_ptr<Context> FetchClientContext(uint64_t seq);

 private:
  bool IsValidMsg(const Request& request);

  bool MayConsensusChangeStatus(int type, int received_count,
                                std::atomic<TransactionStatue>* status);

 private:
  XDBConfig config_;
  uint64_t next_seq_ = 1;

  LockFreeQueue<BatchClientResponse> queue_;
  TxnMemoryDB* txn_db_;
  SystemInfo* system_info_;
  CheckPointManager* checkpoint_manager_;
  std::map<uint64_t, std::vector<std::unique_ptr<RequestInfo>>>
      committed_proof_;
  std::map<uint64_t, Request> committed_data_;

  std::mutex data_mutex_, seq_mutex_;
  std::unique_ptr<TransactionExecutor> transaction_executor_;
  std::unique_ptr<LockFreeCollectorPool> collector_pool_;

  Stats* global_stats_;
};

}  // namespace XXXX
