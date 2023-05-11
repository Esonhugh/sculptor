/*
Package sculptor is flexible and powerful Go library
for transforming data from various formats
(CSV, JSON, etc.) into desired Go struct types.
Getting Start with

	package main

	import (
		"fmt"
		"github.com/esonhugh/sculptor"
	)

	type TestStruct struct {
		Name string `select:"name"`
		Pass string `select:"pass"`
	}

	// json like `{"user": "username","pass": "114514","content": "123444", "id":2}`
	// csv like `user,pass,content,id\nusername,114514,123444,2`

	Doc := sculptor.NewDataSculptor("your_test.json").
		SetDocType(sculptor.JSON_DOCUMENT).
		SetQuery("name", "user").
		SetQuery("pass", "pass").
		SetTargetStruct(&TestStruct{})
	go Doc.Do()
	for i := range Doc.ConstructedOutput {
		fmt.Println(i)
	}
*/
package sculptor
