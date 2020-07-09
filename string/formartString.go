package main

import (
	"fmt"
	"reflect"
)

func main() {
	var result = formatString([]interface{}{"name"})
	fmt.Println("结果：", result)
}

func getRangeData(res map[string]interface{}, keys []string) (in interface{}) {
	for _, key := range keys {
		in = res[key]
		if in != nil {
			break
		}
	}
	return
}

func formatString(in interface{}) (word string) {
	if in == nil {
		return ""
	}
	var kind = reflect.TypeOf(in).Kind()
	switch kind {
	case reflect.String:
		word = in.(string)
		break
	case reflect.Int:
		word = fmt.Sprintf("%d", in.(int))
		break
	case reflect.Int8:
		word = fmt.Sprintf("%d", in.(int8))
		break
	case reflect.Int16:
		word = fmt.Sprintf("%d", in.(int16))
		break
	case reflect.Int32:
		word = fmt.Sprintf("%d", in.(int32))
		break
	case reflect.Int64:
		word = fmt.Sprintf("%d", in.(int64))
		break
	case reflect.Float32:
		word = fmt.Sprintf("%.0f", in.(float32))
		break
	case reflect.Float64:
		word = fmt.Sprintf("%.0f", in.(float64))
		break
	case reflect.Slice:
		var arr, ok = in.([]interface{})
		if ok && len(arr) > 0 {
			word = formatString(arr[0])
		}
		break
	}
	return
}
