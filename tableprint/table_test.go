package tableprint

import "testing"

type TestStructLabels struct {
	Data    string `select:"data"`
	KeyInfo string `select:"key"`
	A       int    `select:"num"`
}

func TestTable(t *testing.T) {
	t.Log("Start")
	PrintStruct("select", []TestStructLabels{
		{
			Data:    "Helloworld",
			KeyInfo: "Key",
			A:       12,
		},
		{
			Data:    "Bad",
			KeyInfo: "value",
			A:       11,
		},
		{
			Data:    "Nothing",
			KeyInfo: "Hello",
			A:       0,
		},
	})
}
