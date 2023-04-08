killall -9 mining_server

SERVER_PATH=./bazel-bin/application/poc/server/mining_server
TRANSACTION_SERVER_CONFIG=application/poc/server/transaction_server.config
MINING_SERVER_CONFIG=application/poc/server/mining_config.config
WORK_PATH=$PWD

bazel build //application/poc/server:mining_server $@
nohup $SERVER_PATH $TRANSACTION_SERVER_CONFIG $MINING_SERVER_CONFIG $WORK_PATH/cert/node6.key.pri $WORK_PATH/cert/cert_6.cert > server6.log &
nohup $SERVER_PATH $TRANSACTION_SERVER_CONFIG $MINING_SERVER_CONFIG $WORK_PATH/cert/node7.key.pri $WORK_PATH/cert/cert_7.cert > server7.log &
nohup $SERVER_PATH $TRANSACTION_SERVER_CONFIG $MINING_SERVER_CONFIG $WORK_PATH/cert/node8.key.pri $WORK_PATH/cert/cert_8.cert > server8.log &
nohup $SERVER_PATH $TRANSACTION_SERVER_CONFIG $MINING_SERVER_CONFIG $WORK_PATH/cert/node9.key.pri $WORK_PATH/cert/cert_9.cert > server9.log &

