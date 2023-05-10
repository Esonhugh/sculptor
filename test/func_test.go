package testing

import (
	gdocExt "github.com/esonhugh/GoDataExtractor"
	"reflect"
	"testing"
)

type TestStruct struct {
	Name string `extract:"name"`
	Pass string `extract:"pass"`
}

func TestReflectReplace(t *testing.T) {
	t.Log("Start")
	testStruct := &TestStruct{
		Name: "1314",
		Pass: "520",
	}
	V := reflect.Indirect(reflect.ValueOf(testStruct))
	for i := 0; i < V.NumField(); i++ {
		tagStr := V.Type().Field(i).Tag.Get("extract")
		t.Log("tagStr:", tagStr)
		if tagStr == "name" {
			t.Log("name:", V)
			V.Field(i).Set(reflect.ValueOf("114514"))
			t.Log("name:", V.String())
		}
	}
	t.Log(testStruct)
}

func TestJsonParse(t *testing.T) {
	t.Log("Start")

	Doc := gdocExt.NewDocExtractor("test.json").
		SetDocType(gdocExt.JSON_DOCUMENT).
		SetQuery("name", "user").
		SetQuery("pass", "pass").
		SetTargetStruct(&TestStruct{})
	Doc.Do()
	for i := range Doc.ConstructedOutput {
		t.Log(i)
	}
}
