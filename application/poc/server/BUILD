package(default_visibility = ["//visibility:public"])

cc_binary(
    name = "transaction_server",
    srcs = ["transaction_server.cpp"],
    deps = [
        "//application/utils:server_factory",
        "//config:resdb_config_utils",
	"//ordering/poc/pbft:transaction_consensor",
	"//ordering/poc/pbft:transaction_executor",
	"//ordering/poc/pbft:transaction_query",
    ],
)

cc_binary(
    name = "mining_server",
    srcs = ["mining_server.cpp"],
    deps = [
    "//application/utils:server_factory",
        "//config:resdb_config_utils",
        "//ordering/poc/pow:consensus_service_pow",
        "//server:resdb_server",
	"//statistic:stats"
    ],
)

