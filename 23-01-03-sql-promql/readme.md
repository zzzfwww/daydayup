# Prometheus

## metrics类型
HELP:说明

TYPE : metrics类型
```text
alertmanager, alects, invalid., total(version= "v1"@139383232 0
```
Histogram参考文档[link](https://www.cnblogs.com/ryanyangcs/p/11309373.html)

## PromQL

瞬时向量:包含该时间序列中最新的一个样本值

区间向量:一段时间范围内的数据

Offset :查看多少分钟之前的数据offset 30m
Labelsets: 
* 辻濾出具有handler="/login"的label 的数据。
* 正则匹配 : http_request_totạl{handler= ~"*login.*"}
* 剔除某个lable: http_request_totạl{handler!=".*login.*"}
* 匹配两个值 : http_request_totạl{handler=~"/login|/password"}

数学运算: +-*/% ^心

查看主机内存总大小(Mi)心

BMit : mode_memery_MemTotal_bytes 11024 /1024

mode_memery_MemTotal_bytes/ 1024 /1024 < 3000

结合运算:
* and or
* mode_memery_MemTotal_bytes/ 1024 /1024 <= 2772 or mode_memery_MemTotal_bytes / 1024 /1024 == 3758.59765625

unless :排除
* mode_memery_MemTotal_bytes/1024/1024 >= 2772 unless mode_memery_MemTotal_bytes/1024/1024 == 3758.59765625

运算符优先级
```text
^
*/%
+ -
==,!=,<=,< >= >
and unless
or
```

聚合操作:
* sum(mode_memery_MemTotal_bytes) / 1024^2求和
* 根据某个字段迸行統汁sum(http_request_totạl) by (statuscode, handler)
* min(mode_memery_MemTotal_bytes) 最小值max 
* avg(mode_memery_MemTotal_bytes) 平均值avg
* 标准差: stdde
* 标准差异: stdvar 
* count(http_request_totạl) 计数。
* count_values("count", mode_memery_MemTotal_bytes)对 value进行统计计数
* topk(5, sum(http_request_totạl) by (statuscode, handler))取前N条时序
* bottomk(3, sum(http_request_totạl) by (statuscode, handler))取后N条时序
* 取当前数据的中位数guantile(0.5, http_request_totạl)
