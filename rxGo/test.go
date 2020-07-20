package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	rxGo4 "luci/rxGo/rxGo"
	"time"
)

func main() {
	//rxGo5()
	rxGo7()
}

func rxGo() {
	observable := rxgo.Just(4, 5, 6)().Map(func(ctx context.Context, i2 interface{}) (i interface{}, err error) {
		return i2.(int) * 10, nil
	})
	observable.DoOnNext(func(i interface{}) {
		fmt.Println(i)
	})
	observable.DoOnCompleted(func() {
		fmt.Println(time.Now().Nanosecond(), "complete1")
		time.Sleep(time.Second)
	})
	disposed, cancel := observable.Connect()
	go func() {
		time.Sleep(time.Second * 60)
		cancel()
	}()
	<-disposed.Done()

}

func rxGo1() {
	ch := make(chan rxgo.Item)
	go func() {
		ch <- rxgo.Of(1)
		ch <- rxgo.Of(2)
		ch <- rxgo.Of(3)
		close(ch)
	}()
	// Create a Connectable Observable
	observable := rxgo.FromChannel(ch, rxgo.WithPublishStrategy())

	// Create the first Observer
	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("First observer: %d\n", i)
	})

	// Create the second Observer
	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("Second observer: %d\n", i)
	})

	disposed, cancel := observable.Connect()
	go func() {
		// Do something
		time.Sleep(time.Second)
		// Then cancel the subscription
		fmt.Println("nih")
		time.Sleep(time.Second)
		cancel()
	}()
	// Wait for the subscription to be disposed
	<-disposed.Done()
	fmt.Println("ces")
}

func rxGo2() {
	observable := rxgo.Defer([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(20)
		next <- rxgo.Of(10)
	}})
	observable.DoOnNext(func(i interface{}) {
		fmt.Println("DoOnNext", i)
		time.Sleep(time.Second * 2)
	})
	observable.DoOnCompleted(func() {
		fmt.Println("complete")
		time.Sleep(time.Second * 3)
	})
	observable.DoOnError(func(err error) {

	})

	disposed, cancel := observable.Connect()
	go func() {
		time.Sleep(time.Second * 60)
		cancel()
	}()
	<-disposed.Done()
}

func rxGo3() {
	observable1 := rxgo.Create([]rxgo.Producer{func(_ context.Context, ch chan<- rxgo.Item) {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
	}})

	// First Observer
	for item := range observable1.Observe() {
		fmt.Println(item.V)
	}

	// Second Observer
	for item := range observable1.Observe() {
		fmt.Println(item.V)
	}
}

func rxGo5() {
	var me = make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		me <- 2
		me <- 6
	}()
	select {
	case num := <-me:
		fmt.Println(num)
		//close(me)
		break
	default:
		fmt.Println("终结")
		break
	}
	time.Sleep(time.Second * 4)
}

func rxGo6() {
	var num = 0

	rxGo4.Create(func(next rxGo4.NextHandler) {
		num++
		if num > 2 {
			next.OnNext(&rxGo4.Event{Data: 22})
		} else {
			next.OnError(errors.New(fmt.Sprintf("第%d次失败状态", num)))
		}
	}).Timer(1).SetRetry(3).Subscribe(rxGo4.Observer{
		OnNext: func(event *rxGo4.Event) {
			fmt.Println("成功", event.Data)
		},
		OnError: func(err error) {
			fmt.Println("失败", err.Error())
		},
		OnComplete: func() {
			fmt.Println("完成")
		},
	})
	time.Sleep(time.Minute)

	rxGo4.Create(func(handler rxGo4.NextHandler) {

	}).Subscribe(rxGo4.Observer{
		OnNext: func(event *rxGo4.Event) {

		},
		OnError: func(err error) {

		},
		OnComplete: func() {

		},
	})
}

func rxGo7() {
	//var num = 0
	//rxGo4.Create(func(next rxGo4.NextHandler) {
	//	num++
	//	if num > 2 {
	//		next.OnNext(&rxGo4.Event{Data: 22})
	//	} else {
	//		next.OnError(errors.New(fmt.Sprintf("第%d次失败状态", num)))
	//	}
	//}).Timer(1).SetRetry(2).Subscribe(rxGo4.Observer{
	//	OnNext: func(event *rxGo4.Event) {
	//		fmt.Println("成功", event.Data)
	//	},
	//	OnError: func(err error) {
	//		fmt.Println("失败", err.Error())
	//	},
	//	OnComplete: func() {
	//		fmt.Println("完成")
	//	},
	//})

	rxGo4.Create(func(handler rxGo4.NextHandler) {
		handler.OnNext(&rxGo4.Event{Data: "s"})
		time.Sleep(time.Second * 3)
		handler.OnNext(&rxGo4.Event{Data: "l"})

	}).Subscribe(rxGo4.Observer{
		OnNext: func(event *rxGo4.Event) {
			fmt.Println("OnNext", event.Data)
		},
		OnError: func(err error) {
			fmt.Println("OnError")
		},
		OnComplete: func() {
			fmt.Println("OnComplete")
		},
	})
	time.Sleep(time.Minute)
}
