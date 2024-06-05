// @Author huzejun 2024/6/6 1:07:00
package main

import "fmt"

// -----抽象层----
type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractPear interface {
	ShowPear()
}

// 抽象的工厂
type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

//-----实现层----
/* 中国产品族 */
type ChinaApple struct {
}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("中国苹果")
}

type ChinaBanana struct {
}

func (ca *ChinaBanana) ShowBanana() {
	fmt.Println("中国香蕉")
}

type ChinaPear struct {
}

func (ca *ChinaPear) ShowPear() {
	fmt.Println("中国梨")
}

type ChinaFactory struct {
}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(ChinaApple)
	return apple
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(ChinaBanana)
	return banana
}

func (cf *ChinaFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(ChinaPear)
	return pear
}

/* 日本产品族 */
type JapanApple struct {
}

func (jb *JapanApple) ShowApple() {
	fmt.Println("日本苹果")
}

type JapanBanana struct {
}

func (jb *JapanBanana) ShowBanana() {
	fmt.Println("日本香蕉")
}

type JapanPear struct {
}

func (jb *JapanPear) ShowPear() {
	fmt.Println("日本梨")
}

type JapanFactory struct {
}

func (jf *JapanFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(ChinaApple)
	return apple
}

func (jf *JapanFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(ChinaBanana)
	return banana
}

func (jf *JapanFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(ChinaPear)
	return pear
}

//针对产品族进行添加，是符合开闭原则的
//针对产品、等级、结构进行添加，不符合开闭原则

// -----逻辑层----
func main() {
	//需求1：需要中国的苹果，香蕉、梨
	//1-创建一个中国工厂
	var cFac AbstractFactory
	cFac = new(ChinaFactory)

	//2-生产苹果
	var cApple AbstractApple
	cApple = cFac.CreateApple()
	cApple.ShowApple()

	//3-生产香蕉
	var cBanana AbstractBanana
	cBanana = cFac.CreateBanana()
	cBanana.ShowBanana()

	//4-生产梨
	var cPear AbstractPear
	cPear = cFac.CreatePear()
	cPear.ShowPear()

	//需求2：中国
}
