// @Author huzejun 2024/5/24 20:23:00
package main

import "fmt"

type ClothesShop struct {
}

type ClothesWork struct {
}

func (cs *ClothesShop) Style() {
	fmt.Println("逛街的装扮")
}

func (cs *ClothesWork) Style() {
	fmt.Println("工作的装扮")
}

func main() {
	//工作的业务
	cw := ClothesWork{}
	cw.Style()

	//逛街的业务
	cs := ClothesShop{}
	cs.Style()
}

/*
// 穿衣服的方式
type Clothes struct {
	//逛街穿衣服
}

// 工作穿衣服
func (c *Clothes) OnWork() {
	fmt.Println("工作的装扮")
}

// 逛街穿衣服
func (c *Clothes) OnShop() {
	fmt.Println("逛街的装扮")
}

func main() {
	c := Clothes{}

	//工作的业务
	fmt.Println("在工作...")
	c.OnWork()

	//逛街的业务
	fmt.Println("在逛街...")
	c.OnShop()
}
*/
