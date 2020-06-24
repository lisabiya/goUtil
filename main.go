package main

import (
	"fmt"
	"golang.org/x/net/proxy"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("启动！")
	//socket()
	//udpSocket()

	//dialogTcp()
	//dialogUdp()
	//test()

	done := make(chan bool)
	for i := 1; i < 50; i++ {
		fmt.Println("sendRequest")
		time.Sleep(1000 * time.Millisecond)
		go dialogSocks5()
	}
	<-done
	//dialogSocks5()
}

func socket() {
	service := ":3333"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		fmt.Println("request: 1")
		conn, err := listener.Accept()
		fmt.Println("request: 2")
		if err != nil {
			continue
		}
		go handleRequest(conn)

	}
}

func handleRequest(conn net.Conn) {
	var r = make([]byte, 16)
	count, _ := conn.Read(r)
	fmt.Println("response: ", string(r[0:count]))

	daytime := time.Now().String()
	var response = "当前时间：" + daytime + ",请求成功"
	time.Sleep(2 * time.Second)
	_, _ = conn.Write([]byte(response)) // don't care about return value
	_ = conn.Close()                    // we're finished with this client
}

func udpSocket() {
	service := ":8804"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	listener, err := net.ListenUDP("udp4", udpAddr)
	checkError(err)
	for {
		var buf = make([]byte, 16)
		nr, clientAddr, err := listener.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("udp read:", err)
			if strings.Contains(err.Error(), "use of closed") {
				break
			} else {
				continue
			}
		}
		fmt.Println("request: ", clientAddr, string(buf[:nr]))
		var response = "请求成功啦啦啦"
		_, _ = listener.WriteToUDP([]byte(response), clientAddr) // don't care about return value
		_ = listener.Close()                                     // we're finished with this client
	}
}

func dialogTcp() {
	//使用Dial建立连接
	conn, err := net.Dial("tcp", "192.168.1.1:3333")
	if err != nil {
		fmt.Println("error dialing", err.Error())
		return
	}
	defer conn.Close()
	msg := "雷猴啊"

	_, err = io.WriteString(conn, msg)

	if err != nil {
		fmt.Println("write string failed", err)
		return
	}

	buf := make([]byte, 4096)

	for {
		count, err := conn.Read(buf)
		if err != nil {
			break
		}
		fmt.Println("返回值：", string(buf[0:count]))
	}

}

func dialogUdp() {
	//DialUDP
	uAddr, err := net.ResolveUDPAddr("udp4", "192.168.1.1:8804")
	conn, err := net.DialUDP("udp4", nil, uAddr)
	if err != nil {
		fmt.Println("error dialing", err.Error())
		return
	}
	defer conn.Close()
	//写入数据
	msg := "雷猴啊"
	_, err = io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("write string failed", err)
		return
	}

	buf := make([]byte, 4096)
	_ = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	count, _, err := conn.ReadFromUDP(buf)
	_ = conn.Close()
	if err != nil {
		fmt.Println("获取返回数据失败!", err)
	}
	fmt.Println(string(buf[0:count]))

}

func checkError(err error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

/**
 * socks客户端
 */
func dialogSocks5() {
	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", "192.168.1.1:3333", nil, proxy.Direct)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}

	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport, Timeout: 10 * time.Second}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial
	//httpTransport.DialContext = func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
	//	return dialer.Dial(network, addr)
	//}

	done := make(chan bool)
	if resp, err := httpClient.PostForm("http://175.24.128.25:8090/api/account/getAllUser", nil); err != nil {
		log.Fatalln(err)
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("回调\n %s\n", body)
	}
	<-done
}
