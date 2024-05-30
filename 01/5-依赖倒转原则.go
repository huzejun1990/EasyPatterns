// @Author huzejun 2024/5/31 1:38:00
package main

import "fmt"

// -----> 抽象层 <---------
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// -----> 实现层 <---------
type BenZ struct {
	//...
}

func (b *BenZ) Run() {
	fmt.Println("Benz is Running...")
}

type Bmw struct {
}

func (b *Bmw) Run() {
	fmt.Println("Bmw is Running...")
}

type Zhangsan struct {
	//...
}

func (z3 *Zhangsan) Drive(car Car) {
	fmt.Println("zhang3 drive car")
	car.Run()
}

type Lisi struct {
	//...
}

func (li4 *Lisi) Drive(car Car) {
	fmt.Println("li4 drive car")
	car.Run()
}

type Wangwu struct {
}

func (w5 *Wangwu) Drive(car Car) {
	fmt.Println("wang5 drive car")
	car.Run()
}

// -----> 业务逻辑层 <---------
func main() {
	//zhang3-->benz
	var benz Car
	benz = new(BenZ)

	var zhang3 Driver
	zhang3 = new(Zhangsan)
	zhang3.Drive(benz)

	//li4-->开宝马
	var bmw Car
	bmw = new(Bmw)
	var li4 Driver
	li4 = new(Lisi)
	li4.Drive(bmw)

	zhang3.Drive(bmw)

	var w5 Driver
	w5 = new(Wangwu)
	w5.Drive(benz)
}
