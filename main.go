package main

import (
	"fmt"
	"reflect"
)

type Item struct {
	ID       int
	Name     string
	Priority int
}

type DirItem struct {
	Item
	Dir string
}

func (DirItem) TableName() string {
	return "device_shutdown_operation"
}

func main() {
	item := DirItem{Dir: "123", Item: Item{ID: 1, Name: "123", Priority: 0}}

	m := StructToMap(item)
	for item, value := range m {
		fmt.Printf("%s: %v (%T)\n", item, value, value)
	}
	fmt.Println("----------------")
	i := MapToStruct(m, reflect.TypeOf(DirItem{})).(DirItem)
	fmt.Println(`test:>key`, i)
	testFn(reflect.TypeOf(i))
}

func testFn(t reflect.Type) {
	v := reflect.New(t)
	fmt.Println(`test:>key`, v.Kind())
}

func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(obj)

	// 如果是指针，先取指针指向的值
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldType := t.Field(i)
		fieldVal := v.Field(i)
		if !fieldVal.CanInterface() {
			continue
		}
		value := fieldVal.Interface()
		if IsZero(value) {
			continue
		}
		if fieldType.Type.Kind() == reflect.Struct {
			result[fieldType.Name] = StructToMap(value)
			continue
		}
		result[fieldType.Name] = value
	}

	return result
}

func MapToStruct(obj map[string]interface{}, t reflect.Type) interface{} {
	// 创建结构体实例
	v := reflect.New(t).Elem()

	for key, value := range obj {
		// 获取结构体字段
		field := v.FieldByName(key)
		if !field.IsValid() || !field.CanSet() {
			continue
		}
		if field.Type().Kind() == reflect.Struct {
			// 递归处理嵌套结构体
			nestedMap, ok := value.(map[string]interface{})
			if !ok {
				continue
			}
			nestedStruct := MapToStruct(nestedMap, field.Type())

			field.Set(reflect.ValueOf(nestedStruct))
			continue
		}
		// 设置字段值
		field.Set(reflect.ValueOf(value))
	}

	return v.Interface()
}

func IsZero(v interface{}) bool {
	val := reflect.ValueOf(v)
	return val.IsZero() // Go 1.13+
}
