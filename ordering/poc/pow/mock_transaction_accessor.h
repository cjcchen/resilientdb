#pragma once

#include "ordering/poc/pow/transaction_accessor.h"
#include "gmock/gmock.h"

namespace XXXX {

class MockTransactionAccessor : public TransactionAccessor {
 public:
	 MockTransactionAccessor(const XDBPoCConfig& config):TransactionAccessor(config, false){}
  MOCK_METHOD(std::unique_ptr<BatchClientTransactions>, ConsumeTransactions, (uint64_t seq),
              (override));
};

}  // namespace XXXX
