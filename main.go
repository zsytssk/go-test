package main

import (
	"encoding/json"
	"fmt"
	"go-test/dbt"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	ID       int    `db:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

type DirItem struct {
	Item
	Hide bool `json:"hidden"`
}

func (DirItem) TableName() string {
	return "dir"
}

func main() {
	db, err := dbt.InitDB("./test.db")
	if err != nil {
		panic(err)
	}
	db1 := dbt.NewModel(db, &DirItem{})

	var count int64
	err = db1.Count(&count).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(`test:>count`, count)
	if count == 0 {
		return
	}
	var item DirItem
	err = db1.First(&item).Error
	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(`test:>item`, string(jsonBytes))
	item.ID = 33
	item.Priority = 20
	item.Hide = true
	err = db1.Delete(&item).Error
	if err != nil {
		panic(err)
	}
	var list []DirItem
	err = db1.Order("ORDER BY priority DESC").Limit(2).Offset(0).Find(&list).Error
	if err != nil {
		panic(err)
	}

	fmt.Println(`test:>items`, len(list[0].Name))
	jsonBytes1, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(`test:>items`, string(jsonBytes1))
}
