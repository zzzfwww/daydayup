# 需求

在开发和使用软件的时候，时常需要生成一些随机字符串，并且还需要设置到环境变量里面，不在 .bash_history 上显示

# 实现

```shell
[root@localhost ~]# if [ "$SECRET_KEY" = "" ]; then SECRET_KEY=`cat /dev/urandom | tr -dc A-Za-z0-9 |head -c 50`; echo "SECRET_KEY=$SECRET_KEY" >> ~/.bashrc; echo $SECRET_KEY; else echo $SECRET_KEY;fi
EyqaoJU49pXFXuNOD0ptp3QbKEBBbc7PCTVIdOpxDZ1jKtkcGx
```

# 验证

```shell
[root@localhost ~]# tail -n 1 .bashrc
SECRET_KEY=EyqaoJU49pXFXuNOD0ptp3QbKEBBbc7PCTVIdOpxDZ1jKtkcGx

[root@localhost ~]# echo $SECRET_KEY
EyqaoJU49pXFXuNOD0ptp3QbKEBBbc7PCTVIdOpxDZ1jKtkcGx
```
