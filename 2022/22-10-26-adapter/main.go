package main

// 适配的目标
type V5 interface {
	Use5V()
}

// 业务类，依赖5V接口
type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v: v}
}

func (p *Phone) Charge() {
	println("Phone charge...")
	p.v.Use5V()
}

// 被适配的角色，适配者
type V220 struct {
}

func (v *V220) User220V() {
	println("use 220v charge")
}

type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	println("use adapter charge...")
	// 调用是陪着的方法
	a.v220.User220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{
		v220: v220,
	}
}

func main() {
	iphone := NewPhone(NewAdapter(new(V220)))
	iphone.Charge()
}
