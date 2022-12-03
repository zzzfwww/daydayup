# 配置要求
对于 Kubernetes 初学者，在搭建K8S集群时，推荐在阿里云或腾讯云采购如下配置：（您也可以使用自己的虚拟机、私有云等您最容易获得的 Linux 环境）

至少2台 2核2G 的服务器
Cent OS  7.8

安装后的软件版本为

- Kubernetes v1.19.x
  - calico 3.13.1
  - nginx-ingress 1.5.5
- Docker 19.03.11

关于二进制安装

> kubeadm 是 Kubernetes 官方支持的安装方式，“二进制” 不是。本文档采用 kubernetes.io 官方推荐的 kubeadm 工具安装 kubernetes 集群。

# 检查 centos / hostname

```bash
# 在 master 节点和 worker 节点都要执行
cat /etc/redhat-release

# 此处 hostname 的输出将会是该机器在 Kubernetes 集群中的节点名字
# 不能使用 localhost 作为节点的名字
hostname

# 请使用 lscpu 命令，核对 CPU 信息
# Architecture: x86_64    本安装文档不支持 arm 架构
# CPU(s):       2         CPU 内核数量不能低于 2
lscpu
```

修改 hostname

> 如果您需要修改 hostname，可执行如下指令：
> ```bash
> # 修改 hostname
> hostnamectl set-hostname k8smaster
> hostnamectl set-hostname k8sworker
> # 查看修改结果
> hostnamectl status
> # 设置 hostname 解析
> echo "127.0.0.1   $(hostname)" >> /etc/hosts
> ```

# 检查网络
```bash
[root@k8smaster ~]# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
2: ens33: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 00:0c:29:3d:87:32 brd ff:ff:ff:ff:ff:ff
    inet 192.168.3.201/24 brd 192.168.3.255 scope global noprefixroute ens33
       valid_lft forever preferred_lft forever
```

> **kubelet使用的IP地址**

> ip route show 命令中，可以知道机器的默认网卡，通常是 eth0或者ens33，如 default via 172.21.0.23 dev eth0

> ip address 命令中，可显示默认网卡的 IP 地址，Kubernetes 将使用此 IP 地址与集群内的其他节点通信，如 172.17.216.80

> 所有节点上 Kubernetes 所使用的 IP 地址必须可以互通（无需 NAT 映射、无安全组或防火墙隔离）

# 安装软件
使用 root 身份在所有节点执行如下代码，以安装软件：

docker
nfs-utils
kubectl / kubeadm / kubelet

docker hub 镜像请根据自己网络的情况任选一个

第四行为腾讯云 docker hub 镜像

第六行为DaoCloud docker hub 镜像

第八行为华为云 docker hub 镜像

第十行为阿里云 docker hub 镜像
```bash
# 在 master 节点和 worker 节点都要执行
# 最后一个参数 1.19.5 用于指定 kubenetes 版本，支持所有 1.19.x 版本的安装
# 腾讯云 docker hub 镜像
# export REGISTRY_MIRROR="https://mirror.ccs.tencentyun.com"
# DaoCloud 镜像
# export REGISTRY_MIRROR="http://f1361db2.m.daocloud.io"
# 华为云镜像
# export REGISTRY_MIRROR="https://05f073ad3c0010ea0f4bc00b7105ec20.mirror.swr.myhuaweicloud.com"
# 阿里云 docker hub 镜像
export REGISTRY_MIRROR=https://registry.cn-hangzhou.aliyuncs.com
curl -sSL https://kuboard.cn/install-script/v1.19.x/install_kubelet.sh | sh -s 1.19.5
```

# 初始化 master 节点

> 关于初始化时用到的环境变量

> APISERVER_NAME 不能是 master 的 hostname
> APISERVER_NAME 必须全为小写字母、数字、小数点，不能包含减号
> POD_SUBNET 所使用的网段不能与 master节点/worker节点 所在的网段重叠。该字段的取值为一个 CIDR 值，如果您对 CIDR 这个概念还不熟悉，请仍然执行 export POD_SUBNET=10.100.0.1/16 命令，不做修改

请将脚本最后的 1.19.5 替换成您需要的版本号， 脚本中间的 v1.19.x 不要替换
```bash
# 只在 master 节点执行
# 替换 x.x.x.x 为 master 节点实际 IP（请使用内网 IP）
# export 命令只在当前 shell 会话中有效，开启新的 shell 窗口后，如果要继续安装过程，请重新执行此处的 export 命令
export MASTER_IP=x.x.x.x
# 替换 apiserver.demo 为 您想要的 dnsName
export APISERVER_NAME=apiserver.demo
# Kubernetes 容器组所在的网段，该网段安装完成后，由 kubernetes 创建，事先并不存在于您的物理网络中
export POD_SUBNET=10.100.0.1/16
echo "${MASTER_IP}    ${APISERVER_NAME}" >> /etc/hosts
curl -sSL https://kuboard.cn/install-script/v1.19.x/init_master.sh | sh -s 1.19.5
```
## 检查 master 初始化结果
```bash
# 只在 master 节点执行

# 执行如下命令，等待 3-10 分钟，直到所有的容器组处于 Running 状态
watch kubectl get pod -n kube-system -o wide

# 查看 master 节点初始化结果
kubectl get nodes -o wide
```

# 初始化 worker节点
## 获得 join命令参数

在 master 节点上执行

```bash
# 只在 master 节点执行
kubeadm token create --print-join-command

# kubeadm token create 命令的输出
kubeadm join apiserver.demo:6443 --token mpfjma.4vjjg8flqihor4vt     --discovery-token-ca-cert-hash sha256:6f7a8e40a810323672de5eee6f4d19aa2dbdb38411845a1bf5dd63485c43d303

```

## 初始化worker
### 针对所有的 worker 节点执行

```bash
# 只在 worker 节点执行
# 替换 x.x.x.x 为 master 节点的内网 IP
export MASTER_IP=x.x.x.x
# 替换 apiserver.demo 为初始化 master 节点时所使用的 APISERVER_NAME
export APISERVER_NAME=apiserver.demo
echo "${MASTER_IP}    ${APISERVER_NAME}" >> /etc/hosts

# 替换为 master 节点上 kubeadm token create 命令的输出
kubeadm join apiserver.demo:6443 --token mpfjma.4vjjg8flqihor4vt     --discovery-token-ca-cert-hash sha256:6f7a8e40a810323672de5eee6f4d19aa2dbdb38411845a1bf5dd63485c43d303
```
# 检查初始化结果
```bash
[root@k8smaster ~]# kubectl get pod -n kube-system -o wide
NAME                                       READY   STATUS    RESTARTS   AGE   IP              NODE        NOMINATED NODE   READINESS GATES
calico-kube-controllers-6c89d944d5-mpmtb   1/1     Running   0          29m   10.100.16.129   k8smaster   <none>           <none>
calico-node-g7hnp                          1/1     Running   0          29m   192.168.3.201   k8smaster   <none>           <none>
calico-node-lkbjd                          1/1     Running   0          18m   192.168.3.202   k8sworker   <none>           <none>
coredns-59c898cd69-vwjmx                   1/1     Running   0          29m   10.100.16.131   k8smaster   <none>           <none>
coredns-59c898cd69-xsdzj                   1/1     Running   0          29m   10.100.16.130   k8smaster   <none>           <none>
etcd-k8smaster                             1/1     Running   0          29m   192.168.3.201   k8smaster   <none>           <none>
kube-apiserver-k8smaster                   1/1     Running   0          29m   192.168.3.201   k8smaster   <none>           <none>
kube-controller-manager-k8smaster          1/1     Running   0          29m   192.168.3.201   k8smaster   <none>           <none>
kube-proxy-7b2rn                           1/1     Running   0          18m   192.168.3.202   k8sworker   <none>           <none>
kube-proxy-wgk4d                           1/1     Running   0          29m   192.168.3.201   k8smaster   <none>           <none>
kube-scheduler-k8smaster                   1/1     Running   0          29m   192.168.3.201   k8smaster   <none>           <none>

[root@k8smaster ~]# kubectl get nodes -o wide
NAME        STATUS   ROLES    AGE   VERSION   INTERNAL-IP     EXTERNAL-IP   OS-IMAGE                KERNEL-VERSION           CONTAINER-RUNTIME
k8smaster   Ready    master   30m   v1.19.5   192.168.3.201   <none>        CentOS Linux 7 (Core)   3.10.0-1127.el7.x86_64   docker://19.3.11
k8sworker   Ready    <none>   18m   v1.19.5   192.168.3.202   <none>        CentOS Linux 7 (Core)   3.10.0-1127.el7.x86_64   docker://19.3.11
```

