// @Author huzejun 2024/6/15 21:25:00
package main

import "fmt"

/*
	商场促销有策略A(0.8折) 策略B(消费满200减100)
*/

// 销售策略
type SellStrategy interface {
	//根据原价得售卖价
	GetPrice(price float64) float64
}

type StrategyA struct {
}

func (sa *StrategyA) GetPrice(price float64) float64 {
	fmt.Println("执行策略A,所有商品打八折")
	return price * 0.8
}

type StrategyB struct {
}

func (sb *StrategyB) GetPrice(price float64) float64 {
	fmt.Println("执行策略B,所有的商品满200，减100元")
	if price >= 200 {
		price -= 100
	}
	return price
}

// 环境类
type Goods struct {
	Price    float64
	Strategy SellStrategy
}

func (g *Goods) SetStrategy(s SellStrategy) {
	g.Strategy = s
}

func (g *Goods) SellPrice() float64 {
	fmt.Println("原价值", g.Price, ".")
	return g.Strategy.GetPrice(g.Price)
}

func main() {
	nick := Goods{
		Price: 200.0,
	}

	//上午，商场执行策略A
	nick.SetStrategy(new(StrategyA))
	fmt.Println("上午Nick鞋卖", nick.SellPrice())
	//下午，商场执行策略B
	nick.SetStrategy(new(StrategyB))
	fmt.Println("下午Nick鞋卖", nick.SellPrice())
}
