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

#include "ordering/poc/pbft/transaction_consensor.h"

#include "ordering/poc/pbft/transaction_executor.h"
#include "ordering/poc/pow/miner_utils.h"

namespace resdb {
namespace poc {

TransactionConsensor::TransactionConsensor(const ResDBConfig& config, 
  std::unique_ptr<TransactionExecutorImpl> executor, std::unique_ptr<CustomQuery> query)
    : ConsensusServicePBFT(config, std::move(executor), std::move(query)) {
}

int TransactionConsensor::ConsensusCommit(std::unique_ptr<Context> context,
                                             std::unique_ptr<Request> request) {
  switch (request->type()) {
    case Request::TYPE_GEO_MINING_RESULT:
      return SaveResult(std::move(request));
  }
  return ConsensusServicePBFT::ConsensusCommit(std::move(context),
                                               std::move(request));
}

int TransactionConsensor::SaveResult(std::unique_ptr<Request> request){
  int received_size = 0;
  //LOG(ERROR)<<"receive result";
  std::unique_ptr<BlockMiningInfo> mining_result = std::make_unique<BlockMiningInfo>();

  if(!mining_result->ParseFromString(request->data())){
    LOG(ERROR)<<"parse data fail";
    return -2;
  }

  uint64_t seq = mining_result->header().min_seq();
  //LOG(ERROR)<<"get result seq:"<<seq;
  std::string hash = GetHashDigest(mining_result->hash());
  {
    std::unique_lock<std::mutex> lck(mutex_);
    results_[std::make_pair(seq, hash)].push_back(std::move(mining_result));
    received_size = results_[std::make_pair(seq, hash)].size();
  }
  //LOG(ERROR)<<"receive seq:"<<seq<<" size:"<<received_size;

  if(received_size != 3){
    return 0;
  }

  BatchClientRequest batch_request;
  batch_request.set_ex_data(request->data());
  batch_request.set_system_data(true);

  auto new_request =
    NewRequest(Request::TYPE_NEW_TXNS, Request(), config_.GetSelfInfo().id());

  batch_request.SerializeToString(new_request->mutable_data());
  if(verifier_) {
    auto signature_or = verifier_->SignMessage(new_request->data());
    if (!signature_or.ok()) {
      LOG(ERROR) << "Sign message fail";
      return -2;
    }
    *new_request->mutable_data_signature() = *signature_or;
  }

  new_request->set_hash(SignatureVerifier::CalculateHash(new_request->data()));
  new_request->set_proxy_id(config_.GetSelfInfo().id());

  return ConsensusServicePBFT::SelfPropose(std::move(new_request));
}

}
}  // namespace resdb
