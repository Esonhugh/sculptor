package query

import (
	"bufio"
	"encoding/csv"
	"github.com/esonhugh/GoDataExtractor/parser"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"reflect"
	"unsafe"
)

type CSV struct {
	Header   []string
	Recorder *csv.Reader
	file     io.Closer
}

func NewCsvReader(filename string) *CSV {
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Error(err)
		return nil
	}
	log.Info("Successfully Opened CSV file")

	reader := csv.NewReader(csvFile)
	// reader.ReuseRecord = true
	Header, err := reader.Read() // give CSV file header line
	if err != nil {
		log.Error("GET csv file header error", err)
	}

	log.Debug("Header: ", Header)

	return &CSV{
		Header:   Header,
		Recorder: reader,
		file:     csvFile,
	}
}

// SetDelimiter is func to set CSV.Recorder.Comma
func (c *CSV) SetDelimiter(r rune) {
	c.Recorder.Comma = r
}

// InitIndex func will set the Index of Column in Selector
func (c *CSV) InitIndex(Selectors []parser.DocumentQuery) {
	for i, header := range c.Header {
		for _, v := range Selectors {
			if header == v.Query {
				Selectors[i].Index = i
			}
		}
	}
}

// Read func will read one line for csv record.
func (c *CSV) Read() ([]string, error) {
	return c.Recorder.Read()
}

func (c *CSV) Close() {
	c.file.Close()
}

func (c *CSV) Select(s []parser.DocumentQuery) error {
	if OneRecord, e := c.Read(); e != nil {
		c.CurrentLineNum()
		log.Error("Read CSV file error", e)
		return e
	} else {
		for i, v := range s {
			s[i].SetValue(reflect.ValueOf(OneRecord[v.Index]))
		}
		return nil
	}
}

// Hack for Got CSV File using reflect get Current Error Line
// Reader csv defined.
type Reader struct {
	// Comma is the field delimiter.
	// It is set to comma (',') by NewReader.
	// Comma must be a valid rune and must not be \r, \n,
	// or the Unicode replacement character (0xFFFD).
	Comma rune

	// Comment, if not 0, is the comment character. Lines beginning with the
	// Comment character without preceding whitespace are ignored.
	// With leading whitespace the Comment character becomes part of the
	// field, even if TrimLeadingSpace is true.
	// Comment must be a valid rune and must not be \r, \n,
	// or the Unicode replacement character (0xFFFD).
	// It must also not be equal to Comma.
	Comment rune

	// FieldsPerRecord is the number of expected fields per record.
	// If FieldsPerRecord is positive, Read requires each record to
	// have the given number of fields. If FieldsPerRecord is 0, Read sets it to
	// the number of fields in the first record, so that future records must
	// have the same field count. If FieldsPerRecord is negative, no check is
	// made and records may have a variable number of fields.
	FieldsPerRecord int

	// If LazyQuotes is true, a quote may appear in an unquoted field and a
	// non-doubled quote may appear in a quoted field.
	LazyQuotes bool

	// If TrimLeadingSpace is true, leading white space in a field is ignored.
	// This is done even if the field delimiter, Comma, is white space.
	TrimLeadingSpace bool

	// ReuseRecord controls whether calls to Read may return a slice sharing
	// the backing array of the previous call's returned slice for performance.
	// By default, each call to Read returns newly allocated memory owned by the caller.
	ReuseRecord bool

	TrailingComma bool // Deprecated: No longer used.

	r *bufio.Reader

	// numLine is the current line being read in the CSV file.
	numLine int

	// rawBuffer is a line buffer only used by the readLine method.
	RawBuffer []byte

	// recordBuffer holds the unescaped fields, one after another.
	// The fields can be accessed by using the indexes in fieldIndexes.
	// E.g., For the row `a,"b","c""d",e`, recordBuffer will contain `abc"de`
	// and fieldIndexes will contain the indexes [1, 2, 5, 6].
	recordBuffer []byte

	// fieldIndexes is an index of fields inside recordBuffer.
	// The i'th field ends at offset fieldIndexes[i] in recordBuffer.
	fieldIndexes []int

	// fieldPositions is an index of field positions for the
	// last record returned by Read.
	fieldPositions []position

	// lastRecord is a record cache and only used when ReuseRecord == true.
	lastRecord []string
}

// Scanner Location
type position struct {
	line, col int
}

func (c *CSV) CurrentLine() string {
	return string((*(*Reader)(unsafe.Pointer(c.Recorder))).RawBuffer)
	// Same implement as below..... But with reflect.
	// get := reflect.ValueOf(*c.Recorder)
	// return string(get.FieldByName("rawBuffer").Bytes())
}

// CurrentLineNum is func to get current line number(include the header line) Sample. Save This and got this line.
func (c *CSV) CurrentLineNum() int {
	Class := reflect.ValueOf(*c.Recorder)
	return int(Class.FieldByName("numLine").Int())
}
