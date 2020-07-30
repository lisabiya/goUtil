package main

import (
	"github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	defer L.Close()
	_ = L.RegisterModule("GO_Util", GoUtil)

	if err := L.DoFile(`lua/run.lua`); err != nil {
		panic(err)
	}

}

var GoUtil = map[string]lua.LGFunction{
	"printMe": printMe,
}

func printMe(L *lua.LState) int {
	src := L.ToString(1)
	println("goStr", src)
	L.Push(lua.LNumber(186))
	L.Push(lua.LString("来自GO的问候"))
	return 2
}
