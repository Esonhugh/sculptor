package tableprint

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"reflect"
)

// TableData is a struct for Table Data which can be print pretty.
type TableData struct {
	Header []string
	Body   [][]string
}

// PrintTable print Tables for TableData Struct.
func PrintTable(data TableData, Caption string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(data.Header)
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	var TableHeaderColor = make([]tablewriter.Colors, len(data.Header))
	for i := range TableHeaderColor {
		TableHeaderColor[i] = tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor}
	}
	table.SetHeaderColor(TableHeaderColor...)
	if Caption != "" {
		table.SetCaption(true, Caption)
	}
	table.AppendBulk(data.Body)
	table.Render()
}

// PrintStruct use tagName as header and value as []body to print Struct
func PrintStruct[T any](tagName string, sample []T) {

	if len(sample) == 0 && tagName != "" {
		return
	}

	Type := reflect.TypeOf(sample[0])
	for _, v := range sample {
		if Type != reflect.TypeOf(v) {
			return
		}
	}

	td := TableData{}
	var fieldId []int
	for i := 0; i < Type.NumField(); i++ {
		if v, ok := Type.Field(i).Tag.Lookup(tagName); ok {
			td.Header = append(td.Header, v) // tagName Value1
			fieldId = append(fieldId, i)
		}
	}

	for _, v := range sample {
		var currentRecord []string
		for i := range fieldId {
			fieldValue := reflect.ValueOf(v).Field(i)
			currentRecord = append(currentRecord, fmt.Sprintf("%v", fieldValue.Interface()))
		}
		td.Body = append(td.Body, currentRecord)
	}
	PrintTable(td, "")
}
