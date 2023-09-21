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

## 4. wait install success...
