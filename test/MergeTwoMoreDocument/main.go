package main

import (
	"fmt"
	"github.com/esonhugh/sculptor"
)

type SubDomains struct {
	Name   string `select:"domain"`
	Source string `select:"source"`
}

// Demo of merge subfinder and oneforall
func main() {
	doc_subfinder := sculptor.NewDataSculptor("subfinder_output.json").
		SetDocType(sculptor.JSON_DOCUMENT).
		SetQuery("domain", "host").
		SetQuery("source", "source").
		SetTargetStruct(&SubDomains{})
	doc_oneforall := sculptor.NewDataSculptor("hackerone.com.csv").
		SetDocType(sculptor.CSV_DOCUMENT).
		SetQuery("domain", "subdomain").
		SetQuery("source", "source").
		SetTargetStruct(&SubDomains{})
	common_output := sculptor.Merge(doc_subfinder, doc_oneforall)
	go doc_oneforall.Do()
	go doc_subfinder.Do()

	if common_output == nil {
		panic("merge error")
	}
	for i := range common_output {
		fmt.Printf("[*] subdomain %v found, via %v \n", i.(*SubDomains).Name, i.(*SubDomains).Source)
	}
}
