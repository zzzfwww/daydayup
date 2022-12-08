<!-- vscode-markdown-toc -->
* 1. [方法一：使用 hostPath 提供持久化](#hostPath)
* 2. [访问 Kuboard](#Kuboard)

<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->
# 安装 Kuboard v3 - kubernetes

在 K8S 中安装 Kuboard，主要考虑的问题是，如何提供 etcd 的持久化数据卷。建议的两个选项有：

- 使用 hostPath 提供持久化存储，将 kuboard 所依赖的 Etcd 部署到 Master 节点，并将 etcd 的数据目录映射到 Master 节点的本地目录；
- 使用 StorageClass 动态创建 PV 为 etcd 提供数据卷；
> 如果使用此方法安装 Kuboard 碰到问题，强烈建议您使用 docker run 或者 static-pod 的方式安装 kuboard .

##  1. <a name='hostPath'></a>方法一：使用 hostPath 提供持久化
```bash
kubectl apply -f https://addons.kuboard.cn/kuboard/kuboard-v3.yaml
# 您也可以使用下面的指令，唯一的区别是，该指令使用华为云的镜像仓库替代 docker hub 分发 Kuboard 所需要的镜像
# kubectl apply -f https://addons.kuboard.cn/kuboard/kuboard-v3-swr.yaml
```
> 定制参数
> 
> 如果您想要定制 Kuboard 的启动参数，请将该 YAML 文件下载到本地，并修改其中的 ConfigMap

等待 Kuboard v3 就绪

执行指令 watch kubectl get pods -n kuboard，等待 kuboard 名称空间中所有的 Pod 就绪，如下所示，

如果结果中没有出现 kuboard-etcd-xxxxx 的容器，请查看  中关于 缺少 Master Role 的描述。
```bash
[root@k8smaster ~]#  kubectl get pods -n kuboard
NAME                              READY   STATUS    RESTARTS   AGE
kuboard-agent-2-cf648f7fd-8xv2c   1/1     Running   0          39s
kuboard-agent-7f4c97f888-gt7f7    1/1     Running   0          39s
kuboard-etcd-x5gqf                1/1     Running   0          91s
kuboard-v3-79797c7b84-fflkj       1/1     Running   0          91s
```

##  2. <a name='Kuboard'></a>访问 Kuboard
在浏览器中打开链接 http://192.168.3.201:30080

输入初始用户名和密码，并登录

用户名： admin

密码： Kuboard123
