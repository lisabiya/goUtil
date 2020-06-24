package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"reflect"
	"strings"
)

func main() {
	//printType()

	//pinyinTest()

	//fmt.Println(createPinyin("nihao"))

	testRangeData()
}
func pinyinTest() {
	hans := "中国人"

	// 默认
	a := pinyin.NewArgs()
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhong] [guo] [ren]]

	// 包含声调
	a.Style = pinyin.Tone
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhōng] [guó] [rén]]

	// 声调用数字表示
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zho1ng] [guo2] [re2n]]

	// 开启多音字模式
	a = pinyin.NewArgs()
	a.Heteronym = true
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhong zhong] [guo] [ren]]
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zho1ng zho4ng] [guo2] [re2n]]

	fmt.Println(pinyin.LazyPinyin(hans, pinyin.NewArgs()))
	// [zhong guo ren]

	fmt.Println(pinyin.Convert(hans, nil))
	// [[zhong] [guo] [ren]]

	fmt.Println(pinyin.LazyConvert(hans, nil))
	// [zhong guo ren]

}

func printType() {
	var res map[string]interface{}
	var ss = "nihao"
	fmt.Println(reflect.TypeOf(res).Name())
	fmt.Println(reflect.TypeOf(res).String() == "map[string]interface {}")
	fmt.Println(reflect.TypeOf(ss).Name())
	fmt.Println(reflect.TypeOf(ss).String())

	var data interface{} = "great"
	if data, ok := data.(int); ok {
		fmt.Println("[is an int] value =>", data)
	} else {
		fmt.Println("[not an int] value =>", data)
		//prints: [not an int] value => 0 (not "great")
	}
}

func testRangeData() {
	var word = getRangeData(map[string]interface{}{
		"aa": "nis",
		"bb": 12,
	}, []string{"cc", "bb", "aa"})
	fmt.Println(word)

	fmt.Printf("%.0f", 12.00)
}

func getRangeData(res map[string]interface{}, keys []string) (word string) {
	for _, key := range keys {
		var s = res[key]
		msg, ok := s.(string)
		if ok {
			fmt.Println(msg)
		} else {
			fmt.Println("fail")
		}

		if s == nil {
			continue
		}

		switch reflect.TypeOf(s).Kind() {
		case reflect.String:
			word = s.(string)
			break
		case reflect.Int:
			word = fmt.Sprintf("%d", s.(int))
			break
		case reflect.Int8:
			word = fmt.Sprintf("%d", s.(int8))
			break
		case reflect.Int16:
			word = fmt.Sprintf("%d", s.(int16))
			break
		case reflect.Int32:
			word = fmt.Sprintf("%d", s.(int32))
			break
		case reflect.Int64:
			word = fmt.Sprintf("%d", s.(int64))
			break
		case reflect.Float32:
			word = fmt.Sprintf("%f", s.(float32))
			break
		case reflect.Float64:
			word = fmt.Sprintf("%f", s.(float64))
			break
		}
	}
	return
}

func createPinyin(name string) string {
	var py = pinyin.LazyPinyin(name, pinyin.NewArgs())
	if len(py) == 0 {
		return name
	}
	return strings.Join(py, "")
}
