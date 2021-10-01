package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/emailquangto/redmine4go"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env-json")

	baseURL := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")
	apiFormat := os.Getenv("API_FORMAT")
	projectId, error := strconv.Atoi(os.Getenv("PROJECT_ID"))

	c := redmine4go.CreateClient(baseURL, apiKey, apiFormat)

	// default parameters of querying issues
	paras := &redmine4go.IssueListParameter{
		0,  // Offset
		25, // Limit
		"", // Sort
		"", // Include
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

	issueList, error := c.GetIssueListOfProject(projectId, paras, filters)
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

	issues, error := c.GetIssuesOfProject(projectId, paras, nil)
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
