package(default_visibility = ["//visibility:public"])

cc_library(
    name = "xxxx_client",
    srcs = ["xxxx_client.cpp"],
    hdrs = ["xxxx_client.h"],
    deps = [
        "//common:comm",
        "//common/data_comm",
        "//common/network:tcp_socket",
        "//config:xxxx_config",
        "//crypto:signature_verifier",
        "//proto:xxxx_cc_proto",
    ],
)

cc_library(
    name = "mock_xxxx_client",
    testonly = True,
    hdrs = ["mock_xxxx_client.h"],
    deps = [
        ":xxxx_client",
    ],
)

cc_test(
    name = "xxxx_client_test",
    srcs = ["xxxx_client_test.cpp"],
    deps = [
        ":xxxx_client",
        "//common/network:mock_socket",
        "//common/test:test_main",
        "//crypto:key_generator",
        "//proto:client_test_cc_proto",
    ],
)

cc_library(
    name = "xxxx_user_client",
    srcs = ["xxxx_user_client.cpp"],
    hdrs = ["xxxx_user_client.h"],
    deps = [
        ":xxxx_client",
        "//common/data_comm",
    ],
)

cc_test(
    name = "xxxx_user_client_test",
    srcs = ["xxxx_user_client_test.cpp"],
    deps = [
        ":xxxx_user_client",
        "//common/network:mock_socket",
        "//common/test:test_main",
        "//crypto:signature_verifier",
        "//proto:client_test_cc_proto",
    ],
)

cc_library(
    name = "xxxx_state_client",
    srcs = ["xxxx_state_client.cpp"],
    hdrs = ["xxxx_state_client.h"],
    deps = [
        ":xxxx_user_client",
        "//common:comm",
        "//proto:replica_info_cc_proto",
        "//proto:xxxx_cc_proto",
    ],
)

cc_test(
    name = "xxxx_state_client_test",
    srcs = ["xxxx_state_client_test.cpp"],
    deps = [
        ":mock_xxxx_client",
        ":xxxx_state_client",
        "//common/test:test_main",
    ],
)

cc_library(
    name = "xxxx_txn_client",
    srcs = ["xxxx_txn_client.cpp"],
    hdrs = ["xxxx_txn_client.h"],
    deps = [
        ":xxxx_client",
        "//common:comm",
        "//proto:replica_info_cc_proto",
        "//proto:xxxx_cc_proto",
    ],
)

cc_library(
    name = "mock_xxxx_txn_client",
    testonly = True,
    hdrs = ["mock_xxxx_txn_client.h"],
    deps = [
        ":xxxx_txn_client",
    ],
)

cc_test(
    name = "xxxx_txn_client_test",
    srcs = ["xxxx_txn_client_test.cpp"],
    deps = [
        ":mock_xxxx_client",
        ":xxxx_txn_client",
        "//common/test:test_main",
        "//config:xxxx_config_utils",
    ],
)
