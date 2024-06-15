// @Author huzejun 2024/6/16 1:16:00
package main

import "fmt"

// ------- 抽象层 -------
// 抽象的观察者
type Listener interface {
	OnTeacherComming() //观察者得到通知后要触发具体动作
}

type Notifier interface {
	AddListenner(listener Listener)
	RemoveListener(listener Listener)
}

// ------- 实现层 -------
// 观察者具体的学生
type StuZhang3 struct {
	Badthing string
}

func (s *StuZhang3) OnTeacherComming() {
	fmt.Println("张3 停止 ", s.Badthing)
}

func (s *StuZhang3) DoBadthing() {
	fmt.Println("张3 正在", s.Badthing)
}

type StuZhao4 struct {
	Badthing string
}

func (s *StuZhao4) OnTeacherComming() {
	fmt.Println("赵4 停止 ", s.Badthing)
}

func (s *StuZhao4) DoBadthing() {
	fmt.Println("赵4 正在", s.Badthing)
}

type StuWang5 struct {
	Badthing string
}

func (s *StuWang5) OnTeacherComming() {
	fmt.Println("王5 停止 ", s.Badthing)
}

func (s *StuWang5) DoBadthing() {
	fmt.Println("王5 正在", s.Badthing)
}

// 编统治者班长
type ClassMonitor struct {
	listenerList []Listener //需要通知的全部观察者集合

}

func (m *ClassMonitor) AddListenner(listener Listener) {
	m.listenerList = append(m.listenerList, listener)
}

func (m *ClassMonitor) RemoveListener(listener Listener) {
	for index, l := range m.listenerList {
		//找到要删除的元素位置
		if listener == l {
			//将删除的元素的前后位置链接起来
			m.listenerList = append(m.listenerList[:index], m.listenerList[index+1:]...)
			break
		}
	}
}

func (m *ClassMonitor) Notify() {
	for _, listener := range m.listenerList {
		listener.OnTeacherComming() //多态现象
	}
}

//------- 抽象层 -------

// 业务逻辑层
func main() {
	s1 := &StuZhang3{
		Badthing: "抄作业",
	}

	s2 := &StuZhao4{
		Badthing: "玩游戏",
	}

	s3 := &StuWang5{
		Badthing: "看赵4玩游戏",
	}

	classMonitor := new(ClassMonitor)
	classMonitor.AddListenner(s1)
	classMonitor.AddListenner(s2)
	classMonitor.AddListenner(s3)

	fmt.Println("上课了，但是老师没有到，学生们都在忙自己的事...")
	s1.DoBadthing()
	s2.DoBadthing()
	s3.DoBadthing()

	fmt.Println("这时候老师来了，班长给学生们使用了眼神er,...")
	classMonitor.Notify()
}
