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
#include "config/xxxx_config.h"
#include "execution/transaction_executor_impl.h"
#include "ordering/poc/proto/pow.pb.h"

namespace XXXX {
namespace poc {

class TransactionManager{
 public:
  TransactionManager();

  virtual ~TransactionManager();

  void AddTransactionRequest(const BatchClientRequest& request);
  std::unique_ptr<std::string> GetTransactionRequest(uint64_t seq, bool is_result);

 private:
  std::unordered_map<uint64_t, std::unique_ptr<BatchClientRequest> > txn_;
  std::unordered_map<uint64_t, std::unique_ptr<BlockMiningInfo>> mining_results_;
  std::atomic<uint64_t> seq_;
  std::mutex txn_mutex_, result_mutex_;

};

}
}  // namespace XXXX
