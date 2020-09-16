package main

/*
#cgo LDFLAGS: -L . -lSKP_SILK_SDK
#include <stdio.h>
#include <stdlib.h>
#include "SKP_Silk_SDK_API.h"
#include "Decoder.h"
static void print_usage() {
     printf("********** Silk Decoder (Fixed Point) v %s ********************\n", SKP_Silk_SDK_get_version());
}


*/
import "C"
import (
	"fmt"
	"luci/transcoder/ffmpeg"
	"reflect"
	"unsafe"
)

func main() {
	test("cgo/testVoice/2020_08_18_09_25_06_769.silk", "cgo/testVoice/newVoice.pcm")
	//testCmd()
}

const voicePath string = "cgo/testVoice/"

func test(inputPath, outputPath string) {
	inputPathC := C.CString(inputPath)
	outPathC := C.CString(outputPath)
	var s = C.Decoder(inputPathC, outPathC) //调用C函数
	fmt.Println(s)
	C.free(unsafe.Pointer(inputPathC))
	C.free(unsafe.Pointer(outPathC))
	//ffmpeg pcm转码mp3
	transPcmToMp3(outputPath, "cgo/testVoice/newM.mp3")
}

func transPcmToMp3(inputPath, OutputPath string) {
	format := "s16le"
	overwrite := true
	audioCodec := "pcm_s16le"
	audioChannels := 2
	audioRate := 12000
	opts := ffmpeg.Options{
		Overwrite:     &overwrite,
		OutputFormat:  &format,
		AudioChannels: &audioChannels,
		AudioRate:     &audioRate,
		AudioCodec:    &audioCodec,
	}
	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath: "cgo/ffmpeg",
	}
	progress, err := ffmpeg.
		New(ffmpegConf).
		Input(inputPath).
		Output(OutputPath).
		WithOptions(opts).
		Start(opts)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(reflect.TypeOf(progress))
	}
}
