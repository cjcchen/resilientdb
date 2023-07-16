package(default_visibility = ["//visibility:public"])

cc_library(
    name = "xxxx_config",
    srcs = ["xxxx_config.cpp"],
    hdrs = ["xxxx_config.h"],
    deps = [
        "//common:comm",
        "//proto:replica_info_cc_proto",
    ],
)

cc_test(
    name = "xxxx_config_test",
    srcs = ["xxxx_config_test.cpp"],
    deps = [
        ":xxxx_config",
        "//common/test:test_main",
    ],
)

cc_library(
    name = "xxxx_config_utils",
    srcs = ["xxxx_config_utils.cpp"],
    hdrs = ["xxxx_config_utils.h"],
    deps = [
        ":xxxx_config",
    ],
)

cc_library(
    name = "xxxx_poc_config",
    srcs = ["xxxx_poc_config.cpp"],
    hdrs = ["xxxx_poc_config.h"],
    deps = [
        ":xxxx_config",
    ],
)
