// @Author huzejun 2024/6/14 20:42:00
package main

import "fmt"

// 抽象类，制作饮料，包裹一个模板，全部实现的步骤
type Beverage interface {
	//煮开水
	BoilWater()

	//冲泡
	Brew()

	//倒入杯中
	PourInCup()

	//添加酌料
	AddThings()

	//是否添加酌料的Hook
	WantAddThings() bool
}

// 封装一套流程模板基类，让具体的制作流程继承而且实现
type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	if t.b.WantAddThings() == true {
		t.b.AddThings()
	}
}

// 具体的制作饮料的流程，制作咖啡
type MakeCoffee struct {
	template //继承模板
}

// 煮开水
func (mc *MakeCoffee) BoilWater() {
	fmt.Println("将水烧到100摄氏度")
}

// 冲泡
func (mc *MakeCoffee) Brew() {
	fmt.Println("用水冲咖啡粉")
}

// 倒入杯中
func (mc *MakeCoffee) PourInCup() {
	fmt.Println("将冲好的咖啡倒入陶瓷杯中")
}

// 添加酌料
func (mc *MakeCoffee) AddThings() {
	fmt.Println("添加牛奶和糖")
}

func (mc *MakeCoffee) WantAddThings() bool {
	return true
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)

	//b 是Beverage,是MakeCoffee的接口，这里需要给接口赋值，让b指向具体的子类
	//来触发b的全部方法的多态特性
	makeCoffee.b = makeCoffee

	return makeCoffee
}

// 具体的制作茶叶的流程
type MakeTea struct {
	template
}

// 煮开水
func (mt *MakeTea) BoilWater() {
	fmt.Println("将水煮到80摄氏度")
}

// 冲泡
func (mt *MakeTea) Brew() {
	fmt.Println("用水冲茶叶")
}

// 倒入杯中
func (mt *MakeTea) PourInCup() {
	fmt.Println("将冲好的茶水倒入茶壶中")
}

// 添加酌料
func (mt *MakeTea) AddThings() {
	fmt.Println("添加柠檬")
}

func (mt *MakeTea) WantAddThings() bool {
	return false
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}

func main() {
	//1、制作一杯咖啡
	makeCoffee := NewMakeCoffee()

	//调用模板方法
	makeCoffee.MakeBeverage()

	// ---------
	fmt.Println("---------------")
	//2. 制作一杯茶
	makeTea := NewMakeTea()
	makeTea.MakeBeverage()
}
