package factory_method

import "testing"

func Test_Apple(t *testing.T) {
	// 需求1 需要一个具体的苹果对象
	// 1. 先要一个具体的苹果工厂
	var appleFac AbstractFactory
	appleFac = new(AppleFactory)

	// 2. 生产一个具体的水果
	var apple Fruit
	apple = appleFac.CreateFruit()
	apple.Show()
}

func Test_Banana(t *testing.T) {
	// 需求1 需要一个具体的香蕉对象
	// 1. 先要一个具体的苹果工厂
	var bananaFac AbstractFactory
	bananaFac = new(BananaFactory)

	// 2. 生产一个具体的水果
	var banana Fruit
	banana = bananaFac.CreateFruit()
	banana.Show()
}

func Test_Pear(t *testing.T) {
	// 需求1 需要一个具体的香蕉对象
	// 1. 先要一个具体的苹果工厂
	var pearFac AbstractFactory
	pearFac = new(PearFactory)

	// 2. 生产一个具体的水果
	var pear Fruit
	pear = pearFac.CreateFruit()
	pear.Show()
}
