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

#include "config/xxxx_config.h"

#include <glog/logging.h>

namespace XXXX {

XDBConfig::XDBConfig(const std::vector<ReplicaInfo>& replicas,
                         const ReplicaInfo& self_info,
                         ResConfigData config_data)
    : config_data_(config_data), replicas_(replicas), self_info_(self_info) {}

XDBConfig::XDBConfig(const std::vector<ReplicaInfo>& replicas,
                         const ReplicaInfo& self_info,
                         const KeyInfo& private_key,
                         const CertificateInfo& public_key_cert_info)
    : replicas_(replicas),
      self_info_(self_info),
      private_key_(private_key),
      public_key_cert_info_(public_key_cert_info) {}

XDBConfig::XDBConfig(const ResConfigData& config_data,
                         const ReplicaInfo& self_info,
                         const KeyInfo& private_key,
                         const CertificateInfo& public_key_cert_info)
    : config_data_(config_data),
      self_info_(self_info),
      private_key_(private_key),
      public_key_cert_info_(public_key_cert_info) {
  for (const auto& region : config_data.region()) {
    if (region.region_id() == config_data.self_region_id()) {
      LOG(INFO) << "get region info:" << region.DebugString();
      for (const auto& replica : region.replica_info()) {
        replicas_.push_back(replica);
      }
      LOG(INFO) << "get region config server size:"
                << region.replica_info_size();
      break;
    }
  }
  if (config_data_.view_change_timeout_ms() == 0) {
    config_data_.set_view_change_timeout_ms(viewchange_commit_timeout_ms_);
  }
  if (config_data_.client_batch_num() == 0) {
    config_data_.set_client_batch_num(client_batch_num_);
  }
  if (config_data_.worker_num() == 0) {
    config_data_.set_worker_num(worker_num_);
  }
  if (config_data_.input_worker_num() == 0) {
    config_data_.set_input_worker_num(input_worker_num_);
  }
  if (config_data_.output_worker_num() == 0) {
    config_data_.set_output_worker_num(output_worker_num_);
  }
  if (config_data_.tcp_batch_num() == 0) {
    config_data_.set_tcp_batch_num(100);
  }
}

void XDBConfig::SetConfigData(const ResConfigData& config_data) {
  config_data_ = config_data;
  replicas_.clear();
  for (const auto& region : config_data.region()) {
    if (region.region_id() == config_data.self_region_id()) {
      LOG(INFO) << "get region info:" << region.DebugString();
      for (const auto& replica : region.replica_info()) {
        replicas_.push_back(replica);
      }
      LOG(INFO) << "get region config server size:"
                << region.replica_info_size();
      break;
    }
  }
  if (config_data_.view_change_timeout_ms() == 0) {
    config_data_.set_view_change_timeout_ms(viewchange_commit_timeout_ms_);
  }
}

KeyInfo XDBConfig::GetPrivateKey() const { return private_key_; }

CertificateInfo XDBConfig::GetPublicKeyCertificateInfo() const {
  return public_key_cert_info_;
}

ResConfigData XDBConfig::GetConfigData() const { return config_data_; }

const std::vector<ReplicaInfo>& XDBConfig::GetReplicaInfos() const {
  return replicas_;
}

const ReplicaInfo& XDBConfig::GetSelfInfo() const { return self_info_; }

size_t XDBConfig::GetReplicaNum() const { return replicas_.size(); }

int XDBConfig::GetMinDataReceiveNum() const {
  int f = (replicas_.size() - 1) / 3;
  return std::max(2 * f + 1, 1);
}

int XDBConfig::GetMinClientReceiveNum() const {
  int f = (replicas_.size() - 1) / 3;
  return std::max(f + 1, 1);
}

size_t XDBConfig::GetMaxMaliciousReplicaNum() const {
  int f = (replicas_.size() - 1) / 3;
  return std::max(f, 0);
}

void XDBConfig::SetClientTimeoutMs(int timeout_ms) {
  client_timeout_ms_ = timeout_ms;
}

int XDBConfig::GetClientTimeoutMs() const { return client_timeout_ms_; }

// Logging
std::string XDBConfig::GetCheckPointLoggingPath() const {
  return checkpoint_logging_path_;
}

void XDBConfig::SetCheckPointLoggingPath(const std::string& path) {
  checkpoint_logging_path_ = path;
}

int XDBConfig::GetCheckPointWaterMark() const {
  return checkpoint_water_mark_;
}

void XDBConfig::SetCheckPointWaterMark(int water_mark) {
  checkpoint_water_mark_ = water_mark;
}

void XDBConfig::EnableCheckPoint(bool is_enable) {
  is_enable_checkpoint_ = is_enable;
}

bool XDBConfig::IsCheckPointEnabled() { return is_enable_checkpoint_; }

bool XDBConfig::HeartBeatEnabled() { return hb_enabled_; }

void XDBConfig::SetHeartBeatEnabled(bool enable_heartbeat) {
  hb_enabled_ = enable_heartbeat;
}

bool XDBConfig::SignatureVerifierEnabled() {
  return signature_verifier_enabled_;
}

void XDBConfig::SetSignatureVerifierEnabled(bool enable_sv) {
  signature_verifier_enabled_ = enable_sv;
}

// Performance setting
bool XDBConfig::IsPerformanceRunning() {
  return is_performance_running_ || GetConfigData().is_performance_running();
}

void XDBConfig::RunningPerformance(bool is_performance_running) {
  is_performance_running_ = is_performance_running;
}

void XDBConfig::SetTestMode(bool is_test_mode) {
  is_test_mode_ = is_test_mode;
}

bool XDBConfig::IsTestMode() const { return is_test_mode_; }

uint32_t XDBConfig::GetMaxProcessTxn() const {
  if (config_data_.max_process_txn()) {
    return config_data_.max_process_txn();
  }
  return max_process_txn_;
}

void XDBConfig::SetMaxProcessTxn(uint32_t num) {
  config_data_.set_max_process_txn(num);
  max_process_txn_ = num;
}

uint32_t XDBConfig::ClientBatchWaitTimeMS() const {
  return client_batch_wait_time_ms_;
}

void XDBConfig::SetClientBatchWaitTimeMS(uint32_t wait_time_ms) {
  client_batch_wait_time_ms_ = wait_time_ms;
}

uint32_t XDBConfig::ClientBatchNum() const {
  return config_data_.client_batch_num();
}

void XDBConfig::SetClientBatchNum(uint32_t num) {
  config_data_.set_client_batch_num(num);
}

uint32_t XDBConfig::GetWorkerNum() const { return config_data_.worker_num(); }

uint32_t XDBConfig::GetInputWorkerNum() const {
  return config_data_.input_worker_num();
}

uint32_t XDBConfig::GetOutputWorkerNum() const {
  return config_data_.output_worker_num();
}

uint32_t XDBConfig::GetTcpBatchNum() const {
  return config_data_.tcp_batch_num();
}

uint32_t XDBConfig::GetViewchangeCommitTimeout() const {
  return viewchange_commit_timeout_ms_;
}

void XDBConfig::SetViewchangeCommitTimeout(uint64_t timeout_ms) {
  viewchange_commit_timeout_ms_ = timeout_ms;
}

}  // namespace XXXX
