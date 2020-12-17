package main

import (
	"fmt"
	socks5 "github.com/armon/go-socks5"
)

func main2() {
	sock5Server()
}

func sock5Server() {
	fmt.Println("启动成功...")
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on port 10080
	if err := server.ListenAndServe("tcp", "0.0.0.0:3333"); err != nil {
		fmt.Println(err.Error())
	}
}
