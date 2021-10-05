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
	//projectId := os.Getenv("PROJECT_ID")

	c := redmine4go.CreateClient(baseURL, apiKey, apiFormat)

	include := "" // trackers, issue_categories, enabled_modules
	projectList, error := c.GetProjects(include)
	if error == nil {
		fmt.Printf("%s\n", "=====get list of projects=====")
		fmt.Printf("Number of projects = %d\n", projectList.TotalCount)
		if projectList.TotalCount > 0 {
			fmt.Printf("project 1 - Name = %s\n", projectList.Projects[0].Name)
			fmt.Printf("project 1 - ID = %d\n", projectList.Projects[0].ID)
			fmt.Printf("project 1 - Status = %d\n", projectList.Projects[0].Status)
		}
	} else {
		fmt.Printf("%s\n", error)
	}

}
