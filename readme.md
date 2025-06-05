https://go.dev/doc/tutorial/

## 2025-06-01 10:42:27

- @ques reflect 的使用
- @ques 数组怎么处理?

### reflect

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
