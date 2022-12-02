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
yum install -y kubelet-1.16.2 kubeadm-1.16.2 kubectl-1.16.2 --disableexcludes=kubernetes

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

kubeadm init --config kubeadm.yaml
[kubelet-check] Initial timeout of 40s passed.
# 出现报错
error execution phase upload-config/kubelet: Error writing Crisocket information for the control-plane node: timed out waiting for the condition
To see the stack trace of this error execute with --v=5 or higher
# 解决办法
[root@k8s-master ~]# swapoff -a && kubeadm reset  && systemctl daemon-reload && systemctl restart kubelet  && iptables -F && iptables -t nat -F && iptables -t mangle -F && iptables -X
[reset] Reading configuration from the cluster...
[root@k8s-master ~]# rm -rf ~/.kube/  /etc/kubernetes/* /var/lib/etcd/*
[root@k8s-master ~]# kubeadm init --config=kubeadm.yaml

## 最终成功启动
[bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstrap-token] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.3.8:6443 --token abcdef.0123456789abcdef \
    --discovery-token-ca-cert-hash sha256:ba0f12e1c8b615b74b1127511f83839a865b5ad3ee48aa80645695176c32fffe
```

13. 查看信息
```bash
[root@k8s-master ~]# kubectl get no
NAME         STATUS     ROLES    AGE   VERSION
k8s-master   NotReady   master   42m   v1.16.2
```
> NotReady 状态没成功是因为还未完成网络配置

14. 子节点加入主节点
```bash

kubeadm join 192.168.3.8:6443 --token abcdef.0123456789abcdef \
    --discovery-token-ca-cert-hash sha256:ba0f12e1c8b615b74b1127511f83839a865b5ad3ee48aa80645695176c32fffe

# 再次查看集群信息
[root@k8s-master ~]# kubectl get no
NAME         STATUS     ROLES    AGE   VERSION
k8s-master   NotReady   master   48m   v1.16.2
k8s-slave1   NotReady   <none>   34s   v1.16.2
k8s-slave2   NotReady   <none>   93s   v1.16.2
```

15. 安装flannel插件
```bash
# 一般会被网络墙
wget https://raw.githubusercontent.com/coreos/flannel/2140ac876ef134e0ed5af15c65e414cf26827915/Documentation/kube-flannel.yml
```
使用[flannel.yml](./kube-flannel.yml)
```
[root@k8s-slave2 ~]# ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: ens33: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000

修改flannel.yml 大概190行 增加
- --iface=ens33
# 出现报错
[root@k8s-master ~]# kubectl delete -f kube-flannel.yml
podsecuritypolicy.policy "psp.flannel.unprivileged" deleted
clusterrole.rbac.authorization.k8s.io "flannel" deleted
clusterrolebinding.rbac.authorization.k8s.io "flannel" deleted
serviceaccount "flannel" deleted
configmap "kube-flannel-cfg" deleted
error: error parsing kube-flannel.yml: error converting YAML to JSON: yaml: line 59: found a tab character that violates indentation

# 安装yaml检测工具
yum -y install yamllint

[root@k8s-master ~]# yamllint kube-flannel.yml
kube-flannel.yml
  42:3      error    wrong indentation: expected 4 but found 2  (indentation)
  87:1      error    wrong indentation: expected 2 but found 0  (indentation)
  167:7     error    wrong indentation: expected 8 but found 6  (indentation)
  171:7     error    wrong indentation: expected 8 but found 6  (indentation)
  174:9     error    wrong indentation: expected 10 but found 8  (indentation)
  176:9     error    wrong indentation: expected 10 but found 8  (indentation)
  180:9     error    wrong indentation: expected 10 but found 8  (indentation)
  185:7     error    wrong indentation: expected 8 but found 6  (indentation)
  188:9     error    wrong indentation: expected 10 but found 8  (indentation)
  190:9     error    wrong indentation: expected 10 but found 8  (indentation)
  192:1     error    syntax error: found character '\t' that cannot start any token (syntax)
# 出错的地方在天剑的网卡信息，这里一定不能用tab键，直接四个空格完成配置，要不然会报错

# 再次执行，无报错，成功
[root@k8s-master ~]# kubectl create -f kube-flannel.yml
podsecuritypolicy.policy/psp.flannel.unprivileged created
clusterrole.rbac.authorization.k8s.io/flannel created
clusterrolebinding.rbac.authorization.k8s.io/flannel created
serviceaccount/flannel created
configmap/kube-flannel-cfg created
daemonset.apps/kube-flannel-ds-amd64 created
daemonset.apps/kube-flannel-ds-arm64 created
daemonset.apps/kube-flannel-ds-arm created
daemonset.apps/kube-flannel-ds-ppc64le created
daemonset.apps/kube-flannel-ds-s390x created
```

16. 查看状态信息
```bash
[root@k8s-master ~]# kubectl -n kube-system get  po
NAME                                 READY   STATUS              RESTARTS   AGE
coredns-58cc8c89f4-b56hl             0/1     Pending             0          117m
coredns-58cc8c89f4-m8rpm             0/1     Pending             0          117m
etcd-k8s-master                      1/1     Running             0          116m
kube-apiserver-k8s-master            1/1     Running             0          116m
kube-controller-manager-k8s-master   1/1     Running             0          116m
kube-flannel-ds-amd64-2l5mb          1/1     Running             0          5m42s
kube-flannel-ds-amd64-gfm8l          1/1     Running             0          5m42s
kube-flannel-ds-amd64-jld4w          0/1     Init:0/1            0          5m42s
kube-proxy-59zqw                     1/1     Running             0          70m
kube-proxy-b8ftz                     0/1     ContainerCreating   0          69m
kube-proxy-px9l5                     1/1     Running             0          117m
kube-scheduler-k8s-master            1/1     Running             0          116m
```