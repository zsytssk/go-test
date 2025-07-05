package main

import (
	"database/sql"
	"fmt"
	"go-test/dbt"
	"go-test/utils"
	"reflect"
	"strings"
	"time"

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

func formatSQLValue(val interface{}) string {
	switch v := val.(type) {
	case string:
		return fmt.Sprintf("'%s'", strings.ReplaceAll(v, "'", "''")) // 防止 SQL 注入
	case time.Time:
		return fmt.Sprintf("'%s'", v.Format("2006-01-02 15:04:05"))
	case nil:
		return "NULL"
	default:
		return fmt.Sprintf("%v", v) // int, float, bool 等直接返回
	}
}

func StructToSQLCreateTable(db *sql.DB, obj DirItem) (err error) {
	if dbt.CheckTableExist(db, obj.TableName()) {
		return
	}
	fields_list := collectFields(obj)
	var columns []string
	for _, field := range fields_list {
		db_type := field["dbType"].(string)
		if db_type == "primaryKey" {
			columns = append(columns, fmt.Sprintf("  %s %s PRIMARY KEY", field["name"], field["sqlType"]))
			continue
		}
		columns = append(columns, fmt.Sprintf("  %s %s", field["name"], field["sqlType"]))
	}

	sqlStr := fmt.Sprintf("CREATE TABLE %s (\n%s\n);",
		strings.ToLower(obj.TableName()),
		strings.Join(columns, ",\n"),
	)
	fmt.Println(sqlStr)
	_, err = db.Exec(sqlStr)
	if err != nil {
		return
	}
	return
}
func StructToSQLInsert(db *sql.DB, obj DirItem) (err error) {
	exists, err := CheckItemExist(db, obj)
	if exists || err != nil {
		return
	}
	fields_list := collectFields(obj)
	var columns []string
	var placeholders []string
	var values []interface{}
	for _, field := range fields_list {
		columns = append(columns, field["name"].(string))
		placeholders = append(placeholders, "?")
		values = append(values, formatSQLValue(field["value"]))
	}
	sqlStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
		strings.ToLower(obj.TableName()),
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	_, err = db.Exec(sqlStr, values...)
	if err != nil {
		return
	}
	return
}
func StructToSQLUpdate(db *sql.DB, obj DirItem, ignore_zero bool) (err error) {
	fields_list := collectFields(obj)
	var columns []string
	var where_str string
	for _, field := range fields_list {
		db_type := field["dbType"].(string)
		if db_type == "primaryKey" {
			where_str = fmt.Sprintf("  %s = %v", field["name"], formatSQLValue(field["value"]))
			continue
		}
		if ignore_zero && utils.IsZero(field["value"]) {
			continue
		}
		columns = append(columns, fmt.Sprintf("  %s = %v", field["name"], formatSQLValue(field["value"])))
	}
	sqlStr := fmt.Sprintf("UPDATE %s\nSET %s \n WHERE %s;",
		strings.ToLower(obj.TableName()),
		strings.Join(columns, ", "),
		where_str,
	)
	_, err = db.Exec(sqlStr)
	if err != nil {
		return
	}
	return
}
func StructToSQLSave(obj DirItem) string {
	return ""
}
func StructToSQLDelete(db *sql.DB, obj DirItem) (err error) {
	// DELETE FROM users WHERE id = 3;
	fields_list := collectFields(obj)
	var where_str string
	for _, field := range fields_list {
		db_type := field["dbType"].(string)
		if db_type == "primaryKey" {
			where_str = fmt.Sprintf("  %s = %v", field["name"], formatSQLValue(field["value"]))
			break
		}
	}
	sqlStr := fmt.Sprintf("DELETE FROM %s WHERE %s;",
		strings.ToLower(obj.TableName()),
		where_str,
	)
	_, err = db.Exec(sqlStr)
	if err != nil {
		return
	}
	return
}
func StructToSQLDeleteTable(db *sql.DB, obj DirItem) (err error) {
	// DROP TABLE IF EXISTS table_name;
	sqlStr := fmt.Sprintf("DROP TABLE IF EXISTS %s;",
		strings.ToLower(obj.TableName()),
	)
	_, err = db.Exec(sqlStr)
	if err != nil {
		return
	}
	return
}
func CheckItemExist(db *sql.DB, obj DirItem) (exists bool, err error) {
	fields_list := collectFields(obj)
	var where_str string
	for _, field := range fields_list {
		db_type := field["dbType"].(string)
		if db_type == "primaryKey" {
			where_str = fmt.Sprintf("  %s = %v", field["name"], formatSQLValue(field["value"]))
			break
		}
	}

	sqlStr := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s);",
		strings.ToLower(obj.TableName()),
		where_str,
	)

	err = db.QueryRow(sqlStr).Scan(&exists)
	return
}
func collectFields(obj interface{}) (connects []map[string]interface{}) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)
		fieldVal := v.Field(i)

		// 跳过未导出字段
		if !fieldType.IsExported() || !fieldVal.CanInterface() {
			continue
		}
		value := fieldVal.Interface()
		// 匿名字段 && 是 struct，递归处理
		if fieldType.Anonymous && fieldType.Type.Kind() == reflect.Struct {
			sub_list := collectFields(value)
			connects = append(connects, sub_list...)
			continue
		}

		// 字段名使用 json tag，否则用字段名
		name := fieldType.Tag.Get("json")
		if name == "" || name == "-" {
			name = utils.CamelToSnake(fieldType.Name)
		}
		dbType := fieldType.Tag.Get("db")

		sqlType := GoTypeToSQLType(fieldType.Type)
		connects = append(connects, map[string]interface{}{
			"dbType":  dbType,
			"name":    name,
			"sqlType": sqlType,
			"value":   value,
		})

	}

	return
}

func main() {
	db, err := dbt.InitDB("./test.db")
	if err != nil {
		panic(err)
	}
	err = StructToSQLCreateTable(db, DirItem{})
	if err != nil {
		panic(err)
	}

	obj := DirItem{Dir: "123", Item: Item{ID: 1, Name: "123", Priority: 123}}
	err = StructToSQLInsert(db, obj)
	if err != nil {
		panic(err)
	}
	obj2 := DirItem{Dir: "1234", Item: Item{ID: 1, Name: "1234", Priority: 0}}
	err = StructToSQLUpdate(db, obj2, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("create -table success")
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
