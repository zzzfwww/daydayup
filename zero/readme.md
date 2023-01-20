## goctl
- goctl macOs[配置](https://segmentfault.com/a/1190000041963346)

## Centos7.8安装go出现报错
* golang官网下载go包
* wget `https://go.dev/dl/go1.19.5.linux-386.tar.gz`
* 配置好环境变量
```shell
export GOROOT=/root/go
export PATH=$GOROOT/bin:$PATH
export GOPATH=/root/gopath
export PATH=$PATH:$GOPATH/bin
export GOPROXY=https://goproxy.cn
```
* go env 报错
```text
错误：/usr/local/bin/rar: /lib/ld-linux.so.2: bad ELF interpreter: No such file or directory

解决：是因为64位系统中安装了32位程序
解决方法：
yum -y install glibc.i686
```