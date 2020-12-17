package queue

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func TestQueue() {
	var que = New()

	//for i := 0; i < 100; i++ {
	//	que.Push(i)
	//}
	//
	//for i := 0; i < 80; i++ {
	//	fmt.Println(que.Pop())
	//}
	que.Push(1)
	que.Push(2)
	que.Push(3)
	que.Push(4)

	fmt.Println(que.Pop())
	fmt.Println(que.Pop())
	que.Push(1)
	que.Push(2)
	for i := 0; i < 10; i++ {
		fmt.Println(que.Pop())
	}

	cronTab := cron.New()
	if err := cronTab.AddFunc("*/5 * * * *", printIt); err != nil {
		fmt.Printf("启动失败上传队列服务失败:%s", err)
		//os.Exit(-1)
	}
	cronTab.Start()
	time.Sleep(time.Minute * 20)
}

func printIt() {
	fmt.Println("测候1")
	time.Sleep(time.Second * 20)
	fmt.Println("测候2")
}
