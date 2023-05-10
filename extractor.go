package GoDataExtractor

import (
	"context"
	"github.com/esonhugh/GoDataExtractor/parser"
	"github.com/esonhugh/GoDataExtractor/parser/query"
)

type DataExtractor[T any] struct {
	// DocType is the type of document to be extracted
	DocType DocumentType
	// Filename is processed filename or path
	Filename string

	// docQueries is the list of DocumentQuery to be executed
	docQueries []parser.DocumentQuery

	// scanner is the scanner used docQueries to extract data from the document
	scanner parser.RawDataParser

	// targetStruct is the sample struct to need to filled with the extracted data
	targetStruct T

	// CTX set for goruntime if thread process.
	CTX               context.Context
	count             uint64
	ConstructedOutput chan<- T

	// lastErr is the last error occurred while processing the record. If nil keep process.
	lastErr error

	// FallbackFunc process if record is bad.
	FallbackFunc func() T
}

// SetQuery sets the query for the given tag name
func (q *DataExtractor[T]) SetQuery(tagName string, Query string) *DataExtractor[T] {
	q.docQueries = append(q.docQueries, parser.DocumentQuery{
		Query:   Query,
		TagName: tagName,
	})
	return q
}

func (q *DataExtractor[T]) setDocType(docType DocumentType) *DataExtractor[T] {
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

func (q *DataExtractor[T]) SetScanner(dataParser parser.RawDataParser) *DataExtractor[T] {
	q.scanner = dataParser
	return q
}
