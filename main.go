package main

import (
	"encoding/json"
	"fmt"
	"go-test/dbt"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	ID       int    `db:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

type DirItem struct {
	Item
	Dir      string `json:"dir"`
	TestName string
}

func (DirItem) TableName() string {
	return "device_shutdown_operation"
}

func main() {
	db, err := dbt.InitDB("./test.db")
	if err != nil {
		panic(err)
	}
	err = dbt.StructToSQLCreateTable(db, DirItem{})
	if err != nil {
		panic(err)
	}

	obj := DirItem{Dir: "123", Item: Item{ID: 1, Name: "123", Priority: 123}}
	err = dbt.StructToSQLInsert(db, obj)
	if err != nil {
		panic(err)
	}
	obj2 := DirItem{Dir: "1234", Item: Item{ID: 1, Name: "1234", Priority: 0}}
	err = dbt.StructToSQLUpdate(db, obj2, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("create -table success")
	list, err := dbt.StructToSQLGetList(db, DirItem{})
	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
	// data, err := json.Marshal(obj)
	// if err != nil {
	// 	panic(err)
	// }
	// var result map[string]interface{}
	// err = json.Unmarshal(data, &result)
	// for k, v := range result {
	// 	t := reflect.TypeOf(v)
	// 	fmt.Printf("key: %-5s → type: %-20s → value: %v\n", k, t.String(), v)

	// }
	// if err != nil {
	// 	panic(err)
	// }
	// item := Item{ID: 1, Name: "123", Priority: 0}
	// m := StructToMap(item)
	// for item, value := range m {
	// 	fmt.Printf("%s: %v (%T)\n", item, value, value)
	// }
	// fmt.Println("----------------")
	// i := MapToStruct(m, reflect.TypeOf(DirItem{})).(DirItem)
	// fmt.Println(`test:>key`, i)
	// testFn(reflect.TypeOf(i))
}

func IsZero(v interface{}) bool {
	val := reflect.ValueOf(v)
	return val.IsZero() // Go 1.13+
}
