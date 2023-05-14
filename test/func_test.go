package testing

import (
	"github.com/esonhugh/sculptor"
	"sync"
	"testing"
)

type TestStruct struct {
	Name string `select:"name"`
	Pass string `select:"pass"`
}

func TestJsonParse(t *testing.T) {
	t.Log("Start")

	Doc := sculptor.NewDataSculptor("test.json").
		SetDocType(sculptor.JSON_DOCUMENT).
		SetQuery("name", "user").
		SetQuery("pass", "pass").
		SetTargetStruct(&TestStruct{})
	Doc.Do()
	go func() {
		Doc.Wg.Wait()
		close(Doc.ConstructedOutput)
	}()
	for i := range Doc.ConstructedOutput {
		t.Log(i)
	}
}

func TestCSVParse(t *testing.T) {
	t.Log("Start")
	wg := sync.WaitGroup{}
	Doc := sculptor.NewDataSculptorWithWg("test.csv", &wg).
		SetDocType(sculptor.CSV_DOCUMENT).
		SetQuery("name", "User").
		SetQuery("pass", "Pass").
		SetTargetStruct(&TestStruct{})
	Doc.Do()
	go func() {
		wg.Wait()
		close(Doc.ConstructedOutput)
	}()
	for i := range Doc.ConstructedOutput {
		t.Log(i)
	}
}

func TestSpaceCSV(t *testing.T) {
	t.Log("Start")

	Doc := sculptor.NewDataSculptor("blank_spilt.csv").
		SetDocType(sculptor.CSV_DOCUMENT).
		SetQuery("name", "User").
		SetQuery("pass", "Pass").
		SetTargetStruct(&TestStruct{}).SetCSVDelimiter(' ')
	Doc.Do()
	go func() {
		Doc.Wg.Wait()
		close(Doc.ConstructedOutput)
	}()
	for i := range Doc.ConstructedOutput {
		t.Log(i)
	}
}
