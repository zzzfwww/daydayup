package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var lock sync.RWMutex
	lock.RLock()
	println("RLock 1...")
	go func() {
		lock.Lock()
		println("lock...")
		time.Sleep(time.Second)
		lock.Unlock()
		println("unlock...")
		wg.Done()
	}()
	time.Sleep(time.Second)

	lock.RLocker()
	println("RLock 2...")
	lock.RUnlock()
	println("RUnlock 1...")
	lock.RUnlock()
	println("RUnlock 2...")
}
