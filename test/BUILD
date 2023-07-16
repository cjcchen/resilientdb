package(default_visibility = ["//visibility:public"])

cc_test(
    name = "xxxx_test",
    srcs = ["xxxx_test.cpp"],
    data = [
        "//test/test_data",
    ],
    deps = [
        "//client:xxxx_user_client",
        "//common:comm",
        "//common/test:test_main",
        "//config:xxxx_config_utils",
        "//execution:transaction_executor_impl",
        "//ordering/pbft:consensus_service_pbft",
        "//server:xxxx_server",
        "//test/proto:xxxx_test_cc_proto",
    ],
)
