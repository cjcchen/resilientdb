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

cc_binary(
    name = "kv_client_tools",
    srcs = ["kv_client_tools.cpp"],
    deps = [
        ":xxxx_kv_client",
        "//config:xxxx_config_utils",
    ],
)
