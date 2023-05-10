package parser

import "reflect"

// DocumentQuery is the struct that holds the query to be executed and other critical information
type DocumentQuery struct {
	// Query string like jquery csv_column_name
	Query string
	// If Query string should used as int
	Index int
	// TagName is the tag name of the struct field. This field will be set with the extracted data.
	TagName string

	// valueType is the type of the struct field.(according to the tag name)
	valueType reflect.Type
	// value is the value of the struct field.(according to the tag name)
	value reflect.Value
}

func (q *DocumentQuery) SetValue(value reflect.Value) {
	q.value.Set(value)
}

func (q *DocumentQuery) GetValue() reflect.Value {
	return q.value
}
