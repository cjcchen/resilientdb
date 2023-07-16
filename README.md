# XDB: Global-Scale Sustainable Blockchain Fabric

## OS Requirements
Ubuntu 20.*

---

## Build and Deploy XDB

Install dependencies:

    ./INSTALL.sh


Run XDB (Providing a Key-Value Service):

    ./example/start_kv_server.sh
    
- This script will start 4 replica and 1 client. Each replica instantiates a key-value store.

Build Interactive Tools:

    bazel build example/kv_server_tools

Run tools to set a value by a key (for example, set the value with key "test" and value "test_value"):

    bazel-bin/example/kv_server_tools example/kv_client_config.config set test test_value
    
You will see the following result if successful:

    client set ret = 0

Run tools to get value by a key (for example, get the value with key "test"):

    bazel-bin/example/kv_server_tools example/kv_client_config.config get test
    
You will see the following result if successful:

    client get value = test_value

Run tools to get all values that have been set:

    bazel-bin/example/kv_server_tools example/kv_client_config.config getvalues

You will see the following result if successful:

    client getvalues value = [test_value]

## Enable SDK

If you want to enable sdk verification, add a flag to the command

    ./example/start_kv_server.sh --define enable_sdk=True

And in example/kv_config.config

    require_txn_validation:true,

## FAQ

If installing bazel fails on a ubuntu server, follow these steps:

> mkdir bazel-src
>
> cd bazel-src
>
> wget https://github.com/bazelbuild/bazel/releases/download/5.4.0/bazel-5.4.0-dist.zip
>
> sudo apt-get install build-essential openjdk-11-jdk python zip unzip
>
> unzip bazel-5.4.0-dist.zip
>
> env EXTRA_BAZEL_ARGS="--tool_java_runtime_version=local_jdk" bash ./compile.sh
>
> sudo mv output/bazel /usr/local/bin
