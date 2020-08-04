package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	luasql "github.com/tengattack/gluasql/sqlite3"
	"github.com/yuin/gopher-lua"
	"luci/db"
	"luci/lua/db_module"
	"net/http"
)

func main() {
	db.Setup()
	//initRouter()
	//db_module.TestDb()

	testMetatable()
	//testSql()
}

func initRouter() {
	r := gin.Default()
	r.GET("/ping", loadLuaModule)
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}

func testMetatable() {
	luaContext := lua.NewState()
	defer luaContext.Close()
	db_module.RegisterOrmDbType(luaContext)
	if err := luaContext.DoFile("lua/run.lua"); err != nil {
		panic(err)
	}
	if err := luaContext.CallByParam(lua.P{
		Fn:      luaContext.GetGlobal("initParams"),
		NRet:    1,
		Protect: true,
	}); err != nil {
		panic(err)
		return
	}
	ret := luaContext.Get(1) // returned value
	fmt.Println(transLuaValue2Map(ret))
}

func testSql() {
	luaContext := lua.NewState()
	defer luaContext.Close()
	luaContext.PreloadModule("sqlite3", luasql.Loader)
	if err := luaContext.DoFile("lua/nativeSql/test.lua"); err != nil {
		panic(err)
	}
	if err := luaContext.CallByParam(lua.P{
		Fn:      luaContext.GetGlobal("testSql"),
		NRet:    1,
		Protect: true,
	}); err != nil {
		panic(err)
		return
	}
	ret := luaContext.Get(1) // returned value
	fmt.Println(transLuaValue2Map(ret))
}

func loadLuaModule(c *gin.Context) {
	luaContext := getDefaultGinStatus(c)
	defer luaContext.Close()
	err := luaContext.DoFile("lua/run.lua")
	if err != nil {
		c.JSON(http.StatusOK, formatError(err))
		return
	}

	if err := luaContext.CallByParam(lua.P{
		Fn:      luaContext.GetGlobal("initParams"),
		NRet:    1,
		Protect: true,
	}); err != nil {
		c.JSON(http.StatusOK, formatError(err))
		return
	}
	ret := luaContext.Get(1) // returned value

	c.JSON(http.StatusOK, formatSuccess(transLuaValue2Map(ret)))
}

func getDefaultGinStatus(c *gin.Context) *lua.LState {
	L := lua.NewState()
	var getParams = L.NewFunction(func(state *lua.LState) int {
		var key = state.ToString(-1)
		var value = c.Query(key)
		L.Push(lua.LString(value))
		return 1
	})
	L.SetGlobal("getParams", getParams)
	var postParams = L.NewFunction(func(state *lua.LState) int {
		var key = state.ToString(-1)
		var value = c.PostForm(key)
		L.Push(lua.LString(value))
		return 1
	})
	L.SetGlobal("postParams", postParams)
	return L
}

//**************************拓展*******************************************
func registerUtil(L *lua.LState) {
	L.RegisterModule("GO_Util", GoUtil)
}

var GoUtil = map[string]lua.LGFunction{
	"printMe": printMe,
}

func printMe(L *lua.LState) int {
	src := L.Get(-1)
	println("goStr", src.String())
	L.Push(lua.LNumber(186))
	L.Push(lua.LString("来自GO的问候1"))
	return 2
}

func transLuaValue2Map(value lua.LValue) interface{} {
	if value.Type() == lua.LTTable {
		var deMap = make(map[string]interface{})
		var list []interface{}
		var table = value.(*lua.LTable)
		table.ForEach(func(key lua.LValue, value lua.LValue) {
			if key.Type() == lua.LTNumber {
				list = append(list, transLuaValue2Map(value))
			} else {
				deMap[key.String()] = transLuaValue2Map(value)
			}
		})
		if len(deMap) > 0 && len(list) > 0 {
			return map[string]interface{}{
				"map":  deMap,
				"list": list,
			}
		}
		if len(deMap) > 0 {
			return deMap
		}
		if len(list) > 0 {
			return list
		}
		return deMap
	} else if value.Type() == lua.LTUserData {
		var table = value.(*lua.LUserData)
		return table.Value
	} else {
		return value
	}
}
