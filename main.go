package main

import (
	"fmt"
	"go-test/utils"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	ID       int    `json:"id"`
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

func GoTypeToSQLType(goType reflect.Type) string {
	switch goType.Kind() {
	case reflect.Int, reflect.Int32:
		return "INT"
	case reflect.Int64:
		return "BIGINT"
	case reflect.Uint, reflect.Uint64:
		return "BIGINT UNSIGNED"
	case reflect.Float32:
		return "FLOAT"
	case reflect.Float64:
		return "DOUBLE"
	case reflect.Bool:
		return "TINYINT(1)"
	case reflect.String:
		return "VARCHAR(255)"
	case reflect.Slice:
		if goType.Elem().Kind() == reflect.Uint8 {
			return "BLOB"
		}
	case reflect.Struct:
		if goType.PkgPath() == "time" && goType.Name() == "Time" {
			return "DATETIME"
		}
	}
	return "TEXT"
}

func StructToSQLDefinition(t reflect.Type) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	fmt.Printf("CREATE TABLE %s (\n", strings.ToLower(t.Name()))

	collectFields(t, "")
	fmt.Println(");")
}

func collectFields(t reflect.Type, indent string) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// 跳过未导出字段
		if !field.IsExported() {
			continue
		}

		// 匿名字段 && 是 struct，递归处理
		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			collectFields(field.Type, indent)
			continue
		}

		// 字段名使用 json tag，否则用字段名
		name := field.Tag.Get("json")
		if name == "" || name == "-" {
			name = utils.CamelToSnake(field.Name)
		}

		sqlType := GoTypeToSQLType(field.Type)
		fmt.Printf("%s  %s %s,\n", indent, name, sqlType)
	}
}

func main() {
	// db, err := dbt.InitDB("test.db")
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = dbt.CreateTable(db, DirItem{})
	// if err != nil {
	// 	panic(err)
	// }

	obj := DirItem{Dir: "123", Item: Item{ID: 1, Name: "123", Priority: 0}}
	StructToSQLDefinition(reflect.TypeOf(obj))
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
