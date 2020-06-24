package main

import (
	"fmt"
	"time"
)

// 模拟动物行为的接口
type IAnimal interface {
	Eat()   //吃
	Drink() //喝
	Play()  //玩
	Enjoy() //乐
}

// 动物 所有动物的父类
type Animal struct {
	Name string
}

// 动物去实现IAnimal中描述的吃的接口
func (a *Animal) Eat() {
	fmt.Printf("%v is Eat\n", a.Name)
}

func (a *Animal) Drink() {
	fmt.Printf("%v is Drink\n", a.Name)
}
func (a *Animal) Play() {
	fmt.Printf("%v is Play\n", a.Name)
}
func (a *Animal) Enjoy() {
	fmt.Printf("%v is Enjoy\n", a.Name)
}

func check(animal IAnimal) {
	animal.Eat()
	animal.Drink()
	animal.Play()
	animal.Enjoy()
}

type Cat struct {
	Animal
}

//进程间通信
func runMe(me chan bool) {
	go func() {
		for {
			var exit = false
			select {
			case exit = <-me:
				// 如果chan1成功读到数据，则进行该case处理语句
				fmt.Println("成功读到数据")
				close(me)
				break
			default:
				fmt.Println("default")
				break
				// 如果上面都没有成功，则进入default处理流程
			}
			if exit {
				fmt.Println("break")
				break
			} else {
				fmt.Println("Sleep")
				time.Sleep(1 * time.Second)
			}
		}
		fmt.Println("finish")
	}()
	time.Sleep(6 * time.Second)
	fmt.Println("timer")
	me <- true
	time.Sleep(3 * time.Second)
}

func timer(me chan bool) {

}

func main() {
	var chan1 = make(chan bool)
	runMe(chan1)
	//var cat = Cat{
	//	Animal{Name: "cat"},
	//}
	//check(&cat)

}
