### java 版锁不死锁
```shell
/Library/Java/JavaVirtualMachines/temurin-17.jdk/Contents/Home/bin/java -javaagent:/Applications/IntelliJ IDEA CE.app/Contents/lib/idea_rt.jar=63593:/Applications/IntelliJ IDEA CE.app/Contents/bin -Dfile.encoding=UTF-8 -classpath /Users/zfw/Documents/study/personalgithub/daydayup/22-08-16-mutex/java/out/production/java demo
read lock
read unlock
write lock
write unlock
```
### 原因
1. Java的ReentrantReadWriteLock支持锁降级，但不能升级，即获取了写锁的线程，可以继续获取读锁，但获取读锁的线程无法再获取写锁；
2. ReentrantReadWriteLock实现了公平和非公平两种锁，公平锁的情况下，获取读锁、写锁前需要看同步队列中是否先线程在我之前排队；非公平锁的情况下：写锁可以直接抢占锁，但是读锁获取有一个让步条件，如果当前同步队列head.next是一个写锁在等待，并且自己不是重入的，就要让步等待。

在Java的实现下，如果一个线程持有了读锁，写锁自然是需要等待的，但是持有读锁的线程也可以再次重入该读锁。