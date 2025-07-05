package dbt

import "database/sql"

type Model struct {
	DB    *sql.DB
	Obj   *TableStruct
	Error error
}

func (m *Model) Where(query string) *Model {
	return m
}
func (m *Model) Count(count *int64) *Model {
	return m
}
func (m *Model) Limit(limit *int64) *Model {
	return m
}
func (m *Model) Offset(offset *int64) *Model {
	return m
}
func (m *Model) Find(dest interface{}) *Model {
	return m
}
func (m *Model) Order(order string) *Model {
	return m
}
func (m *Model) First(first interface{}) *Model {
	return m
}
func (m *Model) Last(first interface{}) *Model {
	return m
}
