# k8s本地虚拟机部署

1. 设置host
```bash
172.21.51.67
172.21.51.68
172.21.51.69
# 在192.168.3.8执行
hostnamectl set-hostname k8s-master
# 在192.168.3.7执行
hostnamectl set-hostname  k8s-slave1
# 在192.168.3.18执行
hostnamectl set-hostname k8s-slave2
```
配置host
```bash
cat >> /etc/hosts << EOF
192.168.3.8 k8s-master
192.168.3.7 k8s-slave1
192.168.3.18 k8s-slave2
EOF
```
2. 配置ssh互信
```bash
# 直接一直回车就行
ssh-keygen

ssh-copy-id -i ~/.ssh/id_rsa.pub root@k8s-master
ssh-copy-id -i ~/.ssh/id_rsa.pub root@k8s-slave1
ssh-copy-id -i ~/.ssh/id_rsa.pub root@k8s-slave2
```
3. 同步时间
```bash
yum install chrony -y
systemctl start chronyd
systemctl enable chronyd
chronyc sources
```
4. 关闭防火墙
```bash
iptables -P FORWARD ACCEPT
systemctl stop firewalld && systemctl disable firewalld
```
5. 关闭swap
```bash
# 临时关闭
swapoff -a
# 可以通过这个命令查看swap是否关闭了
free
# 永久关闭
sed -ri 's/.*swap.*/#&/' /etc/fstab
```
6. 禁用SELinux
```bash
# 临时关闭
setenforce 0
# 永久禁用
sed -i 's/^SELINUX=enforcing$/SELINUX=disabled/' /etc/selinux/config
```
7. 修改内核参数
```
cat <<EOF> /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables=1
net.bridge.bridge-nf-call-iptables=1
net.ipv4.ip_forward=1
vm.max_map_count=262144
EOF
modprobe br_netfilter
sysctl -p /etc/sysctl.d/k8s.conf
```

8. 配置k8s yum源 所有节点
```bash
cat > /etc/yum.repos.d/kubernetes.repo << EOF
[kubernetes]
name=kubernetes
enabled=1
gpgcheck=0
repo_gogcheck=0
baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64/
gogkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
    http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF

yum clean all && yum makecache
```
9. 安装docker
```bash
# 制定版本安装
[root@k8s-master ~]# yum install docker-ce-cli-18.09.9-3.el7 docker-ce-18.09.9-3.el7

```

vi /etc/docker/daemon.json
```json
{
  "insecure-registries":[
    "192.168.3.8:5000"
  ],
  "registry-mirrors":[
    "https://8xpkSwnt.mirror.aliyuncs.com"
  ]
}
```
10. 安装kubeadm kubelet kubectl
```bash
yum install -y kubelet-1.16.2 kubeadm-1.16.2 kubectl-1.16.2

kubeadm version

systemctl enable kubelet
```
11. 初始化配置文件

操作节点：只在master节点执行
```bash
kubeadm config print init-defaults > kubeadm.yaml
```
修改配置文件
```yaml
[root@k8s-master ~]# cat kubeadm.yaml
apiVersion: kubeadm.k8s.io/v1beta2
bootstrapTokens:
- groups:
  - system:bootstrappers:kubeadm:default-node-token
  token: abcdef.0123456789abcdef
  ttl: 24h0m0s
  usages:
  - signing
  - authentication
kind: InitConfiguration
localAPIEndpoint:
  advertiseAddress: 192.168.3.8 # apiserver 地址，因为单master，所以配置master节点内网ip
  bindPort: 6443
nodeRegistration:
  criSocket: /var/run/dockershim.sock
  name: k8s-master
  taints:
  - effect: NoSchedule
    key: node-role.kubernetes.io/master
---
apiServer:
  timeoutForControlPlane: 4m0s
apiVersion: kubeadm.k8s.io/v1beta2
certificatesDir: /etc/kubernetes/pki
clusterName: kubernetes
controllerManager: {}
dns:
  type: CoreDNS
etcd:
  local:
    dataDir: /var/lib/etcd
imageRepository: registry.aliyuncs.com/google_containers # 修改镜像源
kind: ClusterConfiguration
kubernetesVersion: v1.16.2
networking:
  dnsDomain: cluster.local
  podSubnet: 10.244.0.0/16  # pod 网段，flannel插件需要使用这个网段
  serviceSubnet: 10.96.0.0/12
scheduler: {}
```
查看镜像列表
```bash
[root@k8s-master ~]# kubeadm config images list --config kubeadm.yaml
registry.aliyuncs.com/google_containers/kube-apiserver:v1.16.2
registry.aliyuncs.com/google_containers/kube-controller-manager:v1.16.2
registry.aliyuncs.com/google_containers/kube-scheduler:v1.16.2
registry.aliyuncs.com/google_containers/kube-proxy:v1.16.2
registry.aliyuncs.com/google_containers/pause:3.1
registry.aliyuncs.com/google_containers/etcd:3.3.15-0
registry.aliyuncs.com/google_containers/coredns:1.6.2

# 拉取镜像
[root@k8s-master ~]# kubeadm config images pull --config kubeadm.yaml
[config/images] Pulled registry.aliyuncs.com/google_containers/kube-apiserver:v1.16.2
[config/images] Pulled registry.aliyuncs.com/google_containers/kube-controller-manager:v1.16.2
[config/images] Pulled registry.aliyuncs.com/google_containers/kube-scheduler:v1.16.2
[config/images] Pulled registry.aliyuncs.com/google_containers/kube-proxy:v1.16.2
[config/images] Pulled registry.aliyuncs.com/google_containers/pause:3.1
[config/images] Pulled registry.aliyuncs.com/google_containers/etcd:3.3.15-0
[config/images] Pulled registry.aliyuncs.com/google_containers/coredns:1.6.2
```
12. 启动kubeadm
```bash
# 如果系统先前已经不纯粹，需要删除这些内容
rm -fr ~/.kube/  /etc/kubernetes/* var/lib/etcd/*
```