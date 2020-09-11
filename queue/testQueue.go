package queue

import "fmt"

func TestQueue() {
	var que = New()

	for i := 0; i < 100; i++ {
		que.Push(i)
	}

	for i := 0; i < 80; i++ {
		fmt.Println(que.Pop())
	}

}
