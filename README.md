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
	godotenv.Load("../.env")

	baseURL := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")
	apiFormat := os.Getenv("API_FORMAT")
	projectId := os.Getenv("PROJECT_ID")

	c := redmine4go.CreateClient(baseURL, apiKey, apiFormat)

	// default parameters of querying issues
	paras := &redmine4go.IssueListParameter{
		nil, // Offset = 0
		nil, // Limit = 25
		"",  // Sort
		"",  // Include
	}

	// default filters of querying issues
	filters := &redmine4go.IssueListFilter{
		nil, // IssueId = an integer
		nil, // ProjectId = an integer or "project-name"
		nil, // SubprojectId = an integer or "sub-project-name"
		nil, // TrackerId = an integer
		nil, // StatusId = an integer or "status-name"
		nil, // AssignedToId = an integer or "member-name"
		nil, // ParentId = an integer
	}

	// get list of open issues that user can access
	issueList, error := c.GetIssues(paras, filters)
	if error == nil {
		fmt.Printf("Number of issues = %d\n", issueList.TotalCount)
		if issueList.TotalCount > 0 {
			fmt.Printf("issue 1 - Project = %s\n", issueList.Issues[0].Project.Name)
			fmt.Printf("issue 1 - ID = %d\n", issueList.Issues[0].ID)
			fmt.Printf("issue 1 - Subject = %s\n", issueList.Issues[0].Subject)
			fmt.Printf("issue 1 - Status = %s\n", issueList.Issues[0].Status.Name)
			fmt.Printf("issue 1 - Author = %s\n", issueList.Issues[0].Author.Name)
			fmt.Printf("issue 1 - Assigned To = %s\n", issueList.Issues[0].AssignedTo.Name)
		}
	} else {
		fmt.Printf("%s\n", error)
	}

	fmt.Printf("%s\n", "=====*****=====")

	// default filters of querying issues
	filters = &redmine4go.IssueListFilter{
		nil,       // IssueId
		projectId, // ProjectId
		nil,       // SubprojectId
		nil,       // TrackerId
		nil,       // StatusId
		nil,       // AssignedToId
		nil,       // ParentId
	}

	// get list of open issues of a project that user can access
	issueList, error = c.GetIssues(paras, filters)
	if error == nil {
		fmt.Printf("Number of issues = %d\n", issueList.TotalCount)
		if issueList.TotalCount > 0 {
			fmt.Printf("issue 1 - Project = %s\n", issueList.Issues[0].Project.Name)
			fmt.Printf("issue 1 - ID = %d\n", issueList.Issues[0].ID)
			fmt.Printf("issue 1 - Subject = %s\n", issueList.Issues[0].Subject)
			fmt.Printf("issue 1 - Status = %s\n", issueList.Issues[0].Status.Name)
			fmt.Printf("issue 1 - Author = %s\n", issueList.Issues[0].Author.Name)
			fmt.Printf("issue 1 - Assigned To = %s\n", issueList.Issues[0].AssignedTo.Name)
		}
	} else {
		fmt.Printf("%s\n", error)
	}
}
```

### See more examples under _examples_ folder.


## License
   -------

[MIT License](https://github.com/emailquangto/redmine4go/blob/master/LICENSE)
