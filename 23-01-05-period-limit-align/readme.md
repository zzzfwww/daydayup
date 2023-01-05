# period-limit

具体内容可以查看go-zero[官网](https://go-zero.dev/cn/docs/blog/governance/periodlimit)

## 此内容修改点

* 使用原生的redisclient
```text
github.com/go-redis/redis/v8 v8.11.5
```
* 作为源码走读，直接把用到的组件引入进来，其中redistest包也是引入进来，并且使用原生redis8的客户端

## 详细调试align逻辑

* 时间对其，增加当前TZ时间的offset，对其逻辑代码如下
```go
if h.align {
now := time.Now()
_, offset := now.Zone()
unix := now.Unix() + int64(offset)
return h.period - int(unix%int64(h.period))
}
```
* 一开始理解起来不是特别明白，现在增加多一个场景的假设，就明白这段代码的意义
    * 假设目前UTC时间是0
    * 那么东八区的时间就是28800（8h)
    * 如果时间窗口刚好是一天86400（24h)
    * 那么这个时间时间对其则返回的时间是86400-28800（16h）
    * 也就意思当前时间开始计算窗口，则距离结束16h，换算到本地TZ时间刚好是今天的23:59:59
* 上面示例就是解释时间对其逻辑代码
