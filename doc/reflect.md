### reflect

go reflect 是一种运行时处理类型的方式, 在一些处理多种类型的函数是, 需要这个来实现.
我感觉他几乎能实现任何功能, 从创建到设置值, 到一个个的判断都可以. 其中最基本的两个类型就是 `reflect.Type` 和 `reflect.Value`.

很多语言通过泛型来实现类似的功能, 他们付出的代价是编译消耗.

- reflect.Type

  - `reflect.TypeOf(i)` `v.Type()` `t.Field(i)`
  - `.NumField()`

- reflect.Value
  - `v = v.Elem()` `reflect.New(t).Elem()` `reflect.ValueOf(v)`
  - `v.Type()` -> reflect.Type
  - struct `v.NumField()` `v.Field(i)` `v.FieldByName(i)`
  - slice `v.Len()` `v.Index(i)` `reflect.Append(v, ...)`
  - `.IsValid()`
  - `.CanInterface()` `v.Interface()`
  - `.CanAddr()` `.Addr().Interface()`
  - `.CanSet()` `.Set(reflect.ValueOf(value))`
  - `v.IsNil()`

```go
// 这个Find函数展现reflect强大功能, 他将数据的数据查找到 然后放到一个interface{}中
// 这个函数可以用来做几乎任意的struct和对应的table, 只要他们能一一对应.
Find(dest interface{}) *Model {
	destVal := reflect.ValueOf(dest)
	if destVal.Kind() == reflect.Ptr {
		destVal = destVal.Elem() // 解引用 *slice => slice
	}

	sql := fmt.Sprintf("SELECT * FROM %s", m.BuildConditions(m.conditions))
	rows, err := m.DB.Query(sql)
	if err != nil {
		m.Error = err
		return m
	}
	defer rows.Close()
	for rows.Next() {
		typ := reflect.TypeOf(m.Obj)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem() // 获取指针指向的真实类型
		}
		elemPtr := reflect.New(typ) // *T
		elemVal := elemPtr.Elem()   // T
		var scanArgs []interface{}

		for _, field := range m.TableInfo.FieldsList {
			fieldName := field["oriName"].(string)

			// 获取结构体中对应的字段
			structField := elemVal.FieldByName(fieldName) // 需要转换

			// 跳过非法或未导出字段
			if !structField.IsValid() || !structField.CanAddr() {
				continue
			}

			// 添加字段地址作为 Scan 参数
			scanArgs = append(scanArgs, structField.Addr().Interface())
		}
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}
		destVal = reflect.Append(destVal, elemVal)
	}
	reflect.ValueOf(dest).Elem().Set(destVal)

	return m
}
```
