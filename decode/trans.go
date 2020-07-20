package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func main() {
	fmt.Println(convertOctonaryUtf8("\345\223\210\345\223\210"))
	fmt.Println(convertOctonaryUtf8("\344\275\240\345\245\275"))
	fmt.Println(convertOctonaryUtf8("[\345\217\221\345\221\206]"))

	var s = string("\032\153\120\033\123\132")
	fmt.Println(s)
	//var aa = []int8{345,217,221}
	//fmt.Println(string(aa))

	time.Sleep(time.Minute)
}

func convertOctonaryUtf8(in string) string {
	s := []byte(in)
	reg := regexp.MustCompile(`\\[0-7]{3}`)
	out := reg.ReplaceAllFunc(s,
		func(b []byte) []byte {
			fmt.Println(b)
			i, _ := strconv.ParseInt(string(b[1:]), 8, 0)
			return []byte{byte(i)}
		})
	return string(out)
}
