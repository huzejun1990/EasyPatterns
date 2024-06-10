// @Author huzejun 2024/6/10 21:04:00
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singelton struct{}

// 定义一个锁
var lock sync.Mutex

// 标记
var initialized uint32

var instance *singelton

func GetInstance() *singelton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	//如果没有，则加锁申请
	lock.Lock()
	defer lock.Unlock()

	//只有首次GetInstance()方法被调用，才会生成这个到单例的对象
	if instance == nil {
		instance = new(singelton)
		//设置标记为1
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

/*
func GetInstance() *singelton {
	//为了线程安全，增加互斥
	lock.Lock()
	defer lock.Unlock()

	//只有首次GetInstance()方法调用，才会生成这个到单例对象
	if instance == nil {
		instance = new(singelton)
		return instance
	}
	return instance
}*/

func (s *singelton) SomeThing() {
	fmt.Println("单例的某方法")
}
func main() {
	s1 := GetInstance()
	s1.SomeThing()

	s2 := GetInstance()

	if s1 == s2 {
		fmt.Println("s1 == s2")
	}
}
