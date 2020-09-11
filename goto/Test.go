package main

import (
	"encoding/json"
	"github.com/openzipkin/zipkin-go/reporter/http"
	"net"
	"time"
)

func main() {
	testChan()
}

func testChan() {
	ch := make(chan string, 255)

	go func() {
		ch <- "test"
		time.Sleep(time.Second * 3)
		println("test1")
	}()
	<-ch

	println("test2")
	reporter := http.NewReporter("http://zipkinhost:9411/api/v2/spans")
	defer reporter.Close()

	time.Sleep(time.Second * 5)
}

func testGoto() {
	// 消息列表
	ch := make(chan string, 255)
	// 定时从ch中取出所有数据
	ch <- "你好"
	println("开始")
	for {
		select {
		case msg := <-ch:
			println(msg)
		// some actions
		case <-time.After(time.Second):
			// timed out
			println("After")
			goto PrintIt
			//default:
			//	println("default")
			//	ch <- "错误"
			//	time.Sleep(time.Second * 3)
		}
	}

PrintIt:
	{
		println("测试数据")
	}
	var ipArr []net.IP
	for len(ipArr) < 10 {
		iprecords, _ := net.LookupIP("zproxy.lum-superproxy.io")
		ipArr = append(ipArr, iprecords...)
	}
	ips, _ := json.Marshal(ipArr)
	println(string(ips))
	println("ceshi")

}
