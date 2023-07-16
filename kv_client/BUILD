package(default_visibility = ["//visibility:public"])

cc_library(
    name = "xxxx_kv_client",
    srcs = ["xxxx_kv_client.cpp"],
    hdrs = ["xxxx_kv_client.h"],
    deps = [
        "//client:xxxx_user_client",
        "//proto:kv_server_cc_proto",
    ],
)

cc_library(
    name = "xxxx_kv_performance_client",
    srcs = ["xxxx_kv_performance_client.cpp"],
    hdrs = ["xxxx_kv_performance_client.h"],
    deps = [
        ":xxxx_kv_client",
    ],
)

cc_binary(
    name = "kv_performance_client_main",
    srcs = ["kv_performance_client_main.cpp"],
    deps = [
        ":xxxx_kv_performance_client",
        "//config:xxxx_config_utils",
        "//proto:signature_info_cc_proto",
    ],
)
