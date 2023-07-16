package(default_visibility = ["//visibility:public"])

cc_library(
    name = "server_comm",
    hdrs = ["server_comm.h"],
    deps = [
        "//client:xxxx_client",
        "//proto:xxxx_cc_proto",
    ],
)

cc_library(
    name = "xxxx_service",
    srcs = ["xxxx_service.cpp"],
    hdrs = ["xxxx_service.h"],
    deps = [
        ":server_comm",
        "//common/data_comm",
        "//config:xxxx_config",
    ],
)

cc_library(
    name = "mock_xxxx_service",
    hdrs = ["mock_xxxx_service.h"],
    deps = [
        ":xxxx_service",
        "//common/test",
    ],
)

cc_library(
    name = "consensus_service",
    srcs = ["consensus_service.cpp"],
    hdrs = ["consensus_service.h"],
    deps = [
        ":xxxx_replica_client",
        ":xxxx_service",
        "//common:comm",
        "//common/queue:blocking_queue",
        "//config:xxxx_config",
        "//proto:broadcast_cc_proto",
        "//proto:xxxx_cc_proto",
        "//statistic:stats",
    ],
)

cc_test(
    name = "consensus_service_test",
    srcs = ["consensus_service_test.cpp"],
    deps = [
        ":consensus_service",
        ":mock_xxxx_replica_client",
        "//common/test:test_main",
    ],
)

cc_library(
    name = "async_acceptor",
    srcs = ["async_acceptor.cpp"],
    hdrs = ["async_acceptor.h"],
    deps = [
        "//common:asio",
        "//common:comm",
        "//config:xxxx_config",
    ],
)

cc_test(
    name = "async_acceptor_test",
    srcs = ["async_acceptor_test.cpp"],
    deps = [
        ":async_acceptor",
        "//common/network:tcp_socket",
        "//common/test:test_main",
    ],
)

cc_library(
    name = "xxxx_server",
    srcs = ["xxxx_server.cpp"],
    hdrs = ["xxxx_server.h"],
    deps = [
        ":async_acceptor",
        ":xxxx_service",
        "//common/data_comm",
        "//common/network:tcp_socket",
        "//common/queue:lock_free_queue",
        "//proto:broadcast_cc_proto",
        "//statistic:stats",
    ],
)

cc_test(
    name = "xxxx_server_test",
    srcs = ["xxxx_server_test.cpp"],
    deps = [
        ":mock_xxxx_service",
        ":xxxx_server",
        "//common/test:test_main",
    ],
)

cc_library(
    name = "async_replica_client",
    srcs = ["async_replica_client.cpp"],
    hdrs = ["async_replica_client.h"],
    deps = [
        "//client:xxxx_client",
        "//common:asio",
        "//common/queue:blocking_queue",
        "//common/queue:lock_free_queue",
        "//proto:broadcast_cc_proto",
    ],
)

cc_test(
    name = "async_replica_client_test",
    srcs = ["async_replica_client_test.cpp"],
    deps = [
        ":async_replica_client",
        "//client:mock_xxxx_client",
        "//common/network:mock_socket",
        "//common/network:tcp_socket",
        "//common/test:test_main",
        "@boost//:thread",
    ],
)

cc_library(
    name = "mock_async_replica_client",
    testonly = True,
    hdrs = ["mock_async_replica_client.h"],
    deps = [
        ":async_replica_client",
    ],
)

cc_library(
    name = "xxxx_replica_client",
    srcs = ["xxxx_replica_client.cpp"],
    hdrs = ["xxxx_replica_client.h"],
    deps = [
        ":async_replica_client",
        "//client:xxxx_client",
        "//common/queue:batch_queue",
        "//proto:broadcast_cc_proto",
        "//statistic:stats",
    ],
)

cc_library(
    name = "mock_xxxx_replica_client",
    testonly = True,
    hdrs = ["mock_xxxx_replica_client.h"],
    deps = [
        ":xxxx_replica_client",
    ],
)

cc_test(
    name = "xxxx_replica_client_test",
    srcs = ["xxxx_replica_client_test.cpp"],
    deps = [
        ":mock_async_replica_client",
        ":xxxx_replica_client",
        "//client:mock_xxxx_client",
        "//common/network:mock_socket",
        "//common/test:test_main",
    ],
)
