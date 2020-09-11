package neo

import (
	"fmt"
	"reflect"
)

type Level struct {
	Number int8   `json:"number"`
	Name   string `json:"name"`
}

func CreateLevel() {
	var ws = GetWriteSession()
	if ws != nil {
		result, err := ws.Run("MATCH (n:Company) RETURN Properties(n) LIMIT 25", nil)
		if err != nil {
			return
		}

		for result.Next() {
			fmt.Println(reflect.TypeOf(result.Record().GetByIndex(0)))
			//fmt.Println("Created Item ", result.Record().GetByIndex(0).(neo4j.Node).Props())
		}
	}
}
