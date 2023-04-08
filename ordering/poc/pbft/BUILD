package(default_visibility = ["//visibility:public"])

cc_library(
    name = "transaction_consensor",
    srcs = ["transaction_consensor.cpp"],
    hdrs = ["transaction_consensor.h"],
    deps = [
    ":transaction_executor",
        "//ordering/pbft:consensus_service_pbft",
	"//ordering/poc/proto:pow_cc_proto",
	"//ordering/poc/pow:miner_utils",
    ],
)

cc_library(
    name = "transaction_manager",
    srcs = ["transaction_manager.cpp"],
    hdrs = ["transaction_manager.h"],
    deps = [
        "//execution:transaction_executor",
        "//execution:transaction_executor_impl",
      "//ordering/poc/proto:pow_cc_proto"
    ],
)

cc_library(
    name = "transaction_executor",
    srcs = ["transaction_executor.cpp"],
    hdrs = ["transaction_executor.h"],
    deps = [
    ":transaction_manager",
        "//execution:transaction_executor",
        "//execution:transaction_executor_impl",
    ],
)

cc_library(
    name = "transaction_query",
    srcs = ["transaction_query.cpp"],
    hdrs = ["transaction_query.h"],
    deps = [
      ":transaction_manager",
      "//execution:custom_query",
      "//ordering/poc/proto:pow_cc_proto"
    ],
)


