package main

import (
	"fmt"
	"os"

	"github.com/emailquangto/redmine4go"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	baseURL := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")
	projectId := os.Getenv("PROJECT_ID")

	response := redmine4go.GetIssuesOfProject(baseURL, apiKey, projectId)
	fmt.Println(response)
}
