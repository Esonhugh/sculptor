package query

import (
	"github.com/tidwall/gjson"
	"testing"
)

func TestNewJsonReader(t *testing.T) {
	t.Log("Start")
	json := NewJsonReader("test.json")
	defer json.Close()
	if recordLine1, ok := json.Read(); ok == nil {
		t.Log(recordLine1)
	}
	if recordLine2, ok := json.Read(); ok == nil {
		t.Log(recordLine2)
	}
	if recordLine3, ok := json.Read(); ok == nil {
		t.Log(recordLine3)
	}
	t.Log("End")
}

func TestJSON_Select(t *testing.T) {
	json := `{
		"age":37,
			"children": ["Sara","Alex","Jack"],
	"fav.movie": "Deer Hunter",
	"friends": [
	{"age": 44, "first": "Dale", "last": "Murphy"},
	{"age": 68, "first": "Roger", "last": "Craig"},
	{"age": 47, "first": "Jane", "last": "Murphy"}
	],
	"name": {"first": "Tom", "last": "Anderson"}
}`
	value := gjson.Get(json, "name.AAA2121eawdslsd,askc.acds")

	t.Log(value)
}
