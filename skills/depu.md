# depu 工具

在需要检查当前项目是否有依赖更新时可以使用此工具

## 安装
```shell
$ go install github.com/kevwan/depu@latest
```

## 使用
```shell
➜  bookstore git:(master) ✗ depu          
                  PACKAGE                  | CURRENT | LATEST  | GOVERSION  
-------------------------------------------+---------+---------+------------
  github.com/zeromicro/go-zero             | v1.3.2  | v1.4.2  |      1.15  
  github.com/zeromicro/go-zero/tools/goctl | v1.3.3  | v1.4.2  |      1.16  
  google.golang.org/grpc                   | v1.45.0 | v1.51.0 |      1.14  
  google.golang.org/protobuf               | v1.28.0 | v1.28.1 |      1.11  
➜  bookstore git:(master) ✗ 

```
