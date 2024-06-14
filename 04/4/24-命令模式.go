// @Author huzejun 2024/6/15 1:46:00
package main

import "fmt"

// 核心计算模块,Doctor
type Doctor struct {
}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// 抽象的命令
type Command interface {
	Treat() //抽象命令接口
}

// 治疗眼睛的病单
type CommandTreatEye struct {
	doctor *Doctor
}

func (cmd *CommandTreatEye) Treat() {
	cmd.doctor.treatEye()
}

// 治疗鼻子的病单
type CommandTreatNose struct {
	doctor *Doctor
}

func (cmd *CommandTreatNose) Treat() {
	cmd.doctor.treatNose()
}

// 护士，命令的调用者
type Nurse struct {
	CmdList []Command //收集的命令集合
}

// 发关病单，发送命令的方法
func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}

	for _, cmd := range n.CmdList {
		cmd.Treat() //多态现象，调用具体的命令
		//就会调用病单已经绑定的医生的诊断方法
	}
}

// 病人 -业务层
func main() {
	//依赖病单，通过填写病单，让医生看病
	doctor := new(Doctor)

	cmdEye := CommandTreatEye{doctor}

	cmdNose := CommandTreatNose{doctor}

	//护士
	nurse := new(Nurse)
	//收集管理病单
	nurse.CmdList = append(nurse.CmdList, &cmdEye)
	nurse.CmdList = append(nurse.CmdList, &cmdNose)

	//执行病单指令
	nurse.Notify()
}
