// @Author huzejun 2024/6/11 0:37:00
package main

import "fmt"

type Goods struct {
	Kind string //商品的种类
	Fact bool   //商品的真伪
}

// ======== 抽象层 =========
type Shopping interface {
	Buy(goods *Goods) //某任务
}

// ======== 实现层 =========
type KoreaShopping struct {
}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物，买了", goods.Kind)
}

type AmericaShopping struct {
}

func (as *AmericaShopping) Buy(goods *Goods) {
	fmt.Println("去美国进行了购物，买了", goods.Kind)
}

type AfrikaShopping struct {
}

func (as *AfrikaShopping) Buy(goods *Goods) {
	fmt.Println("去非州进行了购物，买了", goods.Kind)
}

// 海外代理
type OverSeasProxy struct {
	shopping Shopping //代理某个主题，这里是抽象的数据类型
}

func NewProxy(shopping Shopping) Shopping {
	return &OverSeasProxy{shopping}
}

func (op *OverSeasProxy) Buy(goods *Goods) {
	//1 +先辨别真伪
	if op.distinguish(goods) == true {
		//2 +调用具体要被代理的财物方法的Buy()方法
		op.shopping.Buy(goods) //多态
		//3 +海关安检
		op.check(goods)
	}
}

func (op *OverSeasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "] 进行了辨别真伪.")
	if goods.Fact == false {
		fmt.Println("发现假货", goods.Kind, ", 不应该购买，")
	}
	return goods.Fact
}

// 安检流程
func (op *OverSeasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "] 进行了海关检查，成功带回祖国..")
}

//======== 业务逻辑层 =========

func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}
	g2 := Goods{
		Kind: "CET4证书",
		Fact: false,
	}

	var KShoping Shopping
	KShoping = new(KoreaShopping)

	/*
		//1 先辨别真伪
		if g1.Fact != true {
			//2 Buy
			shoping.Buy(g1)
			//3 海关安检
			fmt.Println("带回祖国")
		}*/

	var k_proxy Shopping
	k_proxy = NewProxy(KShoping)
	k_proxy.Buy(&g1)

	var aShopping Shopping
	aShopping = new(AmericaShopping)

	var a_proxy Shopping
	a_proxy = NewProxy(aShopping)
	a_proxy.Buy(&g2)
	a_proxy.Buy(&g1)
}
