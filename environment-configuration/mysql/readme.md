# MySQL 本地配置

- alias 命名 mysql1

```shell
alias mysql1="mysql --defaults-file=$HOME/.mysqlconf"
```

- mysqlconf 配置文件

```ini
➜  ~ cat .mysqlconf
[client]
port=3306
user=root
password="yourpassword"
```
