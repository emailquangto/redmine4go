# Redmine API in Go

This library (in progress) supports most if not all of the `Redmine` REST calls.

[![GoDoc](https://godoc.org/github.com/emailquangto/redmine4go?status.svg)](https://godoc.org/github.com/emailquangto/redmine4go) 
[![lint](https://github.com/emailquangto/redmine4go/workflows/golangci-lint/badge.svg?branch=main)](https://github.com/emailquangto/redmine4go/actions?query=workflow%3A%22golangci-lint%22)

### [Redmine API Reference](https://www.redmine.org/projects/redmine/wiki/Rest_api) for details of parameters

## Interfaces

|API                |Implements|Functions      |
|-------------------|----------|---------------|
|Issues             |      100%|	       	   |
|		    		|          |- GetIssues()  |
|             	    |          |- GetIssue()   |
|             	    |          |- CreateIssue()|
|             	    |          |- UpdateIssue()|
|             	    |          |- DeleteIssue()|
|             	    |          |-  AddWatcher()|
|             	    |          |- RemoveWatcher()|
|Projects             |      10%|	       	   |

## Installation

### *go get*

    $ go get -u github.com/emailquangto/redmine4go


## Example
   -------

### CRUD of issues using protocol scheme **JSON**

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
		Offset:  nil, // = 0
		Limit:   nil, // = 25
		Sort:    "",  // Default
		Include: "",  // None
	}

	// get list of open issues
	// default filters of querying issues
	filters := &redmine4go.IssueListFilter{
		IssueId:      nil, // an integer
		ProjectId:    nil, // an integer or "project-name"
		SubprojectId: nil, // an integer or "sub-project-name"
		TrackerId:    nil, // an integer
		StatusId:     nil, // an integer or "status-name"
		AssignedToId: nil, // an integer or "member-name"
		ParentId:     nil, // an integer
	}
	issueList, error := c.GetIssues(paras, filters)
	if error == nil {
		fmt.Printf("%s\n", "=====get list of open issues=====")
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

	// get list of open issues of a project
	// filters of querying issues
	filters = &redmine4go.IssueListFilter{
		IssueId:      nil,
		ProjectId:    projectId,
		SubprojectId: nil,
		TrackerId:    nil,
		StatusId:     nil,
		AssignedToId: nil,
		ParentId:     nil,
	}
	issueList, error = c.GetIssues(paras, filters)
	if error == nil {
		fmt.Printf("%s\n", "=====get list of open issues of a project=====")
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

	// get details of an issue
	issueId := 13430
	include := "" // children, attachments, relations, changesets, journals, watchers, allowed_statuses
	issue, error := c.GetIssue(issueId, include)
	if error == nil {
		fmt.Printf("%s\n", "=====get details of an issue=====")
		fmt.Printf("issue - Project = %s\n", issue.Project.Name)
		fmt.Printf("issue - ID = %d\n", issue.ID)
		fmt.Printf("issue - Subject = %s\n", issue.Subject)
		fmt.Printf("issue - Status = %s\n", issue.Status.Name)
		fmt.Printf("issue - Author = %s\n", issue.Author.Name)
		fmt.Printf("issue - Assigned To = %s\n", issue.AssignedTo.Name)

	} else {
		fmt.Printf("%s\n", error)
	}

	// create a new issue
	issueNew := redmine4go.IssueToSend{
		Project:     16,
		Tracker:     1,
		Status:      1,
		Priority:    2,
		Subject:     "a new issue auto-posted from redmine4go",
		Description: "testing CreateIssue() of Redmine API in Go",
	}
	issueNewWrapper := redmine4go.IssueToSendWrapper{Issue: issueNew}
	issueNewReturn, error := c.CreateIssue(issueNewWrapper)
	if error == nil {
		fmt.Printf("%s\n", "=====create a new issue=====")
		fmt.Printf("issue - Project = %s\n", issueNewReturn.Project.Name)
		fmt.Printf("issue - ID = %d\n", issueNewReturn.ID)
		fmt.Printf("issue - Subject = %s\n", issueNewReturn.Subject)
		fmt.Printf("issue - Status = %s\n", issueNewReturn.Status.Name)
		fmt.Printf("issue - Author = %s\n", issueNewReturn.Author.Name)
		fmt.Printf("issue - Assigned To = %s\n", issueNewReturn.AssignedTo.Name)

	} else {
		fmt.Printf("%s\n", error)
	}

	// update an issue
	issueUpdateWrapper := redmine4go.IssueToSendWrapper{Issue: redmine4go.IssueToSend{
		Status:      2,
		Priority:    1,
		Subject:     "from code",
		Description: "code changed",
	}}
	error = c.UpdateIssue(issueNewReturn.ID, issueUpdateWrapper)
	if error == nil {
		fmt.Printf("%s\n", "=====update an issue=====")
		// get details of updated issue
		include := ""
		issue, error := c.GetIssue(issueNewReturn.ID, include)
		if error == nil {
			fmt.Printf("%s\n", "**details of updated issue**")
			fmt.Printf("issue - Project = %s\n", issue.Project.Name)
			fmt.Printf("issue - ID = %d\n", issue.ID)
			fmt.Printf("issue - Status updated = %s\n", issue.Status.Name)
			fmt.Printf("issue - Priority updated = %s\n", issue.Priority.Name)
			fmt.Printf("issue - Subject updated = %s\n", issue.Subject)
			fmt.Printf("issue - Description updated = %s\n", issue.Description)
			fmt.Printf("issue - Author = %s\n", issue.Author.Name)
			fmt.Printf("issue - Assigned To = %s\n", issue.AssignedTo.Name)
		} else {
			fmt.Printf("%s\n", error)
		}
	}

	// delete an issue
	error = c.DeleteIssue(issueNewReturn.ID)
	if error == nil {
		fmt.Printf("%s\n", "=====delete an issue=====")
		fmt.Printf("issue %d deleted", issueNewReturn.ID)
	} else {
		fmt.Printf("%s\n", error)
	}

}
```

### See more examples under _examples_ folder.


## License
   -------

[MIT License](https://github.com/emailquangto/redmine4go/blob/master/LICENSE)
