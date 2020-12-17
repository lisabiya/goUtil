package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	go HttpProxy("23.225.210.196", 10000)
	time.Sleep(time.Millisecond * 200)
	time.Sleep(time.Hour)
}

func HttpProxy(ip string, port int) {
	urli := url.URL{}
	urlproxy, _ := urli.Parse(fmt.Sprintf("http://%s:%d", ip, port))
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
	}

	rqt, err := http.NewRequest("GET", "https://myip.ipip.net/", nil)
	if err != nil {
		println(err.Error())
		return
	}
	response, err := client.Do(rqt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("http:", string(body))
	return
}
