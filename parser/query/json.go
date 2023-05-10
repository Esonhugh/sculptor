package query

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/esonhugh/GoDataExtractor/parser"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
	"os"
)

type JSON struct {
	file     io.Closer
	Recorder *bufio.Scanner
	CurrLine string
}

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

// InitIndex Set the selectors index
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

func (j *JSON) Select(selector []parser.DocumentQuery) (err error) {
	j.CurrLine, err = j.Read()
	if err != nil {
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

func (j *JSON) CurrentLine() string {
	return j.CurrLine
}
