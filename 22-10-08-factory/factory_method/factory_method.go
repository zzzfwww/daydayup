package factory_method

// === 抽象层 =====

// 水果类（抽象的接口）
type Fruit interface {
	Show()
}

// 工厂类（抽象的接口）
type AbstractFactory interface {
	CreateFruit() Fruit // 生产水果类（抽象）的生产器方法
}

// ==== 基础模块层===
type Apple struct {
	Fruit // 为了易于理解
}

func (a *Apple) Show() {
	println("apple")
}

type Banana struct {
	Fruit // 为了易于理解
}

func (b *Banana) Show() {
	println("banana")
}

type Pear struct {
	Fruit // 为了易于理解
}

// ==== 基础的工厂模块 =====
// 具体的工厂
type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	// 生产一个具体的实例
	fruit = new(Apple)
	return fruit
}

// 具体的工厂
type BananaFactory struct {
	AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit
	// 生产一个具体的实例
	fruit = new(Banana)
	return fruit
}

// 具体的工厂
type PearFactory struct {
	AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
	var fruit Fruit
	// 生产一个具体的实例
	fruit = new(Pear)
	return fruit
}

func (p *Pear) Show() {
	println("pear")
}
