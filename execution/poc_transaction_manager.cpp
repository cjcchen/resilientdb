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

#include "execution/poc_transaction_manager.h"
#include "ordering/poc/proto/transaction.pb.h"
#include "proto/transaction.pb.h"

#include <glog/logging.h>

namespace resdb {
namespace poc {

PoCTransactionManager::PoCTransactionManager(const ResDBConfig& config) : TransactionExecutorImpl(false, false) {
}

std::unique_ptr<BatchClientResponse> PoCTransactionManager::ExecuteBatch(
    const BatchClientRequest& request) {

  //LOG(ERROR)<<"batch:"<<request.DebugString();
  for (auto& sub_request : request.client_requests()) {
        ExecuteOne(sub_request.request().data());
  }

  return nullptr;
}

void PoCTransactionManager::ExecuteOne(const std::string& request){
 TransactionsRequest txn_request;
 if(!txn_request.ParseFromString(request)){
   LOG(ERROR)<<"parse txn fail";
   return;
 }
 for(const Transaction& txn: txn_request.transactions()){
    LOG(ERROR)<<"execute txn :"<<txn.DebugString();
     done_.insert(txn.uid());
 }
}

std::unique_ptr<std::string> PoCTransactionManager::ClientQuery(const std::string& str) {
  TransactionQuery request; 
  TransactionQuery response; 

  request.ParseFromString(str);

//LOG(ERROR)<<"client query:"<<request.DebugString();
  for(uint64_t uid : request.uids()){
    if(done_.find(uid) == done_.end()){
      break;
    }
    response.add_uids(uid);
  }
  //LOG(ERROR)<<"get resp:"<<response.DebugString();
  auto ret = std::make_unique<std::string>();
  response.SerializeToString(ret.get());
  return ret;
}

}
}  // namespace resdb
