// @Author huzejun 2024/5/31 3:58:00
package main

import "fmt"

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

// 给小猫添加一个可以睡觉的方法（使用继承来实现）
type CatB struct {
	Cat
}

func (ch *CatB) Sleep() {
	fmt.Println("小猫睡觉")
}

// 给小猫添加一个可以睡觉的方法（使用组合的方式）
type CatC struct {
	//C *Cat //组合进来一个Cat类
}

func (cc *CatC) Eat(c *Cat) {
	//cc.C.Eat() //调用组合进来的原有类的功能
	c.Eat()
}

func (cc *CatC) Sleep() {
	fmt.Println("小猫睡觉")
}

func main() {
	c := &Cat{}
	c.Eat()

	fmt.Println("---------")
	cb := &CatB{}
	cb.Eat()
	cb.Sleep() //+++

	fmt.Println("---------")
	cc := &CatC{}
	cc.Eat(c)
	cc.Sleep()
}
