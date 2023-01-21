## http middleware

### alice chain
[https://github.com/justinas/alice](https://github.com/justinas/alice)

### alice库封装思想
example.go属于最简版本封装，看不太明白的时候可以详细看下这个封装，相信就能了解alice的核心思想。

### 核心逻辑
```go
for i := range c.constructors {
    h = c.constructors[len(c.constructors)-1-i](h)
}
```
此for循环是整个chain的最核心逻辑

handler示例
- 不使用chain的写法
```go
customizedHandler = logger(timeout(ratelimit(helloHandler)))
```
- 使用chain的写法
```go
customizedHandler = chain.New(logger,timeout,ratelimit).Then(helloHandler)
```
调用栈
```shell
[exec of logger logic]           函数栈: []

[exec of timeout logic]          函数栈: [logger]

[exec of ratelimit logic]        函数栈: [timeout/logger]

[exec of helloHandler logic]     函数栈: [ratelimit/timeout/logger]

[exec of ratelimit logic part2]  函数栈: [timeout/logger]

[exec of timeout logic part2]    函数栈: [logger]

[exec of logger logic part2]     函数栈: []
```

类似俄罗斯套娃，一层套一层，最核心函数的入参是http.Handler，出参也是http.Handler，函数式编程的思想。

### 结论
* 用好工具，能让代码简化，并且提升编码能力