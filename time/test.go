package main

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"time"
)

//限流器

func main() {
	waitPlan()
	//allowPlan()
	//reservePlan()
}

func waitPlan() {
	l := rate.NewLimiter(1, 1)
	c, _ := context.WithCancel(context.TODO())
	fmt.Println(l.Limit(), l.Burst())
	var count = 0
	for {
		count++
		fmt.Println("wait1", count)
		l.Wait(c)
		//time.Sleep(1000 * time.Millisecond)
		fmt.Println(time.Now().Format("2016-01-02 15:04:05"))
	}
}

func allowPlan() {
	l := rate.NewLimiter(10, 1)
	var count = 0
	for {
		count++
		fmt.Println("allow1", count)
		if l.AllowN(time.Now(), 2) {
			fmt.Println("allow2", count)
			fmt.Println(time.Now().Format("2016-01-02 15:04:05.000"))
		} else {
			fmt.Println("allow3", count)
			time.Sleep(3 * time.Second)
		}
	}
}

func reservePlan() {
	l := rate.NewLimiter(1, 3)
	for {
		r := l.ReserveN(time.Now(), 1)
		time.Sleep(r.Delay())
		fmt.Println(time.Now().Format("2016-01-02 15:04:05.000"))
	}
}
