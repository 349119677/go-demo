package routine

import (
	"fmt"
	"time"
)

func sample1() {
	go fmt.Println("1")
	fmt.Println("2")
}

func sample2() {
	var i = 3
	go func(a int) {
		fmt.Println(a)
		fmt.Println("1")
	}(i)
	fmt.Println("2")
}

func sample3() {
	var i = 3
	go func(a int) {
		fmt.Println(a)
		fmt.Println("1")
	}(i)
	fmt.Println("2")
	time.Sleep(time.Second * 1)
}
// 基于主线程，子线程执行时间必须小于主线程的生命周期
func sample4() {
	var i = 3
	go func(a int) {
		time.Sleep(time.Second * 2)
		fmt.Println(a)
		fmt.Println("1")
	}(i)
	fmt.Println("2")
	time.Sleep(time.Second * 1)
}
