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

	// get list of open issues
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
		nil,       // IssueId
		projectId, // ProjectId
		nil,       // SubprojectId
		nil,       // TrackerId
		nil,       // StatusId
		nil,       // AssignedToId
		nil,       // ParentId
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
	issue, err := c.GetIssue(issueId, include)
	if error == nil {
		fmt.Printf("%s\n", "=====get details of an issue=====")
		fmt.Printf("issue - Project = %s\n", issue.Project.Name)
		fmt.Printf("issue - ID = %d\n", issue.ID)
		fmt.Printf("issue - Subject = %s\n", issue.Subject)
		fmt.Printf("issue - Status = %s\n", issue.Status.Name)
		fmt.Printf("issue - Author = %s\n", issue.Author.Name)
		fmt.Printf("issue - Assigned To = %s\n", issue.AssignedTo.Name)

	} else {
		fmt.Printf("%s\n", err)
	}
}
