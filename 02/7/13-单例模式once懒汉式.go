// @Author huzejun 2024/6/10 22:00:00
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type singelton struct{}

var instance *singelton

func GetInstance() *singelton {

	once.Do(func() {
		instance = new(singelton)
	})
	return instance
}

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
