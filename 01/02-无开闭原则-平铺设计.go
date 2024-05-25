// @Author huzejun 2024/5/25 22:31:00
package main

import "fmt"

type Banker struct {
}

// 存款业务
func (b *Banker) Save() {
	fmt.Println("进行了 存款业务...")
}

// 转账业务
func (b *Banker) Tansfer() {
	fmt.Println("进行了 转账业务...")
}

// 支付业务
func (b *Banker) Pay() {
	fmt.Println("进行了 支付业务...")
}

// 股票功能
func (b *Banker) Shares() {
	fmt.Println("进行了 股票业务...")
}

func main() {
	banker := Banker{}

	banker.Save()
	banker.Tansfer()
	banker.Pay()
	banker.Shares()
}
