#pragma once

#include "config/xxxx_poc_config.h"
#include "ordering/poc/proto/pow.pb.h"
#include <condition_variable>

namespace XXXX {

class ShiftManager {
 public:
  ShiftManager(const XDBPoCConfig& config);
  virtual ~ShiftManager() = default;

  void AddSliceInfo(const SliceInfo& slice_info);
  virtual bool Check(const SliceInfo& slice_info, int timeout_ms=10000);

 private:
  XDBPoCConfig config_;
  std::map<std::pair<uint64_t, uint64_t>, std::set<uint32_t>> data_;
	std::mutex mutex_;
	std::condition_variable cv_;
};

}  // namespace XXXX
