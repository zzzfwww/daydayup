# install kubernetes by kubeSphere

## 1. set host
```shell
hostnamectl set-hostname kube
```

## 2. download kk
```bash
## you should set this env to download kk in mainland china
export KKZONE=cn
curl -sfL https://get-kk.kubesphere.io | VERSION=v3.0.7 sh -
```

## 3. install dependencies
* check dependencies
```txt
[root@kube ~]# ./kk create cluster --with-kubesphere v3.4.0


 _   __      _          _   __           
| | / /     | |        | | / /           
| |/ / _   _| |__   ___| |/ /  ___ _   _ 
|    \| | | | '_ \ / _ \    \ / _ \ | | |
| |\  \ |_| | |_) |  __/ |\  \  __/ |_| |
\_| \_/\__,_|_.__/ \___\_| \_/\___|\__, |
                                    __/ |
                                   |___/

03:33:06 EDT [GreetingsModule] Greetings
03:33:07 EDT message: [kube]
Greetings, KubeKey!
03:33:07 EDT success: [kube]
03:33:07 EDT [NodePreCheckModule] A pre-check on nodes
03:33:08 EDT success: [kube]
03:33:08 EDT [ConfirmModule] Display confirmation form
+------+------+------+---------+----------+-------+-------+---------+-----------+--------+--------+------------+------------+-------------+------------------+--------------+
| name | sudo | curl | openssl | ebtables | socat | ipset | ipvsadm | conntrack | chrony | docker | containerd | nfs client | ceph client | glusterfs client | time         |
+------+------+------+---------+----------+-------+-------+---------+-----------+--------+--------+------------+------------+-------------+------------------+--------------+
| kube | y    | y    | y       | y        |       | y     |         |           | y      |        |            |            |             |                  | EDT 03:33:08 |
+------+------+------+---------+----------+-------+-------+---------+-----------+--------+--------+------------+------------+-------------+------------------+--------------+
03:33:08 EDT [ERRO] kube: conntrack is required.
03:33:08 EDT [ERRO] kube: socat is required.

This is a simple check of your environment.
Before installation, ensure that your machines meet all requirements specified at
https://github.com/kubesphere/kubekey#requirements-and-recommendations

```
* install
```bash
[root@kube ~]# yum -y install conntrack socat
```

## 4. wait installation complete...
```txt
04:07:54 EDT [DeployKubeSphereModule] Generate KubeSphere ks-installer crd manifests
04:07:56 EDT success: [kube]
04:07:56 EDT [DeployKubeSphereModule] Apply ks-installer
04:07:57 EDT stdout: [kube]
namespace/kubesphere-system created
serviceaccount/ks-installer created
customresourcedefinition.apiextensions.k8s.io/clusterconfigurations.installer.kubesphere.io created
clusterrole.rbac.authorization.k8s.io/ks-installer created
clusterrolebinding.rbac.authorization.k8s.io/ks-installer created
deployment.apps/ks-installer created
04:07:57 EDT success: [kube]
04:07:57 EDT [DeployKubeSphereModule] Add config to ks-installer manifests
04:07:57 EDT success: [kube]
04:07:57 EDT [DeployKubeSphereModule] Create the kubesphere namespace
04:08:00 EDT success: [kube]
04:08:00 EDT [DeployKubeSphereModule] Setup ks-installer config
04:08:02 EDT stdout: [kube]
secret/kube-etcd-client-certs created
04:08:05 EDT success: [kube]
04:08:05 EDT [DeployKubeSphereModule] Apply ks-installer
04:08:14 EDT stdout: [kube]
namespace/kubesphere-system unchanged
serviceaccount/ks-installer unchanged
customresourcedefinition.apiextensions.k8s.io/clusterconfigurations.installer.kubesphere.io unchanged
clusterrole.rbac.authorization.k8s.io/ks-installer unchanged
clusterrolebinding.rbac.authorization.k8s.io/ks-installer unchanged
deployment.apps/ks-installer unchanged
clusterconfiguration.installer.kubesphere.io/ks-installer created
04:08:14 EDT success: [kube]
Please wait for the installation to complete:     >>---> 
```
* 查看k8s状态
```bash
[root@kube ~]# kubectl version
Client Version: version.Info{Major:"1", Minor:"23", GitVersion:"v1.23.10", GitCommit:"7e54d50d3012cf3389e43b096ba35300f36e0817", GitTreeState:"clean", BuildDate:"2022-08-17T18:32:54Z", GoVersion:"go1.17.13", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"23", GitVersion:"v1.23.10", GitCommit:"7e54d50d3012cf3389e43b096ba35300f36e0817", GitTreeState:"clean", BuildDate:"2022-08-17T18:26:59Z", GoVersion:"go1.17.13", Compiler:"gc", Platform:"linux/amd64"}
```
