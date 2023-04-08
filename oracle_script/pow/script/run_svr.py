#!/bin/sh

import sys

import json
from oracle_script.comm.utils import *
from oracle_script.comm.comm_config import *
from proto.replica_info_pb2 import ResConfigData,ReplicaInfo
from google.protobuf.json_format import MessageToJson
from google.protobuf.json_format import Parse, ParseDict

cert_path="pbft_cert/"

def gen_svr_config(config):
    iplist=get_ips(config["svr_ip_file"])

    info_list = []
    for idx,ip in iplist:
        port = int(config["base_port"]) + int(idx)
        info={}
        info["id"] = idx
        info["ip"] = ip
        info["port"] = port
        info_list.append(info)

    with open(config["pow_config_path"],"w") as f:
        config_data=ResConfigData() 
        region = config_data.region.add()
        region.region_id=1
        for info in info_list:
            replica = Parse(json.dumps(info), ReplicaInfo())
            region.replica_info.append(replica)

        config_data.self_region_id = 1
        config_data.is_performance_running = True
        config_data.max_process_txn = 1024
        config_data.client_batch_num = 200
        config_data.worker_num=16
        if "WORKER_NUM" in os.environ:
            print ("evn worker:",os.environ["WORKER_NUM"])
            config_data.worker_num=int(os.environ["WORKER_NUM"])
        config_data.input_worker_num = 2
        config_data.output_worker_num = 2
        json_obj = MessageToJson(config_data)
        f.write(json_obj)
        print("write to {} done".format(config["pow_config_path"]))

def kill_svr(config):
    iplist=get_ips(config["svr_ip_file"])+get_ips(config["cli_ip_file"])
    cmd_list=[]
    for (idx,svr_ip) in iplist:
        cmd="killall -9 {};".format(config["svr_bin"])
        cmd_list.append((svr_ip,cmd))
        run_remote_cmd(svr_ip, cmd)
    #run_remote_cmd_list(cmd_list)

if __name__ == '__main__':
    config_file=sys.argv[1]
    config = read_config(config_file)
    print("config:{}".format(config))
    gen_svr_config(config)
