# 网站被攻击了，自动封禁ip的脚本

## 备份Nginx访问日志
```sh
# 每天0点执行日志按日期分隔脚本 0 0 * * * cd /www/Home/ && ./log_cut.sh
#!/bin/bash
#此脚本⽤于⾃动分割Nginx的⽇志，包括access.log
#每天00:00执⾏此脚本将前⼀天的access.log重命名为access-xxxx-xx-xx.log格式，并重新打开⽇志⽂件
#Nginx⽇志⽂件所在⽬录 todo 换成你自己的
LOG_PATH=/data/logs/nginx/
#获取昨天的⽇期
YESTERDAY=$(date -d "yesterday" +%Y-%m-%d)
#获取pid⽂件路径 todo 换成你自己的
PID=/var/run/nginx.pid
#分割⽇志
mv ${LOG_PATH}access.log ${LOG_PATH}access-${YESTERDAY}.log
#向Nginx主进程发送USR1信号，重新打开⽇志⽂件
kill -USR1 `cat ${PID}`

# 每天0点执行日志按日期分隔脚本 
0 0 * * * cd /www/Home/ && ./log_cut.sh
```

## 编写封禁ip脚本
```sh
# 每十分钟执行一次封禁ip脚本 */10 * * * * cd /www/Home/ && ./blackip.sh
#!/bin/bash
logdir=/data/logs/nginx/access.log #nginx访问日志文件路径
port=443
#循环遍历日志文件取出访问量大于100的ip（忽略自己本地ip）
for drop_ip in $(cat $logdir | grep -v '127.0.0.1' | awk '{print $1}' | sort | uniq -c | sort -rn | awk '{if ($1>100) print $2}'); do
  # 避免重复添加
  num=$(grep ${drop_ip} /tmp/nginx_deny.log | wc -l)
  if [ $num -ge 1 ]; then
    continue
  fi
  # shellcheck disable=SC2154
  iptables -I INPUT -p tcp --dport ${port} -s ${drop_ip} -j DROP
  echo ">>>>> $(date '+%Y-%m-%d %H%M%S') - 发现攻击源地址 ->  ${drop_ip} " >>/tmp/nginx_deny.log #记录log
done
```

## 如果误操作了
```sh
#清空屏蔽IP
iptables -t filter -D INPUT -s 1.2.3.4 -j DROP

#一键清空所有规则
iptables -F
```

