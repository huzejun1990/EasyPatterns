// @Author huzejun 2024/6/3 23:26:00
package main

import "fmt"

// ------抽象层-------
type Fruit interface {
	Show() //接口的方法
}

// ------实现层-------
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

// =====工厂模块=====
type Factory struct {
}

func (fac *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit

	if kind == "apple" {
		//apple构造初始化业务
		fruit = new(Apple) //满足多态条件的巅值，父类指针指向子类对象
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit
}

// ------业务逻辑层-------
func main() {
	factory := new(Factory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
