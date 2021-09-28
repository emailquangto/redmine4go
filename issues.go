package redmine4go

import (
	"io/ioutil"
	"net/http"
)

// GetIssuesOfProject returns a list of issues
// in a project
// for the default settings (parameters)
func (c *Client) GetIssuesOfProject(projectId string) (string, error) {

	req, err := http.NewRequest(http.MethodGet, c.url+"/issues."+c.format+"?project_id="+projectId, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/"+c.format)
	req.Header.Add("X-Redmine-API-Key", c.key)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return resp.Status, err
	}

	return string(bodyContent), nil
}

// An Issue stores the issue information
type Issue struct {
	id               int
	project_id       int
	project_name     string
	tracker          string
	status           string
	priority         string
	author_id        int
	author_name      string
	assigned_to_id   int
	assigned_to_name string
	parent           int
	subject          string
	description      string
	start_date       string
	due_date         string
	done_ratio       int
	is_private       bool
	estimated_hours  float32
	created_on       string
	updated_on       string
	closed_on        string
}
