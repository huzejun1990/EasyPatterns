// @Author huzejun 2024/6/15 1:54:00
package main

import "fmt"

/*
练习：

	联想路边撸串烧烤场景
	有烤羊肉，烤鸡翅命令，有烤串师傅，和服务员MM
	根据命令模式，设计烤串场景
*/

type Cooker struct {
}

func (c *Cooker) MakeChuaner() {
	fmt.Println("烤串师傅烤了羊肉串儿")
}

func (c Cooker) MakeChicken() {
	fmt.Println("烤串是师傅烤了鸡肉串儿")
}

// 抽象命令
type Command interface {
	Make()
}

type CommandCookChicken struct {
	cooker *Cooker
}

func (cmd *CommandCookChicken) Make() {
	cmd.cooker.MakeChuaner()
}

type CommandCookChuaner struct {
	cooker *Cooker
}

func (cmd *CommandCookChuaner) Make() {
	cmd.cooker.MakeChuaner()
}

// 命令的调用者
type WaiterMM struct {
	CmdList []Command
}

func (w WaiterMM) Notify() {
	if w.CmdList == nil {
		return
	}

	for _, cmd := range w.CmdList {
		cmd.Make()
	}
}

func main() {
	cooker := new(Cooker)
	cmdChicken := CommandCookChicken{cooker}
	cmdChuaner := CommandCookChuaner{cooker}

	mm := new(WaiterMM)
	mm.CmdList = append(mm.CmdList, &cmdChicken)
	mm.CmdList = append(mm.CmdList, &cmdChuaner)

	mm.Notify()

}
