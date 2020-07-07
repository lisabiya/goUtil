package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"time"
)

func main() {
	rxGo5()
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

func rxGO6(day string) {
	Create(func(next NextHandler) {
		next.OnError(errors.New("失败"))
		next.OnNext(&Event{Data: true})
	}).SetRetry(3).Subscribe(Observer{
		OnNext: func(event *Event) {
			fmt.Println(day, "成功")
		},
		OnError: func(err error) {
			fmt.Println(day, "失败")
		},
	})
}
