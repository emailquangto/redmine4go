package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/emailquangto/redmine4go"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")

	baseURL := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")
	apiFormat := os.Getenv("API_FORMAT")
	projectId, error := strconv.Atoi(os.Getenv("PROJECT_ID"))

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
		nil, // IssueId
		nil, // ProjectId
		nil, // SubprojectId
		nil, // TrackerId
		nil, // StatusId
		nil, // AssignedToId
		nil, // ParentId
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
