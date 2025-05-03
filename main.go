package main

import (
	"fmt"
	"sort"
)

// Person 定义一个结构体
type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 20},
		{"Charlie", 30},
	}

	// 按年龄升序排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})

	fmt.Println(people)
}
