package redmine4go

import (
	"io/ioutil"
	"net/http"
)

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
