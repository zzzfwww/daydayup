# go test

## Split Test

```shell
test_slice> go test -bench=Split -benchmem     
goos: windows
goarch: amd64
pkg: demo/test_slice
cpu: Intel(R) Core(TM) i5-10600K CPU @ 4.10GHz
BenchmarkSplit-12        7016485               175.9 ns/op           112 B/op          3 allocs/op
PASS
ok      demo/test_slice 1.468s
```
* 发现一次操作要三次申请内存
* 所以修改代码，增加`result = make([]string, 0, strings.Count(s, sep)+1)`
* 再次测试发现性能提升非常明显
```shell
test_slice> go test -bench=Split -benchmem
goos: windows
goarch: amd64
pkg: demo/test_slice
cpu: Intel(R) Core(TM) i5-10600K CPU @ 4.10GHz
BenchmarkSplit-12       11726937               105.6 ns/op            48 B/op          1 allocs/op
PASS
ok      demo/test_slice 1.833s

```
