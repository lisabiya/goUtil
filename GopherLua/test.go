package main

import (
	"fmt"
	"github.com/lisabiya/GopherLua"
	"github.com/lisabiya/GopherLua/goTool"
	"github.com/parnurzeal/gorequest"
	lua "github.com/yuin/gopher-lua"
)

const metatableName = "request_metatable"

type CustomModule struct {
}

func (CustomModule) RegisterType(L *lua.LState) {
	mt := L.NewTypeMetatable(metatableName)
	//声明全局对象
	L.SetGlobal("httprequest", mt)
	//添加拓展函数
	L.SetField(mt, "get", L.NewFunction(getSimple))
}

func (CustomModule) Close() {

}

func getSimple(L *lua.LState) int {
	var request = L.CheckTable(1)
	var requestMap, ok = goTool.TransLuaValue2Map(request).(map[string]interface{})
	if ok {
		_, body, _ := gorequest.New().
			Get(requestMap["url"].(string)).
			Query(requestMap["query"]).End()
		L.Push(lua.LNumber(0))
		L.Push(lua.LString(body))
		return 2
	} else {
		L.Push(lua.LNumber(1))
		L.Push(lua.LString("参数转map失败"))
		return 2

	}
}

//测试go中跑lua代码
func main() {
	gopherLua := GopherLua.NewState()
	gopherLua.Register(CustomModule{})
	err := gopherLua.DoFile("GopherLua/test.lua")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = gopherLua.ExecuteFunc("TestHttp", 1, lua.LString("测试参数"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(goTool.TransLuaValue2Map(gopherLua.State.Get(-1)))
}
