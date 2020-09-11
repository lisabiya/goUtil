package main

import "fmt"

func main() {
	var s = Child{}
	fmt.Println(s.GetReturn())
}

type callBack interface {
	GetReturn() string
}

type Base struct {
}

func (Base) GetReturn() string {
	println("Base1")
	return "Base"
}

type Child struct {
	Base
}

//func (Child) GetReturn() string {
//	println("Child1")
//	return "Child"
//}
