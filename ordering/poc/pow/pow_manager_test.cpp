#include "ordering/poc/pow/pow_manager.h"

#include <glog/logging.h>
#include <gmock/gmock.h>
#include <gtest/gtest.h>

#include "common/test/test_macros.h"
#include "config/xxxx_config_utils.h"
#include "ordering/poc/pow/mock_transaction_accessor.h"
#include "server/mock_xxxx_replica_client.h"

namespace XXXX {
namespace {

using ::XXXX::testing::EqualsProto;
using ::testing::ElementsAre;
using ::testing::Invoke;
using ::testing::Return;
using ::testing::Pair;
using ::testing::_;
using ::testing::Pointee;
using ::testing::Test;

class MockPoWManager : public PoWManager {
public:
	MockPoWManager(const XDBPoCConfig& config):PoWManager(config, &mock_replica_client_){
	}

  std::unique_ptr<TransactionAccessor> GetTransactionAccessor(const XDBPoCConfig& config) override {
	  auto accessor = std::make_unique<MockTransactionAccessor>(config);
	  mock_transaction_accessor_ = accessor.get();
	  return accessor;
  }
  MOCK_METHOD(int, GetShiftMsg, (const SliceInfo& slice_info), (override));
  MOCK_METHOD(void, NotifyBroadCast, (), (override));
  MockTransactionAccessor * mock_transaction_accessor_;
  MockXDBReplicaClient mock_replica_client_;
};

class PoWManagerBaseTest: public Test {
 protected:
  PoWManagerBaseTest():
	  bft_config_({GenerateReplicaInfo(1, "127.0.0.1", 1234),
                     GenerateReplicaInfo(2, "127.0.0.1", 1235),
                     GenerateReplicaInfo(3, "127.0.0.1", 1236),
                     GenerateReplicaInfo(4, "127.0.0.1", 1237)},
                    GenerateReplicaInfo(3, "127.0.0.1", 1234), KeyInfo(),
                    CertificateInfo()),
	  config_(bft_config_,
                {GenerateReplicaInfo(1, "127.0.0.1", 1234),
                 GenerateReplicaInfo(2, "127.0.0.1", 1235),
                 GenerateReplicaInfo(3, "127.0.0.1", 1236),
                 GenerateReplicaInfo(4, "127.0.0.1", 1237)},
                GenerateReplicaInfo(1, "127.0.0.1", 1234), KeyInfo(),
                CertificateInfo()){ 
		    config_.SetMaxNonceBit(10);
		    config_.SetDifficulty(1);
		    pow_manager_= std::make_unique<MockPoWManager>(config_);
  }

  XDBConfig bft_config_;
  XDBPoCConfig config_;
  std::unique_ptr<MockPoWManager> pow_manager_;
};

/*
TEST_F(PoWManagerBaseTest, NoData) {
	pow_manager_->Reset();
  std::promise<bool> done;
  std::future<bool> done_future = done.get_future();
	int call_time = 0;
	ASSERT_TRUE(pow_manager_->mock_transaction_accessor_!=nullptr);
	EXPECT_CALL(*pow_manager_->mock_transaction_accessor_, ConsumeTransactions).WillRepeatedly(Invoke([&](uint64_t seq){
				LOG(ERROR)<<"seq:"<<seq;
				call_time++;
				if(call_time>2){
					done.set_value(true);
				}
			return nullptr;
			}));
	EXPECT_CALL(*pow_manager_, NotifyBroadCast).Times(0);
	pow_manager_->Start();
  done_future.get();
}
*/

TEST_F(PoWManagerBaseTest, GenerateOneBlock) {
	pow_manager_->Reset();
	std::promise<bool> done;
	std::future<bool> done_future = done.get_future();
	ASSERT_TRUE(pow_manager_->mock_transaction_accessor_!=nullptr);
	EXPECT_CALL(*pow_manager_->mock_transaction_accessor_, ConsumeTransactions(_)).WillRepeatedly(Invoke([&](uint64_t seq){
				int batch_num = 1;
				std::unique_ptr<BatchClientRequest> batch_request = std::make_unique<BatchClientRequest>();
				for(int j = 0; j < 10; ++j){
				auto req = batch_request->add_client_requests();
				req->mutable_request()->set_data(std::to_string(rand()%10000));
				req->mutable_request()->set_seq(j+seq);
				}
				std::string data;
				batch_request->SerializeToString(&data);

				std::unique_ptr<BatchClientTransactions> batch_transactions =
				std::make_unique<BatchClientTransactions>();
				for (uint32_t i = 0; i < batch_num; ++i) {
				std::unique_ptr<ClientTransactions> client_txn = std::make_unique<ClientTransactions>();
				client_txn->set_transaction_data(data);
				client_txn->set_seq(seq+i);

				*batch_transactions->add_transactions() = *client_txn;
				}
				batch_transactions->set_min_seq(seq);
				batch_transactions->set_max_seq(seq + batch_num - 1);
				return batch_transactions;
	}));
	EXPECT_CALL(*pow_manager_, NotifyBroadCast).WillOnce(Invoke([&](){
				done.set_value(true);
				}));
	pow_manager_->Start();
	done_future.get();
}

/*
TEST_F(PoWManagerBaseTest, MineBlockFail) {
        config_.SetMaxNonceBit(5);
	config_.SetDifficulty(9);
	pow_manager_= std::make_unique<MockPoWManager>(config_);
	pow_manager_->Reset();
	std::promise<bool> done;
	std::future<bool> done_future = done.get_future();
	int call_time = 0;
	ASSERT_TRUE(pow_manager_->mock_transaction_accessor_!=nullptr);

	EXPECT_CALL(*pow_manager_->mock_transaction_accessor_, ConsumeTransactions(_)).WillRepeatedly(Invoke([&](uint64_t seq){
				if(seq>1) {
					return (std::unique_ptr<XXXX::BatchClientTransactions>)nullptr;
						}
				int batch_num = 1;
				std::unique_ptr<BatchClientRequest> batch_request = std::make_unique<BatchClientRequest>();
				for(int j = 0; j < 10; ++j){
				auto req = batch_request->add_client_requests();
				req->mutable_request()->set_data(std::to_string(rand()%10000));
				req->mutable_request()->set_seq(j+seq);
				}
				std::string data;
				batch_request->SerializeToString(&data);

				std::unique_ptr<BatchClientTransactions> batch_transactions =
				std::make_unique<BatchClientTransactions>();
				for (uint32_t i = 0; i < batch_num; ++i) {
				std::unique_ptr<ClientTransactions> client_txn = std::make_unique<ClientTransactions>();
				client_txn->set_transaction_data(data);
				client_txn->set_seq(seq+i);

				*batch_transactions->add_transactions() = *client_txn;
				}
				batch_transactions->set_min_seq(seq);
				batch_transactions->set_max_seq(seq + batch_num - 1);
				return batch_transactions;
	}));
	EXPECT_CALL(*pow_manager_, GetShiftMsg).WillOnce(Invoke([&](const SliceInfo& slice_info){
				done.set_value(true);
		  return absl::NotFoundError("No new transaction.");
				}));
	pow_manager_->Start();
	done_future.get();
}

TEST_F(PoWManagerBaseTest, RecvCommitMsg) {
        config_.SetMaxNonceBit(5);
	config_.SetDifficulty(9);
	pow_manager_= std::make_unique<MockPoWManager>(config_);
	pow_manager_->Reset();
	std::promise<bool> done;
	std::future<bool> done_future = done.get_future();
	int call_time = 0;
	ASSERT_TRUE(pow_manager_->mock_transaction_accessor_!=nullptr);

	EXPECT_CALL(*pow_manager_->mock_transaction_accessor_, ConsumeTransactions(_)).WillRepeatedly(Invoke([&](uint64_t seq){
				if(seq>1) {
					return (std::unique_ptr<XXXX::BatchClientTransactions>)nullptr;
						}
				int batch_num = 1;
				std::unique_ptr<BatchClientRequest> batch_request = std::make_unique<BatchClientRequest>();
				for(int j = 0; j < 10; ++j){
				auto req = batch_request->add_client_requests();
				req->mutable_request()->set_data(std::to_string(rand()%10000));
				req->mutable_request()->set_seq(j+seq);
				}
				std::string data;
				batch_request->SerializeToString(&data);

				std::unique_ptr<BatchClientTransactions> batch_transactions =
				std::make_unique<BatchClientTransactions>();
				for (uint32_t i = 0; i < batch_num; ++i) {
				std::unique_ptr<ClientTransactions> client_txn = std::make_unique<ClientTransactions>();
				client_txn->set_transaction_data(data);
				client_txn->set_seq(seq+i);

				*batch_transactions->add_transactions() = *client_txn;
				}
				batch_transactions->set_min_seq(seq);
				batch_transactions->set_max_seq(seq + batch_num - 1);
				return batch_transactions;
	}));
	EXPECT_CALL(*pow_manager_, GetShiftMsg).WillOnce(Invoke([&](const SliceInfo& slice_info){
				done.set_value(true);
		  return absl::NotFoundError("No new transaction.");
				}));
	pow_manager_->Start();
	done_future.get();
}
*/

class MockBlockManager : public BlockManager {
	public:
		MockBlockManager(const XDBPoCConfig& config):BlockManager(config){}
  MOCK_METHOD(absl::Status, Mine, (), (override));
};

class MockShiftManager : public ShiftManager {
	public:
		MockShiftManager(const XDBPoCConfig& config):ShiftManager(config){}
  MOCK_METHOD(bool, Check, (const SliceInfo& slice_info, int), (override));
};

class MockPoWManagerWithBC : public PoWManager {
public:
	MockPoWManagerWithBC(const XDBPoCConfig& config):PoWManager(config, &mock_replica_client_){
	}

  std::unique_ptr<TransactionAccessor> GetTransactionAccessor(const XDBPoCConfig& config) override {
	  auto accessor = std::make_unique<MockTransactionAccessor>(config);
	  mock_transaction_accessor_ = accessor.get();
	  return accessor;
  }
  
  std::unique_ptr<ShiftManager> GetShiftManager(const XDBPoCConfig& config) override {
	  auto manager = std::make_unique<MockShiftManager>(config);
	  mock_shift_manager_= manager.get();
	  return manager ;
  }
 std::unique_ptr<BlockManager> GetBlockManager(const XDBPoCConfig& config) override {
	  auto manager = std::make_unique<MockBlockManager>(config);
	  mock_block_manager_= manager.get();
	  return manager ;
  }


  MockTransactionAccessor * mock_transaction_accessor_;
  MockXDBReplicaClient mock_replica_client_;
  MockShiftManager * mock_shift_manager_;
  MockBlockManager * mock_block_manager_;
};

class PoWManagerTest: public Test {
 protected:
  PoWManagerTest():
	  bft_config_({GenerateReplicaInfo(1, "127.0.0.1", 1234),
                     GenerateReplicaInfo(2, "127.0.0.1", 1235),
                     GenerateReplicaInfo(3, "127.0.0.1", 1236),
                     GenerateReplicaInfo(4, "127.0.0.1", 1237)},
                    GenerateReplicaInfo(3, "127.0.0.1", 1234), KeyInfo(),
                    CertificateInfo()),
	  config_(bft_config_,
                {GenerateReplicaInfo(1, "127.0.0.1", 1234),
                 GenerateReplicaInfo(2, "127.0.0.1", 1235),
                 GenerateReplicaInfo(3, "127.0.0.1", 1236),
                 GenerateReplicaInfo(4, "127.0.0.1", 1237)},
                GenerateReplicaInfo(1, "127.0.0.1", 1234), KeyInfo(),
                CertificateInfo()){ 
		    config_.SetMaxNonceBit(10);
		    config_.SetDifficulty(1);
		    pow_manager_= std::make_unique<MockPoWManagerWithBC>(config_);
  }

  XDBConfig bft_config_;
  XDBPoCConfig config_;
  std::unique_ptr<MockPoWManagerWithBC> pow_manager_;
};

TEST_F(PoWManagerTest, Broadcast) {
	pow_manager_->Reset();
	std::promise<bool> done;
	std::future<bool> done_future = done.get_future();
	int call_time = 0;
	ASSERT_TRUE(pow_manager_->mock_transaction_accessor_!=nullptr);
	EXPECT_CALL(*pow_manager_->mock_transaction_accessor_, ConsumeTransactions(_)).WillRepeatedly(Invoke([&](uint64_t seq){
				int batch_num = 1;
				std::unique_ptr<BatchClientRequest> batch_request = std::make_unique<BatchClientRequest>();
				for(int j = 0; j < 10; ++j){
				auto req = batch_request->add_client_requests();
				req->mutable_request()->set_data(std::to_string(rand()%10000));
				req->mutable_request()->set_seq(j+seq);
				}
				std::string data;
				batch_request->SerializeToString(&data);

				std::unique_ptr<BatchClientTransactions> batch_transactions =
				std::make_unique<BatchClientTransactions>();
				for (uint32_t i = 0; i < batch_num; ++i) {
				std::unique_ptr<ClientTransactions> client_txn = std::make_unique<ClientTransactions>();
				client_txn->set_transaction_data(data);
				client_txn->set_seq(seq+i);

				*batch_transactions->add_transactions() = *client_txn;
				}
				batch_transactions->set_min_seq(seq);
				batch_transactions->set_max_seq(seq + batch_num - 1);
				return batch_transactions;
	}));
	EXPECT_CALL(pow_manager_->mock_replica_client_, BroadCast).WillOnce(Invoke([&](const google::protobuf::Message& message){
				done.set_value(true);
				}));
	pow_manager_->Start();
	done_future.get();
}

TEST_F(PoWManagerTest, Broadcast2) {
	pow_manager_->Reset();
	std::promise<bool> done;
	std::future<bool> done_future = done.get_future();
	int call_time = 0;
	ASSERT_TRUE(pow_manager_->mock_transaction_accessor_!=nullptr);
	EXPECT_CALL(*pow_manager_->mock_transaction_accessor_, ConsumeTransactions(_)).WillRepeatedly(Invoke([&](uint64_t seq){
				int batch_num = 1;
				std::unique_ptr<BatchClientRequest> batch_request = std::make_unique<BatchClientRequest>();
				for(int j = 0; j < 10; ++j){
				auto req = batch_request->add_client_requests();
				req->mutable_request()->set_data(std::to_string(rand()%10000));
				req->mutable_request()->set_seq(j+seq);
				}
				std::string data;
				batch_request->SerializeToString(&data);

				std::unique_ptr<BatchClientTransactions> batch_transactions =
				std::make_unique<BatchClientTransactions>();
				for (uint32_t i = 0; i < batch_num; ++i) {
				std::unique_ptr<ClientTransactions> client_txn = std::make_unique<ClientTransactions>();
				client_txn->set_transaction_data(data);
				client_txn->set_seq(seq+i);

				*batch_transactions->add_transactions() = *client_txn;
				}
				batch_transactions->set_min_seq(seq);
				batch_transactions->set_max_seq(seq + batch_num - 1);
				return batch_transactions;
	}));
	EXPECT_CALL(pow_manager_->mock_replica_client_, BroadCast).WillRepeatedly(Invoke([&](const google::protobuf::Message& message){
				call_time++;
				if(call_time>1){
				done.set_value(true);
				}
				}));
	pow_manager_->Start();
	done_future.get();
}

TEST_F(PoWManagerTest, SendShift) {
        config_.SetMaxNonceBit(5);
	config_.SetDifficulty(9);
	pow_manager_= std::make_unique<MockPoWManagerWithBC>(config_);
	pow_manager_->Reset();
	std::promise<bool> done;
	std::future<bool> done_future = done.get_future();
	ASSERT_TRUE(pow_manager_->mock_transaction_accessor_!=nullptr);

	EXPECT_CALL(*pow_manager_->mock_shift_manager_, Check).WillRepeatedly(Invoke([&](const SliceInfo &info, int){
				LOG(ERROR)<<"check:";
				return true;
				}));
	EXPECT_CALL(*pow_manager_->mock_transaction_accessor_, ConsumeTransactions(_)).WillRepeatedly(Invoke([&](uint64_t seq){
				if(seq>1) {
					return (std::unique_ptr<XXXX::BatchClientTransactions>)nullptr;
						}
				int batch_num = 1;
				std::unique_ptr<BatchClientRequest> batch_request = std::make_unique<BatchClientRequest>();
				for(int j = 0; j < 10; ++j){
				auto req = batch_request->add_client_requests();
				req->mutable_request()->set_data(std::to_string(rand()%10000));
				req->mutable_request()->set_seq(j+seq);
				}
				std::string data;
				batch_request->SerializeToString(&data);

				std::unique_ptr<BatchClientTransactions> batch_transactions =
				std::make_unique<BatchClientTransactions>();
				for (uint32_t i = 0; i < batch_num; ++i) {
				std::unique_ptr<ClientTransactions> client_txn = std::make_unique<ClientTransactions>();
				client_txn->set_transaction_data(data);
				client_txn->set_seq(seq+i);

				*batch_transactions->add_transactions() = *client_txn;
				}
				batch_transactions->set_min_seq(seq);
				batch_transactions->set_max_seq(seq + batch_num - 1);
				return batch_transactions;
	}));
	EXPECT_CALL(pow_manager_->mock_replica_client_, BroadCast).WillOnce(Invoke([&](const google::protobuf::Message& message){
				const Request *request = (const Request *)&message;
				EXPECT_EQ(request->type(), PoWRequest::TYPE_SHIFT_MSG);
				return absl::OkStatus();
		}));
	int call_time = 0;
	EXPECT_CALL(*pow_manager_->mock_block_manager_, Mine).Times(2).WillRepeatedly(Invoke([&](){
				LOG(ERROR)<<"mine time:"<<call_time<<" "<<pow_manager_->mock_block_manager_->GetSliceIdx();
				EXPECT_EQ(pow_manager_->mock_block_manager_->GetSliceIdx(), call_time);
				call_time++;
				if(call_time>1){
					done.set_value(true);
					return absl::OkStatus();
				}
			  return absl::NotFoundError("No new transaction.");
			}));
	pow_manager_->Start();
	done_future.get();
}

TEST_F(PoWManagerTest, ReSendShift) {
        config_.SetMaxNonceBit(5);
	config_.SetDifficulty(9);
	config_.SetMiningTime(1000);
	pow_manager_= std::make_unique<MockPoWManagerWithBC>(config_);
	pow_manager_->Reset();
	std::promise<bool> done;
	std::future<bool> done_future = done.get_future();
	ASSERT_TRUE(pow_manager_->mock_transaction_accessor_!=nullptr);

	EXPECT_CALL(*pow_manager_->mock_shift_manager_, Check).WillRepeatedly(Invoke([&](const SliceInfo &info, int){
				return false;
				}));
	EXPECT_CALL(*pow_manager_->mock_transaction_accessor_, ConsumeTransactions(_)).WillRepeatedly(Invoke([&](uint64_t seq){
				if(seq>1) {
				return (std::unique_ptr<XXXX::BatchClientTransactions>)nullptr;
				}
				int batch_num = 1;
				std::unique_ptr<BatchClientRequest> batch_request = std::make_unique<BatchClientRequest>();
				for(int j = 0; j < 10; ++j){
				auto req = batch_request->add_client_requests();
				req->mutable_request()->set_data(std::to_string(rand()%10000));
				req->mutable_request()->set_seq(j+seq);
				}
				std::string data;
				batch_request->SerializeToString(&data);

				std::unique_ptr<BatchClientTransactions> batch_transactions =
				std::make_unique<BatchClientTransactions>();
				for (uint32_t i = 0; i < batch_num; ++i) {
				std::unique_ptr<ClientTransactions> client_txn = std::make_unique<ClientTransactions>();
				client_txn->set_transaction_data(data);
				client_txn->set_seq(seq+i);

				*batch_transactions->add_transactions() = *client_txn;
				}
				batch_transactions->set_min_seq(seq);
				batch_transactions->set_max_seq(seq + batch_num - 1);
				return batch_transactions;
	}));

	int call_time = 0;
	EXPECT_CALL(pow_manager_->mock_replica_client_, BroadCast).WillRepeatedly(Invoke([&](const google::protobuf::Message& message){
				const Request *request = (const Request *)&message;
				EXPECT_EQ(request->type(), PoWRequest::TYPE_SHIFT_MSG);
				LOG(ERROR)<<"call broad cast:"<<call_time;
				call_time++;
				if(call_time>1){
					done.set_value(true);
					return absl::OkStatus();
				}
				return absl::OkStatus();
		}));
	
	EXPECT_CALL(*pow_manager_->mock_block_manager_, Mine).Times(1).WillRepeatedly(Invoke([&](){
			  return absl::NotFoundError("No new transaction.");
			}));

	pow_manager_->Start();
	done_future.get();
}

}  // namespace
}  // namespace XXXX
