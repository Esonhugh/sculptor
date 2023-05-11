# Sculptor

![FlexFormer Logo](./doc/logo.png)
[![Go Report Card](https://goreportcard.com/badge/github.com/esonhugh/sculptor)](https://goreportcard.com/report/github.com/esonhugh/sculptor)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


Sculptor is a flexible and powerful Go library 
for transforming data from various formats 
(CSV, JSON, etc.) into desired Go struct types. 

It is designed to significantly simplify and accelerate the process of data ingestion and formatting.

## Features

- **Ease of use**: Simply define your struct and let sculptor handle the rest.
- **Support for multiple data formats**: CSV, JSON, jQuery-like queries, and more.
- **Database integration**: Directly ingest data from various databases.
- **Efficient handling of large data sets**: Optimize memory usage and performance.

## Installation

```bash
go get github.com/esonhugh/sculptor
```

## Quick Start

```go
package main

import (
        "github.com/esonhugh/sculptor"
	"log"
)

type TestStruct struct {
	Name string `select:"name"`
	Pass string `select:"pass"`
}

func main() {
	Doc := sculptor.NewDataSculptor("test.json").
		SetDocType(sculptor.JSON_DOCUMENT).
		SetQuery("name", "user").
		SetQuery("pass", "pass").
		SetTargetStruct(&TestStruct{})
	go Doc.Do()
	for i := range Doc.ConstructedOutput {
		log.Print(i)
	}
}
```

## More Example

Checkout the [test](./test) folder to see more examples.

## Contributing

welcome contributions! Free for fork and Pull Request. 

## License

sculptor is licensed under the MIT License. See [LICENSE](LICENSE) for more details.

---

Made with :heart: by [Esonhugh](https://eson.ninja)
