package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

func resumeFunction() {
	L := lua.NewState()
	defer L.Close()
	lunch, err := L.LoadFile("lua/test.lua")
	if err != nil {
		fmt.Println(err)
	}
	co, _ := L.NewThread()
	st, err, values := L.Resume(co, lunch)
	if st == lua.ResumeError {
		fmt.Println("yield break(error)")
		fmt.Println(err.Error())
	}

	for i, lv := range values {
		fmt.Printf("%v : %v\n", i, lv)
	}
	if st == lua.ResumeOK {
		fmt.Println("yield break(ok)")
	}
}
