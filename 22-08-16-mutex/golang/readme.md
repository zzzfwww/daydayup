## golang 死锁
报错信息
```shell
dayup/22-08-16-mutex/golang/main.go #gosetup
/private/var/folders/y6/9zm7qybd7yb_lzjlx0xjb7nc0000gn/T/GoLand/___go_build_main_go
RLock 1...
RLock 2...
RUnlock 1...
fatal error: sync: RUnlock of unlocked RWMutex
lock...

goroutine 1 [running]:
runtime.throw({0x1069174, 0x1})
        /Users/zfw/go/go1.17.4/src/runtime/panic.go:1198 +0x71 fp=0xc000092f00 sp=0xc000092ed0 pc=0x102bb31
sync.throw({0x1069174, 0x102d39b})
        /Users/zfw/go/go1.17.4/src/runtime/panic.go:1184 +0x1e fp=0xc000092f20 sp=0xc000092f00 pc=0x10515be
sync.(*RWMutex).rUnlockSlow(0xc0000a2000, 0xc0000000)
        /Users/zfw/go/go1.17.4/src/sync/rwmutex.go:93 +0x39 fp=0xc000092f48 sp=0xc000092f20 pc=0x1056739
sync.(*RWMutex).RUnlock(...)

```