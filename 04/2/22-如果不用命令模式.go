// @Author huzejun 2024/6/15 0:49:00
package main

import "fmt"

type Doctor struct {
}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// 病人
func main() {
	doctor := new(Doctor)
	doctor.treatNose()
	doctor.treatEye()
}
