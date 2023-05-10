package parser

import (
	"github.com/olekukonko/tablewriter"
	"os"
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
