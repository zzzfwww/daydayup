package solution

import (
	"sync"
	"testing"
	"time"
)

func TestNewMutex(t *testing.T) {
	num := 10
	m := NewMutex()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		m.Lock()
		num += 10
		m.Unlock()
	}()
	wg.Wait()
	println(num)
	if num != 20 {
		t.Errorf("want:%v got:%v", 20, num)
	}
}

func TestMutex_TryLock(t *testing.T) {
	num := 10
	m := NewMutex()
	var wg sync.WaitGroup
	wg.Add(1)
	m.Lock()
	go func() {
		defer wg.Done()
		lock := m.TryLock()
		if lock {
			t.Errorf("want:%v got:%v", false, lock)
		}
	}()
	wg.Wait()
	println(num)
}

func TestMutex_LockTimeout(t *testing.T) {
	num := 10
	m := NewMutex()
	var wg sync.WaitGroup
	wg.Add(1)
	m.Lock()
	go func() {
		defer wg.Done()
		lock := m.LockTimeout(1000 * time.Millisecond)
		if !lock {
			t.Errorf("want:%v got:%v", true, lock)
		}
	}()
	time.Sleep(500 * time.Millisecond)
	m.Unlock()
	wg.Wait()
	println(num)
}
