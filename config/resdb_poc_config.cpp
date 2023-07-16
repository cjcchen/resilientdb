#include "config/xxxx_poc_config.h"

#include "glog/logging.h"

namespace resdb {

XDBPoCConfig::XDBPoCConfig(const XDBConfig& bft_config,
                               const ResConfigData& config_data,
                               const ReplicaInfo& self_info,
                               const KeyInfo& private_key,
                               const CertificateInfo& public_key_cert_info)
    : XDBConfig(config_data, self_info, private_key, public_key_cert_info),
      bft_config_(bft_config) {
  SetHeartBeatEnabled(false);
  SetSignatureVerifierEnabled(false);
  
  for (const auto& region : config_data.region()) {
    if (region.region_id() == config_data.self_region_id()) {
      for (const auto& replica : region.replica_info()) {
        bft_replicas_.push_back(replica);
      }
      break;
    }
  }

}

const XDBConfig* XDBPoCConfig::GetBFTConfig() const { return &bft_config_; }

void XDBPoCConfig::SetMaxNonceBit(uint32_t bit) { max_nonce_bit_ = bit; }

uint32_t XDBPoCConfig::GetMaxNonceBit() const { return max_nonce_bit_; }

void XDBPoCConfig::SetDifficulty(uint32_t difficulty) {
  difficulty_ = difficulty;
}

uint32_t XDBPoCConfig::GetDifficulty() const { return difficulty_; }

uint32_t XDBPoCConfig::GetTargetValue() const { return target_value_; }

void XDBPoCConfig::SetTargetValue(uint32_t target_value) {
  target_value_ = target_value;
}

std::vector<ReplicaInfo> XDBPoCConfig::GetBFTReplicas() {
  return bft_replicas_;
}

void XDBPoCConfig::SetBFTReplicas(const std::vector<ReplicaInfo>& replicas) {
  bft_replicas_ = replicas;
}

// Batch
uint32_t XDBPoCConfig::BatchTransactionNum() const { 
	return batch_num_; }

void XDBPoCConfig::SetBatchTransactionNum(uint32_t batch_num) {
  batch_num_ = batch_num;
}

uint32_t XDBPoCConfig::GetWokerNum() { return worker_num_; }

void XDBPoCConfig::SetWorkerNum(uint32_t worker_num) {
  worker_num_ = worker_num;
}

}  // namespace resdb
