package(default_visibility = ["//visibility:public"])

cc_binary(
    name = "key_generator_tools",
    srcs = ["key_generator_tools.cpp"],
    deps = [
        "//crypto:key_generator",
    ],
)

cc_binary(
    name = "certificate_tools",
    srcs = ["certificate_tools.cpp"],
    deps = [
        "//crypto:key_generator",
        "//crypto:signature_verifier",
    ],
)

cc_binary(
    name = "certificate_tools_test",
    srcs = ["certificate_tools_test.cpp"],
    deps = [
        "//crypto:signature_verifier",
    ],
)

cc_binary(
    name = "xxxx_status_client_tools",
    srcs = ["xxxx_status_client_tools.cpp"],
    deps = [
        "//client:xxxx_state_client",
        "//config:xxxx_config_utils",
    ],
)

cc_binary(
    name = "xxxx_txn_client_tools",
    srcs = ["xxxx_txn_client_tools.cpp"],
    deps = [
        "//client:xxxx_txn_client",
        "//config:xxxx_config_utils",
    ],
)

cc_binary(
    name = "kv_client_txn_tools",
    srcs = ["kv_client_txn_tools.cpp"],
    deps = [
        "//client:xxxx_txn_client",
        "//config:xxxx_config_utils",
        "//proto:kv_server_cc_proto",
    ],
)

py_binary(
    name = "generate_region_config",
    srcs = ["generate_region_config.py"],
    deps = [
        "//proto:replica_info_py_proto",
    ],
)

py_binary(
    name = "generate_mulregion_config",
    srcs = ["generate_mulregion_config.py"],
    deps = [
        "//proto:replica_info_py_proto",
    ],
)
