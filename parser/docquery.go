package parser

// DocumentQuery is the struct that holds the query to be executed and other critical information
type DocumentQuery struct {
	// Query string like jquery csv_column_name
	Query string
	// If Query string should used as int
	Index int
	// TagName is the tag name of the struct field. This field will be set with the extracted data.
	TagName string

	// value is the value of the struct field.(according to the tag name)
	Value any
}

/*
func (q *DocumentQuery) SetValue(value reflect.Value) {
	q.value = value
}

func (q *DocumentQuery) GetValue() reflect.Value {
	return q.value
}
*/
