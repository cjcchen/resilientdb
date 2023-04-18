/*
 * Copyright (c) 2019-2022 ExpoLab, UC Davis
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

#include "ordering/poc/pbft/transaction_manager.h"

#include <glog/logging.h>
#include "proto/transaction.pb.h"

namespace resdb {
namespace poc {

TransactionManager::TransactionManager() { 
  seq_ = 1;
}

TransactionManager::~TransactionManager() {
}

void TransactionManager::AddTransactionRequest(const BatchClientRequest& request) {
  //LOG(ERROR)<<"execute batch request, is sysdata:"<<request.system_data()<<" seq:"<<request.seq();
  if(request.system_data()){
    std::unique_ptr<BlockMiningInfo> info = std::make_unique<BlockMiningInfo>();
    info->ParseFromString(request.ex_data());

    std::unique_lock<std::mutex> lck(result_mutex_);
    if(mining_results_.find(info->header().min_seq()) != mining_results_.end()){
      LOG(ERROR)<<"result has been committed:"<<info->header().min_seq();
      return;
    }
    mining_results_[info->header().min_seq()] = std::move(info);
  }
  else {
    std::unique_lock<std::mutex> lck(txn_mutex_);
    txn_[seq_++] = std::make_unique<BatchClientRequest>(request);
  }
}

std::unique_ptr<std::string> TransactionManager::GetTransactionRequest(uint64_t seq, bool is_result) {
  //LOG(ERROR)<<"get data:"<<seq<<" result:"<<is_result;
  std::unique_ptr<std::string> str = std::make_unique<std::string>();
  if(is_result){
    std::unique_lock<std::mutex> lck(result_mutex_);
    auto it = mining_results_.find(seq);
    if(it == mining_results_.end()) {
      return nullptr;
    }
    it->second->SerializeToString(str.get());
  }
  else {
    std::unique_lock<std::mutex> lck(txn_mutex_);
    auto it = txn_.find(seq);
    if(it == txn_.end()) {
    /*
      BatchClientRequest client_request;
      auto * request = client_request.add_client_requests();
      Transaction txn;
      txn.set_uid(seq);
      txn.SerializeToString(request->mutable_request()->mutable_data());

      client_request.SerializeToString(str.get());
      return str;
      */
      return nullptr;
    }
    it->second->SerializeToString(str.get());
  }
  return str;
}


}
}  // namespace resdb
