redmine4go
--------


redmine4go is a Go client library for accessing the [Redmine](https://www.redmine.org/projects/redmine/wiki/Rest_api/) API

**travis-ci:** [![Build Status](https://travis-ci.org/emailquangto/redmine4go.svg?branch=master)](https://travis-ci.org/emailquangto/redmine4go)

**GoDoc:** [![GoDoc](https://godoc.org/github.com/emailquangto/redmine4go?status.svg)](https://godoc.org/github.com/emailquangto/redmine4go)

**Test Coverage:** xx%


References
----------

[https://godoc.org/github.com/emailquangto/redmine4go](https://godoc.org/github.com/emailquangto/redmine4go)


Example usage
-------------

The following is how to get issues of a project.
See more samples under _examples_ folder.

```
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


License
-------

MIT