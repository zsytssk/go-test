package dbt

import (
	"database/sql"
	"fmt"
	"go-test/utils"
	"reflect"
	"strings"
)

type TableInterface interface {
	TableName() string
}

func InitDB(dbPath string) (db *sql.DB, err error) {
	filePath, err := utils.GetCurDirFilePath(dbPath)
	if err != nil {
		return
	}
	fmt.Println(filePath)
	db, err = sql.Open("sqlite3", filePath)
	if err != nil {
		return
	}
	return
}

func UpdateOrInsert(db *sql.DB, item TableInterface) {
	tableName := item.TableName()
	if !CheckTableExist(db, tableName) {

	}
}

func CreateTable(db *sql.DB, item TableInterface) (sus bool, err error) {

	var tableStmt strings.Builder
	tableName := item.TableName()
	t := reflect.TypeOf(item)
	tableStmt.WriteString(fmt.Sprintf("CREATE TABLE %s\n", tableName))
	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)
		fieldName := fieldType.Name
		keyName := utils.CamelToSnake(fieldName)
		if keyName == "id" {
			tableStmt.WriteString("id INTEGER PRIMARY KEY,\n")
		} else {
			typeStr := mapTypeToDbType(fieldType.Type)
			if i == t.NumField()-1 {
				tableStmt.WriteString(fmt.Sprintf("%s %s\n);", keyName, typeStr))
			} else {
				tableStmt.WriteString(fmt.Sprintf("%s %s,\n", keyName, typeStr))
			}
		}
	}
	fmt.Println(tableStmt.String())
	_, err = db.Exec(tableStmt.String())
	if err != nil {
		return false, err
	}
	return true, nil
}

func mapTypeToDbType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Int:
		return "INTEGER"
	case reflect.String:
		return "TEXT"
	default:
		return "NULL"
	}
}

func CheckTableExist(db *sql.DB, tableName string) (exist bool) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM sqlite_master WHERE type='table' AND name=?)"
	err := db.QueryRow(query, tableName).Scan(&exists)
	return err == nil && exists
}
