package GoDataExtractor

import (
	"context"
	"github.com/esonhugh/GoDataExtractor/parser"
	"github.com/esonhugh/GoDataExtractor/parser/query"
)

type Func func() error

type DataExtractor struct {
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
	fallbackFunc Func

	// customFunc Hooks before send to channel.
	customFunc Func
}

func NewDocumentExtractor(ctx context.Context, file string) *DataExtractor {
	return &DataExtractor{
		CTX:               ctx,
		Filename:          file,
		ConstructedOutput: make(chan any, 10),
	}
}

// SetQuery sets the query for the given tag name
func (q *DataExtractor) SetQuery(tagName string, Query string) *DataExtractor {
	q.docQueries = append(q.docQueries, parser.DocumentQuery{
		Query:   Query,
		TagName: tagName,
	})
	return q
}

func (q *DataExtractor) SetDocType(docType DocumentType) *DataExtractor {
	q.DocType = docType
	var dataParser parser.RawDataParser
	switch docType {
	case CSV_DOCUMENT:
		dataParser = query.NewCsvReader(q.Filename)
	case JSON_DOCUMENT:
		dataParser = query.NewJsonReader(q.Filename)
	default:
		panic("Document Type Not Supported")
	}
	q.SetScanner(dataParser)
	return q
}

func (q *DataExtractor) SetScanner(dataParser parser.RawDataParser) *DataExtractor {
	q.scanner = dataParser
	return q
}

func (q *DataExtractor) SetCustomFunc(f Func) *DataExtractor {
	q.customFunc = f
	return q
}

func (q *DataExtractor) SetFallbackFunc(f Func) *DataExtractor {
	q.fallbackFunc = f
	return q
}
