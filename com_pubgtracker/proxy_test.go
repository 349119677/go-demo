package com_pubgtracker

import (
	"testing"
	"fmt"
)

func TestDo(t *testing.T) {
	// 手动init
	userAgent,apiKey := Init()
	fmt.Println(apiKey)
	fmt.Println("初始化后池子状态为：")
	for _, value := range userAgent.UserAgentItemList {
		fmt.Println(value)
	}
	fmt.Println("执行获取")
	item := userAgent.Get()
	fmt.Println("获取到的对象为", item)
	fmt.Println("获取对象后池子状态为：")
	for _, value := range userAgent.UserAgentItemList {
		fmt.Println(value)
	}

	fmt.Println("获取到的对象为-----------", userAgent.Get())
	fmt.Println("获取到的对象为-----------", userAgent.Get())

}
