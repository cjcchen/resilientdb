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

#include "config/xxxx_config_utils.h"
#include "execution/custom_query.h"
#include "execution/transaction_executor_impl.h"
#include "ordering/pbft/consensus_service_pbft.h"
#include "server/xxxx_server.h"

namespace XXXX {

class ServerFactory {
 public:
  std::unique_ptr<XDBServer> CreateXDBServer(
      char* config_file, char* private_key_file, char* cert_file,
      std::unique_ptr<TransactionExecutorImpl> executor, char* logging_dir,
      std::function<void(XDBConfig* config)> config_handler);

  template <typename ConsensusProtocol = ConsensusServicePBFT>
  std::unique_ptr<XDBServer> CustomCreateXDBServer(
      char* config_file, char* private_key_file, char* cert_file,
      std::unique_ptr<TransactionExecutorImpl> executor, char* logging_dir,
      std::function<void(XDBConfig* config)> config_handler);

  template <typename ConsensusProtocol = ConsensusServicePBFT>
  std::unique_ptr<XDBServer> CustomCreateXDBServer(
      char* config_file, char* private_key_file, char* cert_file,
      std::unique_ptr<TransactionExecutorImpl> executor,
      std::unique_ptr<CustomQuery> query_executor,
      std::function<void(XDBConfig* config)> config_handler);
};

std::unique_ptr<XDBServer> GenerateXDBServer(
    char* config_file, char* private_key_file, char* cert_file,
    std::unique_ptr<TransactionExecutorImpl> executor,
    char* logging_dir = nullptr,
    std::function<void(XDBConfig* config)> config_handler = nullptr);

template <typename ConsensusProtocol>
std::unique_ptr<XDBServer> CustomGenerateXDBServer(
    char* config_file, char* private_key_file, char* cert_file,
    std::unique_ptr<TransactionExecutorImpl> executor,
    char* logging_dir = nullptr,
    std::function<void(XDBConfig* config)> config_handler = nullptr);

template <typename ConsensusProtocol>
std::unique_ptr<XDBServer> CustomGenerateXDBServer(
    char* config_file, char* private_key_file, char* cert_file,
    std::unique_ptr<TransactionExecutorImpl> executor,
    std::unique_ptr<CustomQuery> query_executor,
    std::function<void(XDBConfig* config)> config_handler = nullptr);

// ===================================================================
template <typename ConsensusProtocol>
std::unique_ptr<XDBServer> ServerFactory::CustomCreateXDBServer(
    char* config_file, char* private_key_file, char* cert_file,
    std::unique_ptr<TransactionExecutorImpl> executor, char* logging_dir,
    std::function<void(XDBConfig* config)> config_handler) {
  std::unique_ptr<XDBConfig> config =
      GenerateXDBConfig(config_file, private_key_file, cert_file);

  if (config_handler) {
    config_handler(config.get());
  }
  return std::make_unique<XDBServer>(
      *config,
      std::make_unique<ConsensusProtocol>(*config, std::move(executor)));
}

template <typename ConsensusProtocol>
std::unique_ptr<XDBServer> ServerFactory::CustomCreateXDBServer(
    char* config_file, char* private_key_file, char* cert_file,
    std::unique_ptr<TransactionExecutorImpl> executor,
    std::unique_ptr<CustomQuery> query_executor,
    std::function<void(XDBConfig* config)> config_handler) {
  std::unique_ptr<XDBConfig> config =
      GenerateXDBConfig(config_file, private_key_file, cert_file);

  if (config_handler) {
    config_handler(config.get());
  }
  return std::make_unique<XDBServer>(
      *config, std::make_unique<ConsensusProtocol>(*config, std::move(executor),
                                                   std::move(query_executor)));
}

template <typename ConsensusProtocol>
std::unique_ptr<XDBServer> CustomGenerateXDBServer(
    char* config_file, char* private_key_file, char* cert_file,
    std::unique_ptr<TransactionExecutorImpl> executor, char* logging_dir,
    std::function<void(XDBConfig* config)> config_handler) {
  return ServerFactory().CustomCreateXDBServer<ConsensusProtocol>(
      config_file, private_key_file, cert_file, std::move(executor),
      logging_dir, config_handler);
}

template <typename ConsensusProtocol>
std::unique_ptr<XDBServer> CustomGenerateXDBServer(
    char* config_file, char* private_key_file, char* cert_file,
    std::unique_ptr<TransactionExecutorImpl> executor,
    std::unique_ptr<CustomQuery> query_executor,
    std::function<void(XDBConfig* config)> config_handler) {
  return ServerFactory().CustomCreateXDBServer<ConsensusProtocol>(
      config_file, private_key_file, cert_file, std::move(executor),
      std::move(query_executor), config_handler);
}

}  // namespace XXXX
