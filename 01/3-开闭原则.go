// @Author huzejun 2024/5/26 1:42:00
package main

import "fmt"

// 抽象的银行业务员
type AbstractBanker interface {
	DoBusi() //抽象的处理业务接口

}

// 存款的业务员
type SaveBanker struct {
	//AbstractBanker
}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("进行了存款")
}

// ++++++++++
// 转账的业务员
type TranferBanker struct {
	//AbstractBanker
}

func (sb *TranferBanker) DoBusi() {
	fmt.Println("进行了转账")
}

// ++++++++++
// 股票的业务员
type ShareBanker struct {
	//AbstractBanker
}

func (sb *ShareBanker) DoBusi() {
	fmt.Println("进行了股票")
}

// 实现一个架构层（基于抽象层进行业务封装-针对interface接口进行封装）
func BankBusiness(banker AbstractBanker) {
	//通过接口向下来调用（多态的现象）
	banker.DoBusi()
}

// 逻辑
func main() {
	/*
		//存款的业务
		sb := SaveBanker{}
		sb.DoBusi()

		//转账的业务
		st := TranferBanker{}
		st.DoBusi()

		//股票的业务
		sb2 := ShareBanker{}
		sb2.DoBusi()
	*/

	//存款的业务
	BankBusiness(&SaveBanker{})
	//转账的业务
	BankBusiness(&TranferBanker{})
	//股票的业务
	BankBusiness(&ShareBanker{})
}
