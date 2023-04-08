
CONFIG_PATH=$1

bazel build //application/poc/server:transaction_server
bazel run //oracle_script/pow/pbft/script:run_svr $CONFIG_PATH
