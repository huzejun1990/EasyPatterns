// @Author huzejun 2024/6/11 2:42:00
package main

import "fmt"

// ---- 抽象层 ------
type Phone interface {
	Show() //构建的功能
}

// 抽象的装饰器，装饰器的基础类（该类本应该是interface,
// 但是Golang的interface语法不可以有成员属性）
type Decorator struct {
	phone Phone //组合进来的抽象的Phone
}

func (d *Decorator) Show() {

}

// ---- 实现层 ------
// 具体的构件
type Huawei struct {
}

func (hw *Huawei) Show() {
	fmt.Println("秀出了HuaWei手机")
}

type Xiaomi struct {
}

func (xm *Xiaomi) Show() {
	fmt.Println("秀出了Xiaomi手机")
}

// 具体的装饰器
type MoDecorator struct {
	Decorator //继承基础的装饰器类（主要继承Phone的成员属性）
}

func (md *MoDecorator) Show() {
	md.phone.Show() //调用被装饰的构件的原方法
	//+++++
	fmt.Println("贴膜的手机") //装饰器额外的功能
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
	Decorator //继承基础的装饰器类（主要继承Phone的成员属性）
}

func (kd *KeDecorator) Show() {
	kd.phone.Show() //调用被装饰的构件的原方法
	//+++++
	fmt.Println("手机壳的手机") //装饰器额外的功能
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone}}
}

// ---- 业务逻辑层 ------
func main() {
	var huawei Phone
	huawei = new(Huawei)
	huawei.Show()

	fmt.Println("----------------")
	//用贴膜的装饰器，得到一个新的huawei手机
	var moHuawei Phone
	moHuawei = NewMoDecorator(huawei)
	moHuawei.Show()

	fmt.Println("---------------")
	var keHuawei Phone
	keHuawei = NewKeDecorator(huawei)
	keHuawei.Show()

	fmt.Println("--------------")
	var keMoHuawei Phone
	keMoHuawei = NewMoDecorator(keHuawei)
	keMoHuawei.Show()

}
