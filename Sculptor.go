package sculptor

import (
	"context"
	"errors"
	"github.com/esonhugh/sculptor/parser"
	"github.com/esonhugh/sculptor/parser/query"
	"time"
)

// Func is alias of DataSculptor customFunc and fallbackFunc type
type Func func(*DataSculptor) error

// DocumentType is core struct in package.
// it contains all the information required to extract data from a document during runtime.
type DataSculptor struct {

	// DocType is the type of document to be extracted
	DocType DocumentType

	// Filename is processed filename or path
	Filename string

	// docQueries is the list of DocumentQuery to be executed
	docQueries []parser.DocumentQuery

	// scanner is the scanner used docQueries to extract data from the document
	scanner parser.RawDataParser

	// targetStruct is the sample struct to need to filled with the extracted data
	targetStruct any

	// CTX set for goruntime if thread process.
	CTX context.Context

	// count is the number of records processed
	count uint64

	// ConstructedOutput is the channel to send the extracted data out
	ConstructedOutput chan any

	// lastErr is the last error occurred while processing the record. If nil keep process.
	lastErr error

	// fallbackFunc process if record is bad.
	fallbackFunc []Func

	// customFunc Hooks before send to channel.
	customFunc []Func

	// options make CustomOptions
	options Options
}

// NewDataSculptor returns a new DataSculptor with initialized values
func NewDataSculptor(file string) *DataSculptor {
	return &DataSculptor{
		CTX:               context.Background(),
		Filename:          file,
		ConstructedOutput: make(chan any, 10),
		options: Options{
			TagKey:  "select",
			Latency: 1 * time.Millisecond,
		},
	}
}

// SetQuery sets the query for the given tag name
func (d *DataSculptor) SetQuery(tagName string, Query string) *DataSculptor {
	d.docQueries = append(d.docQueries, parser.DocumentQuery{
		Query:   Query,
		TagName: tagName,
	})
	return d
}

// SetDocType sets the document type for the given filename.
// If Supported, It will automatically set the scanner for the given document type.
// If not supported, it will panic.
// If you want Set you own scanners please use SetScanner() and follow the interface.
func (d *DataSculptor) SetDocType(docType DocumentType) *DataSculptor {
	d.DocType = docType
	var dataParser parser.RawDataParser
	switch docType {
	case CSV_DOCUMENT:
		dataParser = query.NewCsvReader(d.Filename)
	case JSON_DOCUMENT:
		dataParser = query.NewJsonReader(d.Filename)
	default:
		panic("Document Type Not Supported")
	}
	d.SetScanner(dataParser)
	return d
}

// SetCSVDelimiter sets the delimiter for the CSV document if you set the document type to CSV.
// Else it will make error.
func (d *DataSculptor) SetCSVDelimiter(r rune) *DataSculptor {
	if d.DocType != CSV_DOCUMENT {
		d.lastErr = errors.New("your Document Type is not CSV. Please check your document type")
	}
	d.scanner.(*query.CSV).SetDelimiter(r)
	return d
}

// SetScanner sets the scanner (parser.RawDataParser) which used to extract data from the document.
// SetScanner is helpful if you want to use your own scanner to process your file.
func (d *DataSculptor) SetScanner(dataParser parser.RawDataParser) *DataSculptor {
	d.scanner = dataParser
	return d
}

// SetCustomFunc sets the customFunc which will be called
// between constructing targetStruct complete and sending the extracted data to the channel.
func (d *DataSculptor) SetCustomFunc(f ...Func) *DataSculptor {
	d.customFunc = append(d.customFunc, f...)
	return d
}

// SetFallbackFunc sets the fallbackFunc which will be called when framework can't handle the record.
func (d *DataSculptor) SetFallbackFunc(f ...Func) *DataSculptor {
	d.fallbackFunc = append(d.fallbackFunc, f...)
	return d
}

// SetTargetStruct sets the target struct with the given struct pointer.
// Helpful in SetFallbackFunc and SetCustomFunc.
// It will be called when init before the Do() func.
func (d *DataSculptor) SetTargetStruct(targetStruct any) *DataSculptor {
	d.targetStruct = targetStruct
	return d
}

// CurrentTarget func returns the current target struct during process.
// Helpful in SetFallbackFunc and SetCustomFunc.
func (d *DataSculptor) CurrentTarget() any {
	return d.targetStruct
}

// Error() func returns the last error occurred while processing the record.
// Helpful in SetFallbackFunc and SetCustomFunc.
func (d *DataSculptor) Error() error {
	return d.lastErr
}

// Send func sends the extracted data to the channel.
func (d *DataSculptor) send() {
	d.ConstructedOutput <- d.targetStruct
	time.Sleep(d.options.Latency)
}
