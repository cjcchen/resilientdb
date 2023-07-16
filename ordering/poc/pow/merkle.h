#pragma once

#include "ordering/poc/proto/pow.pb.h"

namespace XXXX {

class Merkle {
 public:
  static HashValue MakeHash(const BatchClientTransactions& transaction);
};

}  // namespace XXXX
