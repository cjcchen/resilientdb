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

#include "kv_server/kv_server_executor.h"

#include <glog/logging.h>

#include "proto/transaction.pb.h"

namespace resdb {

KVServerExecutor::KVServerExecutor(const ResConfigData& config_data,
                                   char* cert_file)
    : l_storage_layer_(cert_file, config_data),
      r_storage_layer_(cert_file, config_data) {
  equip_rocksdb_ = config_data.rocksdb_info().enable_rocksdb();
  equip_leveldb_ = config_data.leveldb_info().enable_leveldb();
}

KVServerExecutor::KVServerExecutor(void) {}

std::unique_ptr<std::string> KVServerExecutor::ExecuteData(
    const std::string& request) {
  TransactionsRequest tx_request;
  TransactionsResponse tx_response;

  if (!tx_request.ParseFromString(request)) {
    LOG(ERROR) << "parse data fail";
    return nullptr;
  }

  LOG(ERROR)<<"get txn:"<<tx_request.DebugString();
  std::unique_ptr<std::string> resp_str = std::make_unique<std::string>();

  for(const auto& txn : tx_request.transactions()){
    Set(txn.from(), -txn.amount());
    Set(txn.to(), txn.amount());
    auto * res = tx_response.add_result();
    res->set_ret(1);
    res->set_uid(txn.uid());
  }

  tx_response.SerializeToString(resp_str.get());
  return resp_str;
}

void KVServerExecutor::Set(const std::string& key, const int64_t& value) {
    if(value > 0){
      kv_map_[key] += value;
    }
    else {
      int64_t old_value = kv_map_[key];
      if(old_value==0){
        old_value = 10000;
      }
      kv_map_[key] -= value;
    }
}

}  // namespace resdb
