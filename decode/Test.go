package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("trans.so")
	if err != nil {
		panic(err)
	}
	m, err := p.Lookup("ConvertOctonaryUtf8")
	if err != nil {
		panic(err)
	}
	res := m.(func(string) string)("\345\223\210\345\223\210")
	fmt.Println(res)
}
