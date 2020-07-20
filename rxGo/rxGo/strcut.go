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
		wg         *CustomWaitGroup
		times      int
		delayTime  int
		timeout    int
		retryTimes int
		task       func(NextHandler)
	}

	Observer struct {
		OnNext     func(*Event)
		OnError    func(error)
		OnComplete func()
	}

	CustomWaitGroup struct {
		Count int
		Wg    *sync.WaitGroup
	}
)

func (observable *Observable) run() {
	if observable.delayTime > 0 {
		time.Sleep(time.Second * time.Duration(observable.delayTime))
	}
	observable.task(observable)

}

func (observable *Observable) retryOnErr() {
	if observable.retryTimes < observable.times {
		observable.retryTimes = observable.retryTimes + 1
		observable.run()
	} else {
		observable.wg.Done()
	}
}
func (observable *Observable) isTimeout() {
	if observable.timeout != 0 {
		go func() {
			time.Sleep(time.Second * time.Duration(observable.timeout))
			observable.wg.Done()
		}()
	}
}

func cwgInstance() *CustomWaitGroup {
	return &CustomWaitGroup{
		Count: 0,
		Wg:    &sync.WaitGroup{},
	}
}

func (wg *CustomWaitGroup) Add(delta int) {
	wg.Count = wg.Count + delta
	wg.Wg.Add(delta)
}

func (wg *CustomWaitGroup) Done() {
	if wg.Count > 0 {
		wg.Add(-wg.Count)
	}
}

func (wg *CustomWaitGroup) Wait() {
	wg.Wg.Wait()
}

//
/**************操作符***************/
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

/**
 *	@params timeout 过时时间 秒
 */
func (observable *Observable) Timeout(timeout int) *Observable {
	observable.timeout = timeout
	return observable
}
