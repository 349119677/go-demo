package main

import (
"fmt"
"sync"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func(argvs ...interface{}), argvs ...interface{}) {
	w.Add(1)
	go func() {
		cb(argvs...)
		w.Done()
	}()
}

type MyStruct struct {
	Age  int
	Name string
	Sex  bool
}

func GetAge(argvs ...interface{}) {
	age := argvs[0].(int)
	my := argvs[1].(*MyStruct)
	my.Age = age
	fmt.Println("age done")
}

func GetName(argvs ...interface{}) {
	name := argvs[0].(string)
	my := argvs[1].(*MyStruct)
	my.Name = name
	fmt.Println("name done")
}

func GetSex(argvs ...interface{}) {
	my := argvs[0].(*MyStruct)
	my.Sex = true
	fmt.Println("sex done")
}

func main() {
	var wg WaitGroupWrapper
	var my MyStruct
	wg.Wrap(GetAge, 10, &my)
	wg.Wrap(GetName, "test", &my)
	wg.Wrap(GetSex, &my)
	wg.Wait()
	fmt.Println(my)
}