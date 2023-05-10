package parser

type RawDataParser interface {
	// Close func will close the file io and destruct the data flow Scanners
	Close()
	// InitIndex Set the selectors index
	// Not Required But
	InitIndex(Selectors []DocumentQuery)
	// Select func will select the data from the record.
	// Read RawDataParser one-Lined Record
	// and using query.Selector
	// to select the values from the record.
	// Value will in query.Selector[i].value
	// using query.Selector[i].GetValue to get the value.
	// if EOF or Other error, will got error != nil.
	Select(selectors []DocumentQuery) error

	// CurrentLine func will return the current decoding record.
	CurrentLine() string
}
