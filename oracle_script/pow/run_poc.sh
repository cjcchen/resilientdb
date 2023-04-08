
CONFIG=$1

bazel build //application/poc/server:mining_server
bazel run //oracle_script/pow/script:run_svr $CONFIG

