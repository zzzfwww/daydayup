<!-- vscode-markdown-toc -->
* 1. [yum源变更](#yum)
* 2. [安装docker](#docker)
* 3. [docker-compose安装](#docker-compose)

<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->

# Docker 安装

##  1. <a name='yum'></a>yum源变更
1、安装wget
```bash
yum install -y wget
```
2、下载CentOS 7的repo文件
```bash
wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
```
##  2. <a name='docker'></a>安装docker
```bash
# step 1: 安装必要的一些系统工具
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
# Step 2: 添加软件源信息
sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
# Step 3: 更新并安装 Docker-CE
yum makecache fast
# 选定版本安装
yum install -y docker-ce-20.10.12 docker-ce-cli-20.10.12
# Step 4: 开启Docker服务
systemctl enable docker
systemctl start docker
```
##  3. <a name='docker-compose'></a>docker-compose安装
```bash
curl -L https://github.com/docker/compose/releases/download/1.27.4/docker-compose-`uname -s `-`uname -m` > /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

[root@localhost ~]# docker-compose version
docker-compose version 1.27.4, build 40524192
docker-py version: 4.3.1
CPython version: 3.7.7
OpenSSL version: OpenSSL 1.1.0l  10 Sep 2019
```

## 安装gitlab
1. 关闭防火墙，拉取镜像
```bash
# 关闭防火墙
systemctl stop firewalld
# 生成目录
mkdir -p /usr/local/docker/gitlab_docker
# 拉取镜像
docker pull gitlab/gitlab-ce:latest

vim /etc/docker/daemon.json
{
  "registry-mirrors": ["https://registry.docker-cn.com"]
}

systemctl restart docker
```
2. 编写docker-compose文件
```yaml
version: "3.1"
services:
  gitlab:
    image: "gitlab/gitlab-ce:latest"
    container_name: gitlab
    restart: always
    environment:
      GITLAB_OMNIBUS_CONFIG: |
        external_url 'http://192.168.3.101:8929'
        gitlab_rails['gitlab_shell_ssh_port'] = 2224
    ports:
      - "8929:8929"
      - "2224:2224"
    volumes:
      - "./config:/etc/gitlab"
      - "./logs:/var/log/gitlab"
      - "./data:/var/opt/gitlab"
```

3. 启动gitlab

```bash
docker-compose up -d
```
4. 查看docker-compose 日志
```bash
docker-compose -f logs
```
5. 登录gitlab
```bash
192.168.3.101 root  password

[root@localhost gitlab_docker]# docker exec -it gitlab bash
root@d7a0162df5ee:/# cat /etc/gitlab/initial_root_password
# WARNING: This value is valid only in the following conditions
#          1. If provided manually (either via `GITLAB_ROOT_PASSWORD` environment variable or via `gitlab_rails['initial_root_password']` setting in `gitlab.rb`, it was provided before database was seeded for the first time (usually, the first reconfigure run).
#          2. Password hasn't been changed manually, either via UI or via command line.
#
#          If the password shown here doesn't work, you must reset the admin password following https://docs.gitlab.com/ee/security/reset_user_password.html#reset-your-root-password.

Password: eyPdoT1d4eHJEqHReGKmoIoefcurJCiSrwEiCtieUfw=

# NOTE: This file will be automatically deleted in the first reconfigure run after 24 hours.
```
