package rxGo

import (
	"sync"
	"time"
)

type (
	Event struct {
		Data interface{}
	}
	NextHandler interface {
		OnNext(*Event)
		OnError(err error)
		OnComplete()
	}

	Observable struct {
		observer   *Observer
		wg         *sync.WaitGroup
		times      int
		delayTime  int
		retryTimes int
		task       func(NextHandler)
	}

	Observer struct {
		OnNext     func(*Event)
		OnError    func(error)
		OnComplete func()
	}
)

func Create(task func(NextHandler)) (observer *Observable) {
	observer = &Observable{}
	observer.task = task
	return
}

/**************Observable***************/
/**
 *	@params times 失败重试次数
 */
func (observable *Observable) SetRetry(times int) *Observable {
	observable.times = times
	return observable
}

/**
 *	@params time 延时时间 秒
 */
func (observable *Observable) Timer(delayTime int) *Observable {
	observable.delayTime = delayTime
	return observable
}

func (observable *Observable) retryOnErr() {
	if observable.retryTimes < observable.times {
		observable.retryTimes = observable.retryTimes + 1
		observable.run()
	} else {
		observable.wg.Done()
	}

}

func (observable *Observable) Subscribe(observer Observer) {
	observable.observer = &observer
	observable.wg = &sync.WaitGroup{}
	observable.wg.Add(1)
	observable.run()
	observable.wg.Wait()
	observable.OnComplete()
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

func (observable *Observable) run() {
	go func() {
		if observable.delayTime > 0 {
			time.Sleep(time.Second * time.Duration(observable.delayTime))
		}
		observable.task(observable)
	}()
}
