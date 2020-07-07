package main

import "time"

type (
	Event struct {
		Data interface{}
	}
	NextHandler interface {
		OnNext(*Event)
		OnError(err error)
	}

	Observable struct {
		observer   *Observer
		times      int
		retryTimes int
		task       func(NextHandler)
	}

	Observer struct {
		OnNext  func(*Event)
		OnError func(error)
	}
)

func Create(task func(NextHandler)) (observer *Observable) {
	observer = &Observable{}
	observer.task = task
	return
}

/**************Observable***************/
func (observable *Observable) SetRetry(times int) *Observable {
	observable.times = times
	return observable
}

func (observable *Observable) retryOnErr() {
	if observable.retryTimes < observable.times {
		time.Sleep(time.Second * 2)
		observable.retryTimes = observable.retryTimes + 1

		observable.run()
	}
}

func (observable *Observable) Subscribe(observer Observer) {
	observable.observer = &observer
	observable.run()
}

func (observable *Observable) OnNext(event *Event) {
	if observable.observer != nil {
		observable.observer.OnNext(event)
	}
}
func (observable *Observable) OnError(err error) {
	if observable.observer != nil {
		observable.observer.OnError(err)
		observable.retryOnErr()
	}
}

func (observable *Observable) run() {
	go observable.task(observable)
}
