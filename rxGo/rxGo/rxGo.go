package rxGo

import (
	"errors"
	"fmt"
)

func Create(task func(NextHandler)) (observer *Observable) {
	observer = &Observable{}
	observer.task = task
	return
}

func (observable *Observable) Subscribe(observer Observer) {
	observable.observer = &observer
	observable.wg = cwgInstance()
	observable.wg.Add(1)
	go func() {
		defer func() {
			// 获取异常信息
			if err := recover(); err != nil {
				fmt.Println(err)
				observable.OnError(errors.New("获取异常信息"))
			}
		}()
		//
		observable.isTimeout()

		observable.run()
		//
		observable.wg.Wait()
		observable.OnComplete()
	}()
}

func (observable *Observable) OnNext(event *Event) {
	if observable.observer != nil {
		observable.observer.OnNext(event)
	}
	observable.wg.Done()
}
func (observable *Observable) OnError(err error) {
	if observable.observer != nil {
		observable.observer.OnError(err)
	}
	observable.retryOnErr()
}

//当程序走完时
func (observable *Observable) OnComplete() {
	if observable.observer != nil {
		observable.observer.OnComplete()
	}
}
