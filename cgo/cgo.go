package main

/*
#cgo LDFLAGS: -L . -lSKP_SILK_SDK
#include <stdio.h>
#include "SKP_Silk_SDK_API.h"
#include "Decoder.h"
static void print_usage() {
     printf("********** Silk Decoder (Fixed Point) v %s ********************\n", SKP_Silk_SDK_get_version());
}


*/
import "C"
import "fmt"

func main() {
	test()
}
func test() {
	//cs := C.CString("Hello World\n")
	//var pcm = C.CString("")
	//var arr = []C.char{*silk, *pcm}
	//fmt.Println(reflect.TypeOf(&pcm))
	var s = C.Decoder() //调用C函数
	fmt.Println(s)
	//C.add(2, 4)
}

func PrintCString(cs *C.char) {
	print(C.GoString((*C.char)(cs)))
}
