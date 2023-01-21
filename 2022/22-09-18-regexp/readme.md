## regexp
### 发现问题
1. go mod init `go mod init regexp `
2. 出现报错
```shell
go: creating new go.mod: module regexp
go: to add module requirements and sums:
        go mod tidy
go mod tidy       
go: warning: "all" matched no packages
```
3. 点击运行报错
```shell
Compilation finished with exit code 0

Compiled binary cannot be executed
```

### 修复问题
- 修改go.mod`module test-regexp`
- 运行正确结果
```shell
/Users/zfw/go/go1.17.13/bin/go build -o /private/var/folders/y6/9zm7qybd7yb_lzjlx0xjb7nc0000gn/T/GoLand/___go_build_test_regexp test-regexp #gosetup
/private/var/folders/y6/9zm7qybd7yb_lzjlx0xjb7nc0000gn/T/GoLand/___go_build_test_regexp
true
true
true
Escaping symbols like: \.\+\*\?\(\)\|\[\]\{\}\^\$
true
true
true
正则表达式： foo(.?)
foo!
[6 10]
foo!
[6 10]
["foo!"]
[[6 10]]
[foo!]
[[6 10]]
[6 10]
[Hello  ]
foo false
Hello World!
Hello World!
Hello $World
Hello $World
Hello World!
Hello World!
false
1
[ ]
[food d]
[3 7 6 7]
[food d]
[3 7 6 7]
[["food" "d"] ["fool" "l"]]
[[3 7 6 7] [8 12 11 12]]
[[food d] [fool l]]
[[3 7 6 7] [8 12 11 12]]
[3 7 6 7]
option1=value1
option2=value2
option3=value3

option1=value1
option2=value2
option3=value3


Process finished with the exit code 0
```
