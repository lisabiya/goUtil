package rxGo

import (
	"errors"
)

func Create(task func(NextHandler)) (observer *Observable) {
	observer = &Observable{}
	observer.task = task
	return
}

func (observable *Observable) Subscribe(observer Observer) {
	observable.observer = &observer
	observable.done = false
	go func() {
		defer func() {
			// 获取异常信息
			if rsp := recover(); rsp != nil {
				observable.retryTimes = observable.times
				err, ok := rsp.(error)
				if ok {
					observable.OnError(err)
				} else {
					str, ok := rsp.(string)
					if ok {
						observable.OnError(errors.New(str))
					} else {
						observable.OnError(errors.New("未知错误"))
					}
				}
				observable.OnComplete()
			}
		}()
		//
		observable.isTimeout()
		observable.run()
	}()
}

func (observable *Observable) OnNext(event *Event) {
	if observable.done {
		return
	}
	if observable.observer != nil {
		observable.observer.OnNext(event)
	}
}

//程序错误
func (observable *Observable) OnError(err error) {
	if observable.done {
		return
	}
	if observable.observer != nil {
		observable.observer.OnError(err)
	}
	observable.retryOnErr()
}

//程序完成回调，由调用者自行控制
func (observable *Observable) OnComplete() {
	observable.done = true
	if observable.observer != nil {
		observable.observer.OnComplete()
	}
}
