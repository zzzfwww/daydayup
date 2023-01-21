## 问题
goland编译器直接运行test文件报错
```shell
# command-line-arguments [command-line-arguments.test]
./test_test.go:6:9: undefined: NewTest
```

## 解决办法
直接terminal运行`go test -v test_test.go test.go`
