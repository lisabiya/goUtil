package db_module

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
	"luci/db"
)

var exports = map[string]lua.LGFunction{
	"getDB": getDB,
}

func LoadDBModule(L *lua.LState) {
	L.PreloadModule("db_module", loader)
}

func loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)

	L.SetField(mod, "_DEBUG", lua.LBool(false))
	L.SetField(mod, "_VERSION", lua.LString("0.0.0"))

	return 1
}

func getDB(L *lua.LState) int {
	var tableName = L.ToString(-1)
	rows, err := db.GetDB().Table(tableName).Rows()
	if err != nil {
		L.Push(lua.LNumber(1))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	//返回所有列
	cols, _ := rows.Columns()
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	var table = lua.LTable{}
	i := 1
	for rows.Next() {
		//rows
		//填充数据
		_ = rows.Scan(scans...)
		//每行数据
		row := make(map[string]interface{})
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		table.Insert(i, luar.New(L, row))
		i++
	}
	L.Push(lua.LNumber(0))
	L.Push(&table)
	return 2
}

func testUserTable(L *lua.LState) int {
	ud := L.NewUserData()
	//L.SetMetatable(ud, L.NewTypeMetatable())
	L.Push(ud)
	return 2
}

func TestDb() {
	db.Setup()
	rows, _ := db.GetDB().Table("t_salary").Limit(10).Rows()
	//返回所有列
	cols, _ := rows.Columns()
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	i := 0
	result := make(map[int]map[string]string)
	for rows.Next() {
		//填充数据
		rows.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		fmt.Println(row)
		result[i] = row
		i++
	}
	//fmt.Println(result)
	for k, v := range result {
		fmt.Printf("第%d行", k)
		fmt.Println(v["gentuanyouid"] + "===>" + v["title"])
	}
}
