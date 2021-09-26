package redmine4go

import (
	"io/ioutil"
	"net/http"
)

func GetIssuesOfProject(baseURL, apiKey, projectId string) string {
	request, _ := http.NewRequest(http.MethodGet, baseURL+"/projects/"+projectId+"/issues.json", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Redmine-API-Key", apiKey)

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		return "Error in getting issues of project " + projectId
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return string(data)
	}
}
