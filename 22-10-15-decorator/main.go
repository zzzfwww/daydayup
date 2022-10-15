package main

// ---- 抽象层 ------
// 抽象的构件
type Phone interface {
	Show() // 构件的功能
}

// 装饰器基础类（该类应该为interface，但是interface在golang语言中不可以有成员属性）
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {

}

// ---------实现层 ------
// 基础的构件

type HuaWei struct {
}

func (hw *HuaWei) Show() {
	println("show huawei phone")
}

type XiaoMi struct {
}

func (xm *XiaoMi) Show() {
	println("show xiaomi phone")
}

// 具体的装饰器类

type MoDecorator struct {
	Decorator // 继承基础装饰器类（主要继承Phone成员属性）
}

func (md *MoDecorator) Show() {
	md.phone.Show()
	println("贴膜的手机")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{
		Decorator{
			phone: phone,
		},
	}
}

type KeDecorator struct {
	Decorator // 继承基础装饰器类（主要是Phone成员属性）
}

func (kd *KeDecorator) Show() {
	kd.phone.Show()
	println("戴手机壳的手机")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{
		Decorator{
			phone: phone,
		},
	}
}

// 业务逻辑层
func main() {

	var huawei Phone
	huawei = new(HuaWei)
	huawei.Show()

	println("=======")
	var moHuawei Phone
	moHuawei = NewMoDecorator(huawei)
	moHuawei.Show()

	println("=========")
	var keHuawei Phone
	keHuawei = NewKeDecorator(huawei)
	keHuawei.Show()
	println("==========")
	var keMoHuawei Phone
	keMoHuawei = NewMoDecorator(keHuawei)
	keMoHuawei.Show()

}
