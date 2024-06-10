// @Author huzejun 2024/6/11 2:20:00
package main

import "fmt"

// 抽象主题
type BeautyMoman interface {
	//对男人抛媚眼
	MakeEyesWithMan()
	//和男人共度美好时光
	HappyWithMan()
}

// 具体主题
type PanJinlian struct {
}

// 对象男人抛媚眼
func (p *PanJinlian) MakeEyesWithMan() {
	fmt.Println("潘小姐对本官抛了个媚眼")
}

// 和男人共度美好时光
func (p *PanJinlian) HappyWithMan() {
	fmt.Println("潘小姐和本官共度了浪漫的约会...")
}

// 中间的代理，王婆
type WangPo struct {
	woman BeautyMoman
}

func NewProxy(woman BeautyMoman) BeautyMoman {
	return &WangPo{woman}
}

// 对象男人抛媚眼
func (w *WangPo) MakeEyesWithMan() {
	w.woman.MakeEyesWithMan()
}

// 和男人共度美好时光
func (w *WangPo) HappyWithMan() {
	w.woman.HappyWithMan()
}

// main业务逻辑-西门大官人
func main() {
	// 大官人想找金莲，让王婆来安排
	wangpo := NewProxy(new(PanJinlian))
	//王婆命令金莲抛媚眼
	wangpo.MakeEyesWithMan()
	//王婆命令金莲和西门大官人约会
	wangpo.HappyWithMan()
}
