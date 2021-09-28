Redmine API in Go [![GoDoc](https://godoc.org/github.com/emailquangto/redmine4go?status.svg)](https://godoc.org/github.com/emailquangto/redmine4go) [![lint](https://github.com/emailquangto/redmine4go/workflows/lint/badge.svg?branch=master)](https://github.com/emailquangto/redmine4go/actions?query=workflow%3A%22lint%22)
[![codecov](https://codecov.io/gh/emailquangto/redmine4go/branch/master/graph/badge.svg)](https://codecov.io/gh/emailquangto/redmine4go)
===============

This library supports most if not all of the `Redmine` REST calls.


## Installation

### *go get*

    $ go get -u github.com/emailquangto/redmine4go

## Example

### Get issues of a Redmine project

```go
package main

import (
	"fmt"
	"os"

	"github.com/emailquangto/redmine4go"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	baseURL := os.Getenv("BASE_URL")        // https://redmine.domain-name.com
	apiKey := os.Getenv("API_KEY")          // xxxxxa9a660079fe55yyyyy22979c9fa015xxxxx
	apiFormat := os.Getenv("API_FORMAT")    // json or xml
	projectId := os.Getenv("PROJECT_ID")    // 1

	c := redmine4go.CreateClient(baseURL, apiKey, apiFormat)
	resp, err := c.GetIssuesOfProject(projectId)
	if err == nil {
		fmt.Println(resp)
	}
}
```

### See more examples under _examples_ folder.


## License

[MIT License](https://github.com/emailquangto/redmine4go/blob/master/LICENSE)