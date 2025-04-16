package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 模拟 ShouldBindWith 函数
func ShouldBindWith(obj any, _ interface{}) error {
	jsonData := []byte(`{"name": "John", "age": 30}`)

	// 使用反射获取 obj 的类型和值
	value := reflect.ValueOf(obj)
	if value.Kind() != reflect.Ptr || value.IsNil() {
		return fmt.Errorf("obj must be a non-nil pointer")
	}
	// 获取指针指向的值
	elem := value.Elem()

	// 使用 json.Unmarshal 解析 JSON 数据
	err := json.Unmarshal(jsonData, elem.Addr().Interface())
	if err != nil {
		return err
	}
	return nil
}

// Person 结构体
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := &Person{}
	err := ShouldBindWith(p, nil)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
	}
}
