# harbor

## harbor安装
1. 进入`https://goharbor.io/`点击`Download now`,跳转到github地址`https://github.com/goharbor/harbor/releases`，选择对应的版本下载。

2. 我选择的版本地址：
```text
https://objects.githubusercontent.com/github-production-release-asset-2e65be/50613991/10ea29a0-3b2f-4929-9e2c-323793be8752?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20221211%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20221211T042651Z&X-Amz-Expires=300&X-Amz-Signature=0138651dcc46af9dce3317180bdb3f6fb7221d721b4dd354be0ccfb785c8aaf6&X-Amz-SignedHeaders=host&actor_id=3351405&key_id=0&repo_id=50613991&response-content-disposition=attachment%3B%20filename%3Dharbor-offline-installer-v2.3.4.tgz&response-content-type=application%2Foctet-stream
```
3. 下载完成之后scp到目标机器上
```bash
scp harbor-offline-installer-v2.3.4.tgz root@192.168.3.102:h
arbor.tgz
```
4. 解压压缩包到/usr/local下
```bash
tar -zxvf harbor.tgz -C /usr/local

[root@jenkins harbor]# pwd
/usr/local/harbor
[root@jenkins harbor]# ls
common.sh             harbor.yml.tmpl  LICENSE
harbor.v2.3.4.tar.gz  install.sh       prepare
```

5. 修改配置文件安装harbor
```bash
[root@jenkins harbor]# vim harbor.yml
# 修改harbor.yml最终文件内容如下，主要设置hostname，以及把账号密码设置好
[root@jenkins harbor]# grep -Ev "^$|#" harbor.yml
hostname: 192.168.3.102
http:
  port: 80
harbor_admin_password: Harbor12345
database:
  password: root123
  max_idle_conns: 100
  max_open_conns: 900
data_volume: /data
trivy:
  ignore_unfixed: false
  skip_update: false
  insecure: false
jobservice:
  max_job_workers: 10
notification:
  webhook_job_max_retry: 10
chart:
  absolute_url: disabled
log:
  level: info
  local:
    rotate_count: 50
    rotate_size: 200M
    location: /var/log/harbor
_version: 2.3.0
proxy:
  http_proxy:
  https_proxy:
  no_proxy:
  components:
    - core
    - jobservice
    - trivy

# 执行安装脚本
[root@jenkins harbor]# ./install.sh

[Step 0]: checking if docker is installed ...

Note: docker version: 20.10.21

[Step 1]: checking docker-compose is installed ...

Note: docker-compose version: 1.27.4

[Step 2]: loading Harbor images ...
```
6. 启动成功，多了很多有关harbor的容器运行
```bash
[root@jenkins harbor]# docker ps
CONTAINER ID   IMAGE                                COMMAND                  CREATED             STATUS                   PORTS                                                                                      NAMES
a850ab989465   goharbor/nginx-photon:v2.3.4         "nginx -g 'daemon of…"   6 minutes ago       Up 6 minutes (healthy)   0.0.0.0:80->8080/tcp, :::80->8080/tcp                                                      nginx
8eaea927eadb   goharbor/harbor-jobservice:v2.3.4    "/harbor/entrypoint.…"   6 minutes ago       Up 6 minutes (healthy)                                                                                              harbor-jobservice
a3b13147ac37   goharbor/harbor-core:v2.3.4          "/harbor/entrypoint.…"   6 minutes ago       Up 6 minutes (healthy)                                                                                              harbor-core
0b45dc6cc3e2   goharbor/redis-photon:v2.3.4         "redis-server /etc/r…"   6 minutes ago       Up 6 minutes (healthy)                                                                                              redis
364aa11fa9e0   goharbor/harbor-db:v2.3.4            "/docker-entrypoint.…"   6 minutes ago       Up 6 minutes (healthy)                                                                                              harbor-db
c1866a6a5db9   goharbor/harbor-registryctl:v2.3.4   "/home/harbor/start.…"   6 minutes ago       Up 6 minutes (healthy)                                                                                              registryctl
bf880a2cedd0   goharbor/harbor-portal:v2.3.4        "nginx -g 'daemon of…"   6 minutes ago       Up 6 minutes (healthy)                                                                                              harbor-portal
a47c8ea21cc3   goharbor/registry-photon:v2.3.4      "/home/harbor/entryp…"   6 minutes ago       Up 6 minutes (healthy)
```

