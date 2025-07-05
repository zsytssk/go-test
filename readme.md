https://go.dev/doc/tutorial/

## 2025-07-05 10:32:32

- @ques 如何使用 interface

- @ques

  - get order
  - get item
  - ***
  - get list
  - delete item
  - create table
  - insert item
  - update item

- @ques 如何组合 conditions

```go
limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db = db.Model(&process.Classify{}).Joins("left join sys_users on sys_users.id = p_process_classify.creator").
		Select("p_process_classify.*, sys_users.username as create_user, sys_users.nick_name as create_name") // 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&list).Error

// ----

var file example.ExaFileUploadAndDownload
	err := db.Where("id = ?", ID).First(&file).Error
	return file, err

```

- @ques 如何转换类型

```
	// UPDATE users
	// SET name = 'Alice', age = 30
	// WHERE id = 1;
```

- @ques
  - 数组怎么处理?
  - 递归类型怎么处理
  - interface {} 怎么处理

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
