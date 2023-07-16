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

#include <glog/logging.h>

#include "application/utils/server_factory.h"
#include "config/xxxx_config_utils.h"
#include "statistic/stats.h"
#include "ordering/poc/pbft/transaction_consensor.h"
#include "ordering/poc/pbft/transaction_executor.h"
#include "ordering/poc/pbft/transaction_query.h"

using namespace XXXX;
using namespace XXXX::poc;

void ShowUsage() {
  printf("<config> <private_key> <cert_file> [logging_dir]\n");
}

int main(int argc, char** argv) {
  if (argc < 4) {
    ShowUsage();
    exit(0);
  }

  char* config_file = argv[1];
  char* private_key_file = argv[2];
  char* cert_file = argv[3];

  std::unique_ptr<XDBConfig> config =
      GenerateXDBConfig(config_file, private_key_file, cert_file);
  ResConfigData config_data = config->GetConfigData();

  XXXX::poc::TransactionManager manager;

  auto server = CustomGenerateXDBServer<TransactionConsensor>(
		  config_file, private_key_file, cert_file,
		  std::make_unique<XXXX::poc::TransactionExecutor>(*config, &manager), std::make_unique<TransactionQuery>(*config, &manager));


  server->Run();
}

