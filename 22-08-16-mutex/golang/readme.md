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

## 原因
> go 版本 1.17.13
```go
// RLock locks rw for reading.
//
// It should not be used for recursive read locking; a blocked Lock
// call excludes new readers from acquiring the lock. See the
// documentation on the RWMutex type.
func (rw *RWMutex) RLock() {
	if race.Enabled {
		_ = rw.w.state
		race.Disable()
	}
	if atomic.AddInt32(&rw.readerCount, 1) < 0 {
		// A writer is pending, wait for it.
		runtime_SemacquireMutex(&rw.readerSem, false, 0)
	}
	if race.Enabled {
		race.Enable()
		race.Acquire(unsafe.Pointer(&rw.readerSem))
	}
}
```
> atomic.AddInt32(&rw.readerCount,1) <0 如果有写锁在等待，读锁需要等写锁！

```go
// If a goroutine holds a RWMutex for reading and another goroutine might
// call Lock, no goroutine should expect to be able to acquire a read lock
// until the initial read lock is released. In particular, this prohibits
// recursive read locking. This is to ensure that the lock eventually becomes
// available; a blocked Lock call excludes new readers from acquiring the
// lock.
type RWMutex struct {
	w           Mutex  // held if there are pending writers
	writerSem   uint32 // semaphore for writers to wait for completing readers
	readerSem   uint32 // semaphore for readers to wait for completing writers
	readerCount int32  // number of pending readers
	readerWait  int32  // number of departing readers
}
```
翻译：
>如果一个协程持有读锁，另一个协程可能会调用Lock加写锁，那么再也没有一个协程可以获得读锁，直到前一个读锁释放，这是为了禁止读锁递归。也确保了锁最终可用，一个阻塞的写锁调用会将新的读锁排除在外。

## 结语：
- Go的设计者比较「偏执」，认为「不好」的设计坚决不去实现，就如锁的实现不应该依赖线程、协程信息；可重入（递归）锁是一种不好的设计。所以这种看似有BUG的设计，也存在一定的道理。