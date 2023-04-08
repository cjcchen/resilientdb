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

#include <glog/logging.h>

#include "application/utils/server_factory.h"
#include "config/resdb_config_utils.h"
#include "ordering/poc/pow/consensus_service_pow.h"
#include "server/resdb_server.h"
#include "statistic/stats.h"

using namespace resdb;

void ShowUsage() {
  printf("<config> <private_key> <cert_file> [logging_dir]\n");
}

int main(int argc, char** argv) {
  if (argc < 4) {
    ShowUsage();
    exit(0);
  }

  std::string bft_config_file = argv[1];
  std::string pow_config_file = argv[2];
  std::string private_key_file = argv[3];
  std::string cert_file = argv[4];
  LOG(ERROR) << "pow_config:" << pow_config_file;

  std::unique_ptr<ResDBConfig> transaction_server_config = GenerateResDBConfigFromJson(bft_config_file);

  std::unique_ptr<ResDBConfig> mining_config = GenerateResDBConfig(
      pow_config_file, private_key_file, cert_file, std::nullopt,
      [&](const ResConfigData& replicas,
          const ReplicaInfo& self_info, const KeyInfo& private_key,
          const CertificateInfo& public_key_cert_info) {
        return std::make_unique<ResDBPoCConfig>(
            *transaction_server_config, replicas, self_info, private_key, public_key_cert_info);
      });

  LOG(ERROR)<<"elf ip:"<<mining_config->GetSelfInfo().ip();
  //Stats::InitGlobalPrometheus("0.0.0.0:8091");
  ResDBPoCConfig* pow_config_ptr =
      static_cast<ResDBPoCConfig*>(mining_config.get());

  pow_config_ptr->SetMaxNonceBit(42);
  pow_config_ptr->SetDifficulty(28);
  //pow_config_ptr->SetDifficulty(32);

  poc::PoCTransactionManager manager(*pow_config_ptr);
  
  ResDBServer server(*pow_config_ptr,
                     std::make_unique<ConsensusServicePoW>(*pow_config_ptr, &manager));
  server.Run();
}

