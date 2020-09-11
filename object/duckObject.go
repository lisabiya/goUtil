package main

import "fmt"

func testDuck() {
	var chicken = Chicken{}
	DoDuck(chicken)
}

func DoDuck(d Duck) {
	d.Quack()
	d.DuckGo()
}

type Duck interface {
	Quack()  // 鸭子叫
	DuckGo() // 鸭子走
}

type Chicken struct {
}

func (c Chicken) IsChicken() bool {
	fmt.Println("我是小鸡")
	return true
}

func (c Chicken) Quack() {
	fmt.Println("嘎嘎")
}

func (c Chicken) DuckGo() {
	fmt.Println("大摇大摆的走")
}
