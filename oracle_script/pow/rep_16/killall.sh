
iplist=`cat iplist.txt`
key=~/.ssh/junchao.pem
dir=$PWD

count=1
for ip in ${iplist[@]};
do
	echo ${ip}
	ssh -i ${key} -n -o BatchMode=yes -o StrictHostKeyChecking=no ubuntu@${ip} " killall -9 mining_server; " &
	((count++))
done

while [ $count -gt 0 ]; do
wait $pids
count=`expr $count - 1`
done

echo "================ kill done ======"

count=1
for ip in ${iplist[@]};
do
	ssh -i ${key} -n -o BatchMode=yes -o StrictHostKeyChecking=no ubuntu@${ip} " rm -rf /home/ubuntu/mining_server; rm server*.log; rm -rf server.config; rm -rf pow_server.config; rm -rf cert; mkdir -p pbft_cert/; " &
	((count++))
done

while [ $count -gt 0 ]; do
wait $pids
count=`expr $count - 1`
done

count=1
idx=1
for ip in ${iplist[@]};
do
scp -i ${key} /home/ubuntu/resilientdb/bazel-bin/application/poc/server/mining_server ubuntu@${ip}:/home/ubuntu &
scp -i ${key} ${dir}/server.config ubuntu@${ip}:/home/ubuntu/mining_server.config &
scp -i ${key} ${dir}/../pbft/rep_16/server.config ubuntu@${ip}:/home/ubuntu/transaction_server.config &
scp -i ${key} ${dir}/cert/node_${idx}.key.pri ubuntu@${ip}:/home/ubuntu/pbft_cert/ &
scp -i ${key} ${dir}/cert/cert_${idx}.cert ubuntu@${ip}:/home/ubuntu/pbft_cert/ &
	((count++))
	((count++))
	((count++))
	((count++))
	((idx++))
done

while [ $count -gt 0 ]; do
wait $pids
count=`expr $count - 1`
done

echo "================ rm done ======"

idx=1
count=1
for ip in ${iplist[@]};
do
	ssh -i ${key} -n -o BatchMode=yes -o StrictHostKeyChecking=no ubuntu@${ip} " nohup /home/ubuntu/mining_server /home/ubuntu/transaction_server.config /home/ubuntu/mining_server.config pbft_cert//node_${idx}.key.pri pbft_cert//cert_${idx}.cert > server${idx}.log 2>&1 & " &
	((count++))
	((idx++))
done

while [ $count -gt 0 ]; do
wait $pids
count=`expr $count - 1`
done

echo "================ start done ======"
