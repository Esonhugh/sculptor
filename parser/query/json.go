package query

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/esonhugh/sculptor/parser"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
	"os"
)

// JSON is real data parser for JSON file. It should implement parser.RawDataParser interface
type JSON struct {
	file     io.Closer
	Recorder *bufio.Scanner
	CurrLine string
}

// NewJsonReader is func to create a JSON data parser
func NewJsonReader(filename string) *JSON {
	jsonfile, err := os.Open(filename)
	if err != nil {
		log.Error(err)
		return nil
	}
	log.Info("Successful Opened JSON file")
	fileScanner := bufio.NewScanner(jsonfile)
	fileScanner.Split(bufio.ScanLines)
	return &JSON{
		file:     jsonfile,
		Recorder: fileScanner,
	}
}

// InitIndex Set the selectors index. Useless for json there.
func (j *JSON) InitIndex(selectors []parser.DocumentQuery) {

}

// Read func will read one line for json record
func (j *JSON) Read() (string, error) {
	if j.Recorder.Scan() {
		return j.Recorder.Text(), nil
	} else {
		if j.Recorder.Err() == nil {
			return "", io.EOF
		}
		return "", errors.New(fmt.Sprintf("EOF or Other Error: %v", j.Recorder.Err()))
	}
}

// Close func will close the file
func (j *JSON) Close() {
	j.file.Close()
}

// Select func will select the data from the one line in file.
func (j *JSON) Select(selector []parser.DocumentQuery) (err error) {
	j.CurrLine, err = j.Read()
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Error(err)
		return
	}
	for i, selects := range selector {
		value := gjson.Get(j.CurrLine, selects.Query)
		if !value.Exists() {
			return errors.New(fmt.Sprintf("Value Not Found! "+
				"json: >>%v<<"+
				"json query: >>%v<<",
				j.CurrLine,
				selects.Query))
		}
		selector[i].Value = value.Value()
	}
	return
}

// CurrentLine func will return the current decoding record.
func (j *JSON) CurrentLine() string {
	return j.CurrLine
}
