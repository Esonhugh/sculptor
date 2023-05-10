package parser

// DocumentQuery is the struct that holds the query to be executed and the value of query result.
type DocumentQuery struct {
	// Query string like jquery csv_column_name
	Query string

	// If Query string should be used as int to index the record. when Split is or other things.
	Index int

	// TagName is the tag name of the struct field which will be set with the extracted data.
	TagName string

	// Value is the value of the struct field. And This attribute will hold the query result.
	Value any
}
