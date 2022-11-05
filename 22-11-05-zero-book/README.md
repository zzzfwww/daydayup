# zero-book

## go-zero.dev官网错误点
* model的sql文件应该修改
```sql
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名称',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '用户密码',
  `gender` char(5) NOT NULL COMMENT '男｜女｜未公开',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```
> 插入一条数据 
```sql
INSERT INTO `user`
(id, name, password, gender, create_time, update_time)
VALUES(1, 'name', 'name', '男', '2022-11-05 11:01:56', '2022-11-05 11:01:56');
```

* [https://go-zero.dev/cn/docs/advance/business-coding](https://go-zero.dev/cn/docs/advance/business-coding)

代码
```go
// findOneByNumber应该变为FindOneByName
    userInfo, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, req.Username)
    switch err {
    case nil:
    case model.ErrNotFound:
        return nil, errors.New("用户名不存在")
    default:
        return nil, err
    }
    
    if userInfo.Password != req.Password {
        return nil, errors.New("用户密码不正确")
    }
```
## 启动api服务
* 配置文件添加和修改
```yaml
Name: user-api
Host: 0.0.0.0
Port: 8888

Mysql:
  DataSource: root:yourpassword@tcp(localhost:3306)/gozero?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
CacheRedis:
  - Host: 127.0.0.1:6379
    Pass:
    Type: node

Auth:
  AccessSecret: 2hJx9ul9buoA73P
  AccessExpire: 1800
```
* 请求数据
```shell
curl --location --request POST 'http://localhost:8888/user/login' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "username":"name",
  "password":"name"
  }'
```
* 返回数据
```shell
{
    "id": 1,
    "name": "name",
    "gender": "男",
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Njc2MTk4NTQsImlhdCI6MTY2NzYxODA1NCwidXNlcklkIjoxfQ.iqn5Cb8i6sPY_JpVEo683DL1vmxhRjBY4aPdSN9mt3c",
    "accessExpire": 1667619854,
    "refreshAfter": 1667618954
}
```
## 启动 search-api
* api配置
```yaml
Name: search-api
Host: 0.0.0.0
Port: 8889

Auth:
  AccessSecret: 2hJx9ul9buoA73P
  AccessExpire: 1800
```

* 请求数据
```shell
curl --location --request GET 'localhost:8889/search/do?name=name' \
--header 'Content-Type: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Njc2MTk4NTQsImlhdCI6MTY2NzYxODA1NCwidXNlcklkIjoxfQ.iqn5Cb8i6sPY_JpVEo683DL1vmxhRjBY4aPdSN9mt3c' \
--data-raw '{
    "inputData":"sss"
}'
```
* 返回数据
* jwt鉴权失败httpcode返回401
* http返回200
```shell
{
    "name": "name",
    "count": 0
}
```