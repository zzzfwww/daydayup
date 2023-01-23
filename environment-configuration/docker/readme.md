# docker

## 环境准备
```text
wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Cnetos-7.repo

wget -O /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-7.repo

yum clean all

yum makecache

yum install -y bash-completion vim lrzsz wgt expect net-tools nc nmap tree dos2unix htop iftop iotop unzip telnet sl psmisc nethogs glances bc ntpdate openIdap-devel
```
* 安装docker
* 开启Linux内核的流量转发
```conf
cat <<EOF > /etc/sysctl.d/docker.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.conf.default.rp_filter = 0
net.ipv4.conf.all.rp_filter = 0
net.ipv4.ip_forward = 1
EOF

sysctl -p /etc/sysctl.d/docker.conf

如果报错

sysctl cannot stat /proc/sys/net/bridge/bridge-nf-call-ip6tables: No such file or directory

执行 modprobe br_netfilter
```
* SELinux配置[link](https://help.aliyun.com/document_detail/157022.html)

## 利用yum快速安装docker

* 查看yum源里面是否有docker 包
```shell
yum list docker-ce --showduplicates |sort -r
```
* 配置yum源 docker-ce
```shell
curl -o /etc/yum.repos.d/docker-ce.repo http:mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
```
* 更新yum缓存：`yum clean && yum makecache`
* yum -y install docker-ce:20.10.21
* 安装完成之后 `docker version`查看版本
```text
[~]$ docker version
Client: Docker Engine - Community
 Version:           20.10.21
 API version:       1.41
 Go version:        go1.18.7
 Git commit:        baeda1f
 Built:             Tue Oct 25 18:04:24 2022
 OS/Arch:           linux/amd64
 Context:           default
 Experimental:      true

Server: Docker Engine - Community
 Engine:
  Version:          20.10.21
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.18.7
  Git commit:       3056208
  Built:            Tue Oct 25 18:02:38 2022
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.6.12
  GitCommit:        a05d175400b1145e5e6a735a6710579d181e7fb0
 runc:
  Version:          1.1.4
  GitCommit:        v1.1.4-0-g5fd4c4d
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0
```
* 配置docker镜像源
```text
[~]$ cat /etc/docker/daemon.json
{
  "registry-mirrors": ["https://registry.docker-cn.com"],
  "insecure-registries": ["192.168.3.102:80"]
}
# 或者地址
{
  "registry-mirrors": ["https://8xpk5wnt.mirror.aliyuncs.com"]
}
```
* 启动 重新启动， 开机启动
```shell
systemctl daemon-reload
systemctl enable docker
systemctl restart docker
```
