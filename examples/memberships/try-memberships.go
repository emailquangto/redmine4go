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
	projectName := os.Getenv("PROJECT_NAME")

	c := redmine4go.CreateClient(baseURL, apiKey, apiFormat)

	// get memberships of a project
	projectMembership, error := c.GetProjectMemberships(projectId)
	if error == nil {
		fmt.Printf("%s\n", "=====get memberships of a project (using project Id)=====")
		fmt.Printf("Number of memberships = %d\n", projectMembership.TotalCount)
		if projectMembership.TotalCount > 0 {
			fmt.Printf("membership 1 - Project = %s\n", projectMembership.Memberships[0].Project.Name)
			fmt.Printf("membership 1 - Name = %s\n", projectMembership.Memberships[0].User.Name)
			fmt.Printf("membership 1 - Role = %s\n", projectMembership.Memberships[0].Roles[0].Name)
		}
	} else {
		fmt.Printf("%s\n", error)
	}

	projectMembership, error = c.GetProjectMemberships(projectName)
	if error == nil {
		fmt.Printf("%s\n", "=====get memberships of a project (using project name)=====")
		fmt.Printf("Number of memberships = %d\n", projectMembership.TotalCount)
		if projectMembership.TotalCount > 0 {
			fmt.Printf("membership 1 - Project = %s\n", projectMembership.Memberships[0].Project.Name)
			fmt.Printf("membership 1 - Name = %s\n", projectMembership.Memberships[0].User.Name)
			fmt.Printf("membership 1 - Role = %s\n", projectMembership.Memberships[0].Roles[0].Name)
		}
	} else {
		fmt.Printf("%s\n", error)
	}
}
