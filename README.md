# Redmine API in Go

This library supports most if not all of the `Redmine` REST calls.

[![GoDoc](https://godoc.org/github.com/emailquangto/redmine4go?status.svg)](https://godoc.org/github.com/emailquangto/redmine4go) 
[![lint](https://github.com/emailquangto/redmine4go/workflows/golangci-lint/badge.svg?branch=main)](https://github.com/emailquangto/redmine4go/actions?query=workflow%3A%22golangci-lint%22)


## Installation

### *go get*

    $ go get -u github.com/emailquangto/redmine4go


## Example
   -------

### Get issues of a Redmine project using protocol scheme **JSON**

```go
package main

import (
	"fmt"
	"os"

	"github.com/emailquangto/redmine4go"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env-json")

	baseURL := os.Getenv("BASE_URL")        // https://redmine.domain-name.com
	apiKey := os.Getenv("API_KEY")          // xxxxxa9a660079fe55yyyyy22979c9fa015xxxxx
	apiFormat := os.Getenv("API_FORMAT")    // json
	projectId := os.Getenv("PROJECT_ID")    // 1

	c := redmine4go.CreateClient(baseURL, apiKey, apiFormat)

	issueList, error := c.GetIssueListOfProject(projectId)
	if error == nil {
		fmt.Printf("Number of issues = %d\n", issueList.TotalCount)
		fmt.Printf("issue 1 - Project = %s\n", issueList.Issues[0].Project.Name)
		fmt.Printf("issue 1 - ID = %d\n", issueList.Issues[0].ID)
		fmt.Printf("issue 1 - Subject = %s\n", issueList.Issues[0].Subject)
		fmt.Printf("issue 1 - Status = %s\n", issueList.Issues[0].Status.Name)
		fmt.Printf("issue 1 - Author = %s\n", issueList.Issues[0].Author.Name)
		fmt.Printf("issue 1 - Assigned To = %s\n", issueList.Issues[0].AssignedTo.Name)
	} else {
		fmt.Printf("%s\n", error)
	}

	fmt.Printf("%s\n", "=====*****=====")

	issues, error := c.GetIssuesOfProject(projectId)
	if error == nil {
		fmt.Printf("Number of issues = %d\n", len(issues))
		fmt.Printf("issue 1 - Project = %s\n", issues[0].Project.Name)
		fmt.Printf("issue 1 - ID = %d\n", issues[0].ID)
		fmt.Printf("issue 1 - Subject = %s\n", issues[0].Subject)
		fmt.Printf("issue 1 - Status = %s\n", issues[0].Status.Name)
		fmt.Printf("issue 1 - Author = %s\n", issues[0].Author.Name)
		fmt.Printf("issue 1 - Assigned To = %s\n", issues[0].AssignedTo.Name)
	} else {
		fmt.Printf("%s\n", error)
	}
}
```

### Get issues of a Redmine project using protocol scheme **XML**

```go
package main

import (
	"fmt"
	"os"

	"github.com/emailquangto/redmine4go"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env-xml")

	baseURL := os.Getenv("BASE_URL")        // https://redmine.domain-name.com
	apiKey := os.Getenv("API_KEY")          // xxxxxa9a660079fe55yyyyy22979c9fa015xxxxx
	apiFormat := os.Getenv("API_FORMAT")    // xml
	projectId := os.Getenv("PROJECT_ID")    // 1

	c := redmine4go.CreateClient(baseURL, apiKey, apiFormat)

	issueList, error := c.GetIssueListOfProjectXML(projectId)
	if error == nil {
		fmt.Printf("Number of issues = %s\n", issueList.TotalCount)
		fmt.Printf("issue 1 - Project = %s\n", issueList.Issues[0].Project.Name)
		fmt.Printf("issue 1 - ID = %s\n", issueList.Issues[0].ID)
		fmt.Printf("issue 1 - Subject = %s\n", issueList.Issues[0].Subject)
		fmt.Printf("issue 1 - Status = %s\n", issueList.Issues[0].Status.Name)
		fmt.Printf("issue 1 - Author = %s\n", issueList.Issues[0].Author.Name)
		fmt.Printf("issue 1 - Assigned To = %s\n", issueList.Issues[0].AssignedTo.Name)
	} else {
		fmt.Printf("%s\n", error)
	}

	fmt.Printf("%s\n", "=====*****=====")

	issues, error := c.GetIssuesOfProjectXML(projectId)
	if error == nil {
		fmt.Printf("Number of issues = %d\n", len(issues))
		fmt.Printf("issue 1 - Project = %s\n", issues[0].Project.Name)
		fmt.Printf("issue 1 - ID = %d\n", issues[0].ID)
		fmt.Printf("issue 1 - Subject = %s\n", issues[0].Subject)
		fmt.Printf("issue 1 - Status = %s\n", issues[0].Status.Name)
		fmt.Printf("issue 1 - Author = %s\n", issues[0].Author.Name)
		fmt.Printf("issue 1 - Assigned To = %s\n", issues[0].AssignedTo.Name)
	} else {
		fmt.Printf("%s\n", error)
	}
}
```

### See more examples under _examples_ folder.


## License
   -------

[MIT License](https://github.com/emailquangto/redmine4go/blob/master/LICENSE)
