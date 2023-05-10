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
	"reflect"
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
		if j.Recorder.Err() == io.EOF {
			return "", io.EOF
		}
		return j.Recorder.Text(), nil
	} else {
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
			return errors.New(fmt.Sprintf("Value Not Found! \n"+
				"- json: %v\n"+
				"- json query: %v\n",
				j.CurrLine,
				selects.Query))
		}
		selector[i].SetValue(reflect.ValueOf(value.Value()))
	}
	return
}

func (j *JSON) CurrentLine() string {
	return j.CurrLine
}
