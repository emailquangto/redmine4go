package redmine4go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetProjects() Returns a paginated list of the project memberships
// project_id can be either the project numerical id or the project identifier
// from protocol scheme JSON
// Ref: https://www.redmine.org/projects/redmine/wiki/Rest_Memberships#Project-Memberships
func (c *Client) GetProjectMemberships(projectIdOrName interface{}) (ProjectMembership, error) {
	// variable to store return value
	projectMembership := ProjectMembership{}

	// set up request
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/projects/%v/memberships.%v", c.url, projectIdOrName, c.format), nil)
	if err != nil {
		return projectMembership, err
	}
	// add headers to the request
	req.Header.Add("Content-Type", "application/"+c.format)
	req.Header.Add("X-Redmine-API-Key", c.key)
	// send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return projectMembership, err
	}
	defer resp.Body.Close()

	// return error if status code is not OK
	if resp.StatusCode >= http.StatusBadRequest {
		return projectMembership, err
	}

	// parse the response's body
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return projectMembership, err
	}
	err = json.Unmarshal([]byte(bodyContent), &projectMembership)

	return projectMembership, err
}

type ProjectMembership struct {
	Memberships []Membership `json:"memberships"`
	TotalCount  int          `json:"total_count"`
	Offset      int          `json:"offset"`
	Limit       int          `json:"limit"`
}

type Membership struct {
	ID      int         `json:"id"`
	Project BriefInfo   `json:"project"`
	User    BriefInfo   `json:"user"`
	Roles   []BriefInfo `json:"roles"`
}
