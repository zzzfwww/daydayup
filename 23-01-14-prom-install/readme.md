# Prometheus

安装Prometheus并且使用

## 同步时间
```shell
# 设置时区
timedatectl set-timezone Asia/Shanghai

# 更新时间
* * * * * ntpdate -u cn.pool.ntp.org
```

## 下载安装包
```shell
wget https://github.com/prometheus/prometheus/releases/download/v2.41.0/prometheus-2.41.0.linux-amd64.tar.gz
```

## 解压
```shell
tar -zxvf prometheus-2.41.0.linux-amd64.tar.gz

mv prometheus-2.41.0.linux-amd64 /usr/local/prometheus
```

## 使用screen后台启动服务
```shell

[root@prom prometheus]# ./prometheus
ts=2023-01-14T10:13:35.853Z caller=main.go:512 level=info msg="No time or size retention was set so using the default time retention" duration=15d
ts=2023-01-14T10:13:35.853Z caller=main.go:556 level=info msg="Starting Prometheus Server" mode=server version="(version=2.41.0, branch=HEAD, revision=c0d8a56c69014279464c0e15d8bfb0e153af0dab)"
ts=2023-01-14T10:13:35.853Z caller=main.go:561 level=info build_context="(go=go1.19.4, platform=linux/amd64, user=root@d20a03e77067, date=20221220-10:40:45)"
ts=2023-01-14T10:13:35.853Z caller=main.go:562 level=info host_details="(Linux 3.10.0-1127.el7.x86_64 #1 SMP Tue Mar 31 23:36:51 UTC 2020 x86_64 prom (none))"
ts=2023-01-14T10:13:35.853Z caller=main.go:563 level=info fd_limits="(soft=4096, hard=4096)"
ts=2023-01-14T10:13:35.853Z caller=main.go:564 level=info vm_limits="(soft=unlimited, hard=unlimited)"
ts=2023-01-14T10:13:35.857Z caller=web.go:559 level=info component=web msg="Start listening for connections" address=0.0.0.0:9090
ts=2023-01-14T10:13:35.859Z caller=main.go:993 level=info msg="Starting TSDB ..."
ts=2023-01-14T10:13:35.863Z caller=head.go:562 level=info component=tsdb msg="Replaying on-disk memory mappable chunks if any"
ts=2023-01-14T10:13:35.863Z caller=head.go:606 level=info component=tsdb msg="On-disk memory mappable chunks replay completed" duration=6.687µs
ts=2023-01-14T10:13:35.863Z caller=head.go:612 level=info component=tsdb msg="Replaying WAL, this may take a while"
ts=2023-01-14T10:13:35.864Z caller=tls_config.go:232 level=info component=web msg="Listening on" address=[::]:9090
ts=2023-01-14T10:13:35.864Z caller=tls_config.go:235 level=info component=web msg="TLS is disabled." http2=false address=[::]:9090
ts=2023-01-14T10:13:35.864Z caller=head.go:683 level=info component=tsdb msg="WAL segment loaded" segment=0 maxSegment=1
ts=2023-01-14T10:13:35.864Z caller=head.go:683 level=info component=tsdb msg="WAL segment loaded" segment=1 maxSegment=1
ts=2023-01-14T10:13:35.864Z caller=head.go:720 level=info component=tsdb msg="WAL replay completed" checkpoint_replay_duration=24.27µs wal_replay_duration=1.132416ms wbl_replay_duration=121ns total_replay_duration=1.182225ms
ts=2023-01-14T10:13:35.865Z caller=main.go:1014 level=info fs_type=XFS_SUPER_MAGIC
ts=2023-01-14T10:13:35.865Z caller=main.go:1017 level=info msg="TSDB started"
ts=2023-01-14T10:13:35.865Z caller=main.go:1197 level=info msg="Loading configuration file" filename=prometheus.yml
ts=2023-01-14T10:13:35.913Z caller=main.go:1234 level=info msg="Completed loading of configuration file" filename=prometheus.yml totalDuration=47.579942ms db_storage=3.341µs remote_storage=3.292µs web_handler=238ns query_engine=577ns scrape=47.244025ms scrape_sd=48.51µs notify=44.135µs notify_sd=12.513µs rules=5.638µs tracing=21.951µs
ts=2023-01-14T10:13:35.913Z caller=main.go:978 level=info msg="Server is ready to receive web requests."
ts=2023-01-14T10:13:35.913Z caller=manager.go:953 level=info component="rule manager" msg="Starting rule manager..."

[root@prom prometheus]# screen -ls
There is a screen on:
	1873.pts-0.prom	(Detached)
1 Socket in /var/run/screen/S-root.

[root@prom prometheus]# screen -r 1873
```

## 一般正式环境参数启动
```sh
[root@prometheus ~]# cat /data/prometheus/up.sh
/data/prometheus/prometheus --web.listen-address="0.0.0.0:9090" --web.read-timeout=5m --web.max-connections=10 --storage.tsdb.retention=15d --storage.tsdb.path="data/" --query.max-concurrency=20 --query.timeout=2m

# 一般使用daemonzie 启动

daemonzie -c /data/prometheus /data/prometheus/up.sh
``