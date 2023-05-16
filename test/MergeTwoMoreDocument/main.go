package main

import (
	"fmt"
	"github.com/esonhugh/sculptor"
	"sync"
)

type SubDomains struct {
	Name   string `select:"domain"`
	Source string `select:"source"`
}

// Demo of merge subfinder and oneforall
func main() {
	wg := &sync.WaitGroup{}
	doc_subfinder := sculptor.NewDataSculptorWithWg("subfinder_output.json", wg).
		SetDocType(sculptor.JSON_DOCUMENT).
		SetQuery("domain", "host").
		SetQuery("source", "source").
		SetTargetStruct(&SubDomains{})
	doc_oneforall := sculptor.NewDataSculptorWithWg("hackerone.com.csv", wg).
		SetDocType(sculptor.CSV_DOCUMENT).
		SetQuery("domain", "subdomain").
		SetQuery("source", "source").
		SetTargetStruct(&SubDomains{})
	/* 	// Method 1 using merge and Do by yourself
	common_output := sculptor.Merge(doc_subfinder, doc_oneforall)

	doc_oneforall.Do()
	doc_subfinder.Do()

	go func() {
		wg.Wait()
		close(common_output)
	}()
	*/
	/*  // Method 2 using MergeAndDo
	common_output := sculptor.MergeAndDo(doc_subfinder, doc_oneforall)
	*/
	/*
		// Method 3 using MergeAndDoWithWg
		common_output := sculptor.MergeAndDoWithWg(wg, doc_subfinder, doc_oneforall)
	*/
	/*
		// Method 4 using MergeV2
		commWg, common_output := sculptor.MergeV2(doc_subfinder, doc_oneforall)
		doc_oneforall.Do()
		doc_subfinder.Do()
		go func() {
			commWg.Wait()
			close(common_output)
		}()
	*/
	// Method 5 using MergeV3
	common_output := sculptor.MergeV3(doc_subfinder, doc_oneforall)
	doc_oneforall.Do()
	doc_subfinder.Do()
	sculptor.AutoCloseV3()

	if common_output == nil {
		panic("merge error")
	}
	for i := range common_output {
		fmt.Printf("[*] subdomain %v found, via %v \n", i.(*SubDomains).Name, i.(*SubDomains).Source)
	}
}
