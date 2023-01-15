#!/bin/bash

instance_name=`hostname -f |cut -d'.' -f1`

if [ $instance_name == "localhost" ];then
	echo "mush FQDN hostname"
	exit 1
fi

# For waiting connections

label="count_netstat_wait_connections"
count_netstat_wait_connections=`netstat -an |grep -i wait |wc -l `

echo "$label:$count_netstat_wait_connections"
echo "$label $count_netstat_wait_connections" |curl --data-binary @- http://prometheus.server.com:9091/metrics/job/pushgateway1/instance/$instance_name
