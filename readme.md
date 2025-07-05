https://go.dev/doc/tutorial/

## 2025-07-05 10:32:32

- @ques

  - delete item
  - ***
  - create table
  - insert item
  - update item

- @ques 如何使用 interface

- @ques 如何转换类型

```
	// UPDATE users
	// SET name = 'Alice', age = 30
	// WHERE id = 1;
```

- @ques 数组怎么处理?

## 2025-06-01 10:42:27

- @ques reflect 的使用

### reflect

- @ques

  - create table
  - insert item
  - update item
  - delete item

- reflect.Type

  - `reflect.TypeOf(i)` `v.Type()`
  - `.NumField()`
  - `t.Field(i)`

- reflect.Value
  - `v = v.Elem()` `reflect.New(t).Elem()` `reflect.ValueOf(v)`
  - `v.Type()` -> reflect.Type
  - `v.NumField()` `v.Field(i)`
  - `.CanInterface()` `fieldVal.Interface()`
  - `.IsValid()` `.CanSet()`
  - `field.Set(reflect.ValueOf(value))`
  - `v.FieldByName(key)`
