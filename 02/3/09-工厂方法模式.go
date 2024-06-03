// @Author huzejun 2024/6/4 3:06:00
package main

import "fmt"

// ===== 抽象层 ====
// 水果类（抽象的接口）
type Fruit interface {
	Show()
}

// 工厂类（抽象的接口）
type AbstractFactory interface {
	CreateFruit() Fruit //生产水果类（抽象）的生产器方法
}

// ===== 基础模块层 ====
type Apple struct {
	Fruit //为了易于理解
}

func (apple *Apple) Show() {
	fmt.Println("我是Apple")
}

type Banana struct {
	Fruit //为了易于理解
}

func (banana *Banana) Show() {
	fmt.Println("我是Banana")
}

type Pear struct {
	Fruit //为了易于理解
}

func (pear *Pear) Show() {
	fmt.Println("我是Pear")
}

type JanpanPear struct {
	Fruit
}

func (jp *JanpanPear) Show() {
	fmt.Println("我是JapanPear")
}

// ===== 基础的工厂模块 ====
// 具体的苹果工厂
type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit

	//生产一个具体的苹果
	fruit = new(Apple)

	return fruit
}

// 具体的香蕉工厂
type BananaFactory struct {
	AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit

	//生产一个具体的香蕉
	fruit = new(Banana)

	return fruit
}

// 具体的梨工厂
type PearFactory struct {
	AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
	var fruit Fruit

	//生产一个具体的梨
	fruit = new(Pear)

	return fruit
}

type JanpanPearFactory struct {
	AbstractFactory
}

func (fac *JanpanPearFactory) CreateFruit() Fruit {
	var fruit Fruit

	//生产一个具体的梨
	fruit = new(JanpanPear)

	return fruit
}

// 业务逻辑层
func main() {
	//需求1：需要一个具体的苹果对象
	//1-先要一个具体的苹果工厂
	var appleFac AbstractFactory
	appleFac = new(AppleFactory)

	//2-生产一个具体的水果
	var apple Fruit
	apple = appleFac.CreateFruit()

	apple.Show()

	//需求2：需要一个具体的香蕉对象
	var bananaFac AbstractFactory
	bananaFac = new(BananaFactory)

	var banana Fruit
	banana = bananaFac.CreateFruit() //多态
	banana.Show()

	//需求3：需要一个具体的梨对象
	pearFac := new(PearFactory)
	pear := pearFac.CreateFruit()
	pear.Show()

	jppearFac := new(JanpanPearFactory)
	jppear := jppearFac.CreateFruit()
	jppear.Show()
}
