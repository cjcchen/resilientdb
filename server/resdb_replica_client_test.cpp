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

#include "server/xxxx_replica_client.h"

#include <gtest/gtest.h>

#include <future>

#include "client/mock_xxxx_client.h"
#include "common/network/mock_socket.h"
#include "server/mock_async_replica_client.h"

namespace XXXX {
namespace {

using ::testing::Invoke;
using ::testing::Return;

ReplicaInfo GenerateReplicaInfo(const std::string& ip, int port) {
  ReplicaInfo replica_info;
  replica_info.set_ip(ip);
  replica_info.set_port(port);
  return replica_info;
}

class MockXDBReplicaClient : public XDBReplicaClient {
 public:
  MockXDBReplicaClient(const std::vector<ReplicaInfo>& replicas,
                         bool use_lonn_conn = false)
      : XDBReplicaClient(replicas, nullptr, use_lonn_conn){};
  MOCK_METHOD(std::unique_ptr<XDBClient>, GetClient,
              (const std::string&, int), (override));
  MOCK_METHOD(AsyncReplicaClient*, GetClientFromPool, (const std::string&, int),
              (override));
};

TEST(XDBReplicaClientTest, SendMessage) {
  std::vector<ReplicaInfo> replicas;
  replicas.push_back(GenerateReplicaInfo("127.0.0.1", 1234));
  replicas.push_back(GenerateReplicaInfo("127.0.0.1", 1235));

  MockXDBReplicaClient client(replicas);
  EXPECT_CALL(client, GetClient("127.0.0.1", 1234))
      .WillOnce(Invoke([&](const std::string& ip, int port) {
        return std::make_unique<MockXDBClient>(ip, port);
      }));
  EXPECT_CALL(client, GetClient("127.0.0.1", 1235))
      .WillOnce(Invoke([&](const std::string& ip, int port) {
        return std::make_unique<MockXDBClient>(ip, port);
      }));

  Request expected_request;
  expected_request.set_type(Request::TYPE_HEART_BEAT);
  EXPECT_EQ(client.SendMessage(expected_request), 2);
}

TEST(XDBReplicaClientTest, SendMessageToClient) {
  std::vector<ReplicaInfo> replicas;
  replicas.push_back(GenerateReplicaInfo("127.0.0.1", 1234));
  replicas.push_back(GenerateReplicaInfo("127.0.0.1", 1235));

  MockXDBReplicaClient client(replicas);

  std::vector<ReplicaInfo> client_replicas;
  client_replicas.push_back(GenerateReplicaInfo("127.0.0.1", 1236));
  client_replicas.back().set_id(3);

  client.UpdateClientReplicas(client_replicas);

  EXPECT_CALL(client, GetClient("127.0.0.1", 1236))
      .WillOnce(Invoke([&](const std::string& ip, int port) {
        return std::make_unique<MockXDBClient>(ip, port);
      }));

  Request expected_request;
  expected_request.set_type(Request::TYPE_HEART_BEAT);
  client.SendMessage(expected_request, 3);
}

TEST(XDBReplicaClientTest, PartialSendMessage) {
  std::vector<ReplicaInfo> replicas;
  replicas.push_back(GenerateReplicaInfo("127.0.0.1", 1234));
  replicas.push_back(GenerateReplicaInfo("127.0.0.1", 1235));

  MockXDBReplicaClient client(replicas);
  EXPECT_CALL(client, GetClient("127.0.0.1", 1234))
      .WillOnce(Invoke([&](const std::string& ip, int port) {
        return std::make_unique<MockXDBClient>(ip, port);
      }));
  EXPECT_CALL(client, GetClient("127.0.0.1", 1235))
      .WillOnce(
          Invoke([&](const std::string& ip, int port) { return nullptr; }));

  Request expected_request;
  expected_request.set_type(Request::TYPE_HEART_BEAT);
  EXPECT_EQ(client.SendMessage(expected_request), 1);
}

TEST(XDBReplicaClientTest, Lonnconnection) {
  std::promise<bool> bc;
  std::future<bool> bc_done = bc.get_future();
  std::vector<ReplicaInfo> replicas;
  replicas.push_back(GenerateReplicaInfo("127.0.0.1", 1234));

  boost::asio::io_service io_service;
  auto xxxx_client = std::make_unique<MockAsyncReplicaClient>(&io_service);
  MockXDBReplicaClient client(replicas, true);
  EXPECT_CALL(client, GetClientFromPool("127.0.0.1", 1234))
      .WillOnce(Invoke([&](const std::string& ip, int port) {
        bc.set_value(true);
        return xxxx_client.get();
      }));

  Request expected_request;
  expected_request.set_type(Request::TYPE_HEART_BEAT);
  EXPECT_EQ(client.SendMessage(expected_request), 0);
  bc_done.get();
}

}  // namespace

}  // namespace XXXX
