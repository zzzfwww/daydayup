## 锁对比
### golang mutex
#### 背景
- golang 锁不可重入
#### 结果
- 导致一些场景死锁
#### 解决办法
- 需要代码健壮性更强
#### 抑或
- golang本来就存在这个bug

### Java mutex
- 不存在golang锁的问题