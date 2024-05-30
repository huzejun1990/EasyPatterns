// @Author huzejun 2024/5/27 3:34:00
package main

import "fmt"

//司机张三，李四 汽车 宝马 奔驰
//1 -张3 开奔驰
//2 -李4 开宝马

// 奔驰汽车
type Benz struct {
}

func (b *Benz) Run() {
	fmt.Println("Benz is running...")
}

// 宝马汽车
type BMW struct {
}

func (b *BMW) Run() {
	fmt.Println("BMW is running...")
}

// 司机张三
type Zhang3 struct {
	//...
}

func (z3 *Zhang3) DriveBenz(benz *Benz) {
	benz.Run()
	fmt.Println("zhang3 Drive benz")
}

// +
func (z3 *Zhang3) DriveBmw(bmw *BMW) {
	fmt.Println("zhang3 Drive bmw")
	bmw.Run()
}

// 司机李四
type Li4 struct {
	//
}

func (li4 *Li4) DriveBmw(bmw *BMW) {
	fmt.Println("li4 Drive Bmw")
	bmw.Run()
}

// +
func (li4 *Li4) DriveBenz(benz *Benz) {
	fmt.Println("li4 Drive Benz")
	benz.Run()
}

func main() {
	// 1 -张3 开奔驰
	benz := &Benz{}
	zhang3 := &Zhang3{}
	zhang3.DriveBenz(benz)

	//2 -李4 开宝马
	bmw := &BMW{}
	li4 := &Li4{}
	li4.DriveBmw(bmw)
}
