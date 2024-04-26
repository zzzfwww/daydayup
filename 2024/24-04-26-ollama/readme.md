# Ollama
Ollama 是一个本地推理框架客户端，可一键部署如 Llama 2, Mistral, Llava 等大型语言模型。 Dify 支持接入 Ollama 部署的大型语言模型推理和 embedding 能力。

## 快速接入

### 下载并启动Ollama
1. 下载Ollama
  访问 https://ollama.ai/download，下载对应系统 Ollama 客户端。
> 大陆地区下载会报错，
> 需要解决，所以linux环境下，直接运行install.sh就可以完成安装


2. 运行 Ollama 并与 Llava 聊天
```shell
ollama run llava
```
启动成功后，ollama 在本地 11434 端口启动了一个 API 服务，可通过 http://localhost:11434 访问。

## 在Linux上设置环境变量
如果Ollama作为systemd服务运行，应该使用systemctl设置环境变量：

通过调用systemctl edit ollama.service编辑systemd服务。这将打开一个编辑器。

对于每个环境变量，在[Service]部分下添加一行Environment：

```conf
[Service]
Environment="OLLAMA_HOST=0.0.0.0"
```
保存并退出。

重载systemd并重启Ollama：

```shell
systemctl daemon-reload
systemctl restart ollama
```
