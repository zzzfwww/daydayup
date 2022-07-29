## 单测方式
1. go test -run "TestNewMutex" -v lock_test.go lock.go

2. go test -run "TestMutex_TryLock" -v lock_test.go lock.go 
3.  go test -run "TestMutex_LockTimeout" -v lock_test.go lock.go 
