# telnet

## 问题点
* 现象：
```shell
192.168.3.101 启动 nc监听端口 `nc -l 3333`

本机telnet出现如下报错
➜  sshremote telnet 192.168.3.101 3333
Trying 192.168.3.101...
telnet: connect to address 192.168.3.101: Connection refused
telnet: Unable to connect to remote host
```

* 解决办法参考[link](https://bobcares.com/blog/telnet-connection-refused-by-remote-host/)



## 实际解决办法

* 停掉Centos的防火墙
    1. 命令行界面输入命令`systemctl status firewalld.service`并按下回车键。

    2. 然后在下方可度以查看得到`active（running）`，此时说明防火墙已经被打开了。

    3. 在命令行中输入`systemctl stop firewalld.service`命令，进行关闭防火墙。

    4. 然后再使用命令`systemctl status firewalld.service`，在下方出现`disavtive（dead）`，这权样就说明防火墙已经关闭。

    5. 再在命令行中输入命令`systemctl disable firewalld.service` 命令，即可永久关闭防火墙。

* 也可直接输入命令`iptables -F` 清除所有防火墙设置

## 结果验证
* 成功
```text
➜  local telnet 192.168.3.101 3333
Trying 192.168.3.101...
Connected to 192.168.3.101.
Escape character is '^]'.
^]
telnet> q
Connection closed.
```



