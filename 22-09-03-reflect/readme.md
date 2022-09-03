## 反射
不要使用反射，除非你真的需要。但是当你不使用反射时，不要认为这是因为反射很慢，它也可以很快。

### normal版本
```shell
goos: darwin
goarch: amd64
pkg: test-reflect/normal
cpu: Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz
BenchmarkPopulateStructReflect
BenchmarkPopulateStructReflect-8   	11285409	        99.30 ns/op	       8 B/op	       1 allocs/op
PASS
```

### cache版本
```shell
goos: darwin
goarch: amd64
pkg: test-reflect/cache
cpu: Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz
BenchmarkPopulateStructReflectCache
BenchmarkPopulateStructReflectCache-8   	39221967	        29.53 ns/op	       0 B/op	       0 allocs/op
PASS
```

### unsafe版本
```shell
goos: darwin
goarch: amd64
pkg: test-reflect/unsafe
cpu: Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz
BenchmarkPopulateStructReflectCache
BenchmarkPopulateStructReflectCache-8    	62634748	        18.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkPopulateStructReflectCache2
BenchmarkPopulateStructReflectCache2-8   	201304066	         5.780 ns/op	       0 B/op	       0 allocs/op
BenchmarkPopulateUnsafe3
BenchmarkPopulateUnsafe3-8               	1000000000	         0.4690 ns/op	       0 B/op	       0 allocs/op
BenchmarkPopulate
BenchmarkPopulate-8                      	1000000000	         0.2664 ns/op	       0 B/op	       0 allocs/op
PASS
```

## 结论
反射并不一定很慢，但是你必须付出相当大的努力，通过运用 Go 内部机理知识，在你的代码中随意撒上不安全的味道 ，以使其真正加速。

## 参考资料
[orgin](https://philpearl.github.io/post/aintnecessarilyslow/)

[wechat](https://mp.weixin.qq.com/s/fzmN6zFVioQGedTdSDmyqQ)