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

#include "ordering/poc/pbft/transaction_query.h"

#include "ordering/poc/proto/pow.pb.h"

#include <glog/logging.h>

namespace resdb {
namespace poc {

TransactionQuery::TransactionQuery(const XDBConfig& config, TransactionManager * manager)
    :manager_(manager){
}

std::unique_ptr<std::string> TransactionQuery::Query(const std::string& request_str) {
  TxnQueryRequest request;
  if(!request.ParseFromString(request_str)){
    LOG(ERROR)<<"parse data fail";
    return nullptr;
  }

  TxnQueryResponse response;

  uint64_t min_seq = request.min_seq();
  uint64_t max_seq = request.max_seq();
  //LOG(ERROR)<<"query:["<<min_seq<<"-"<<max_seq<<"]";
  for(int i = min_seq; i <= max_seq; ++i){
    std::unique_ptr<std::string> ret = manager_->GetTransactionRequest(i, request.is_query_results());
    if(ret == nullptr){
      break;
    }
    response.add_data(*ret);
    response.add_seq(i);
    //LOG(ERROR)<<"query get seq:"<<i<<" is result:"<<request.is_query_results();
  }

  std::unique_ptr<std::string> ret = std::make_unique<std::string>();
  response.SerializeToString(ret.get());
  return ret;
}

}
}  // namespace resdb
