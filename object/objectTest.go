package main

import (
	"fmt"
	"reflect"
)

func main() {
	//testDuck()
	test()
}

func test() {
	var s = Child{Name: "chishis"}
	fmt.Println("child类型", reflect.TypeOf(s))
	s.RegisterChild(&s)
	fmt.Println("InitState", s.InitState())
	s.PrintChild("nih")
	println(&s)
	println(s.Name)
	println(s.name)
}

type callBack interface {
	InitState() string
	RegisterChild(child ChildIml)
}

type ChildIml interface {
	PrintChild(msg string) string
}

/**************Base****************/
type Base struct {
	name     string
	childIml ChildIml
}

func (b *Base) InitState() string {
	b.name = "Basename"
	println(b)

	if b.childIml != nil {
		return b.childIml.PrintChild("msg")
	}
	return "Base"
}

func (b *Base) RegisterChild(child ChildIml) {
	fmt.Println("RegisterChild", b.name)
	b.childIml = child
}

/**************Child****************/

type Child struct {
	Base
	Name string
}

func (c *Child) PrintChild(msg string) string {
	fmt.Println("PrintChildname", c.name)
	fmt.Println("PrintChildName", c.Name)
	c.Name = "新测试"
	println("**")
	println(&c)
	println("**")

	fmt.Println(c.Base)
	return msg + "测试"
}
