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
yum -y install docker-ce
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