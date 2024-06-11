// @Author huzejun 2024/6/12 4:03:00
package main

import "fmt"

// 适配目标 抽象的技能
type Attack interface {
	Fight()
}

// 具体的技能
type Dabaojian struct {
}

func (d *Dabaojian) Fight() {
	fmt.Println("使用了大保健技能，将敌人击杀...")
}

type Hero struct {
	Name   string
	attack Attack //攻击方式
}

func (h *Hero) Skill() {
	fmt.Println(h.Name, "使用了技能.")
	h.attack.Fight() //使用具体的战斗方式
}

// 适配者
type PowerOff struct {
}

func (p *PowerOff) ShutDown() {
	fmt.Println("电脑即将关机...")
}

// 适配器
type Adapter struct {
	poweroff *PowerOff
}

func (a *Adapter) Fight() {
	a.poweroff.ShutDown()
}

func NewAdapter(p *PowerOff) *Adapter {
	return &Adapter{p}
}

func main() {
	/*	gailun := Hero{"盖伦", new(Dabaojian)}
		gailun.Skill()*/

	gailun := Hero{"盖伦", NewAdapter(new(PowerOff))}
	gailun.Skill()
}
