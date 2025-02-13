#pragma once

#include "config/xxxx_poc_config.h"

namespace XXXX {

class MinerManager {
 public:
  MinerManager(const XDBPoCConfig& config);

  std::vector<ReplicaInfo> GetReplicas();

 private:
  XDBPoCConfig config_;
  std::vector<ReplicaInfo> replicas_;
};

}  // namespace XXXX
