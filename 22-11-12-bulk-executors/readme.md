## bulk-executors

### 测试go-zero bulk insert mysql

单测文件运行[bulkinsert](./sqlx/bulkinsert_test.go)

运行函数 `TestBulkInserter`

本意想简单测试bulk insert 但是最终发现引入了非常多的其它组件和包，待后续再慢慢修改，简化。

### change log
* 2022-11-12 去除sql 熔断器代码