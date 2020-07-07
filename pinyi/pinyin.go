package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	//printType()

	//pinyinTest()

	//fmt.Println(createPinyin("nihao"))

	//testRangeData()

	TestStructReflect()
}
func pinyinTest() {
	hans := "施氏，嗜狮，誓食十狮，氏时时适市视狮。十时，适十狮适市；是时，适施氏适市。氏视是十狮，恃矢势，使是十狮逝世，氏拾是十狮尸，适氏石室。石室湿，氏拭室。氏始试食十狮尸，食时，始识是十狮尸实十石狮。试释是事"

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

func TestStructReflect() {
	type CreateProcessInstanceReq struct {
		AgentId          int64  `json:"agent_id,omitempty"`
		ProcessCode      string `json:"process_code,omitempty"`
		OriginatorUserId string `json:"originator_user_id,omitempty"`
		DeptId           int64  `json:"dept_id"`
		Approvers        string `json:"approvers"`
	}

	CreateProcessInstance(CreateProcessInstanceReq{
		AgentId:          100,
		ProcessCode:      "c2704922-044f-4369-8e75-9e0770d341f4",
		OriginatorUserId: "PROC-779A76B6-85C3-4049-A32A-4E12C73CFBFC",
		DeptId:           132880747,
		Approvers:        "王开勋",
	})
}

func CreateProcessInstance(createParams interface{}) {
	params := make(map[string]string)
	typ := reflect.TypeOf(createParams)
	val := reflect.ValueOf(createParams)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i).Interface()
		v := ""
		tname := field.Type.Name()
		if value != nil {
			if tname == "string" {
				v = value.(string)
			} else if tname == "int" {
				v = strconv.Itoa(value.(int))
			} else if tname == "int64" {
				v = strconv.FormatInt(value.(int64), 10)
			} else {
				v = "err"
			}
		}
		if v != "" {
			params[field.Tag.Get("json")] = v
		}
	}
	fmt.Println(params)
	return
}
