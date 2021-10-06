package redmine4go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GetProjects() returns all projects (all public projects and private projects where user have access to)
// with given parameters
// from protocol scheme JSON
// Ref: https://www.redmine.org/projects/redmine/wiki/Rest_Projects#Listing-projects
func (c *Client) GetProjects(parameters string) (ProjectList, error) {
	// variable to store return value
	projectList := ProjectList{}

	// set up request
	req, err := http.NewRequest(http.MethodGet, c.url+"/projects."+c.format+"?include="+parameters, nil)
	if err != nil {
		return projectList, err
	}
	// add headers to the request
	req.Header.Add("Content-Type", "application/"+c.format)
	req.Header.Add("X-Redmine-API-Key", c.key)
	// send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return projectList, err
	}
	defer resp.Body.Close()

	// return error if status code is not OK
	if resp.StatusCode >= http.StatusBadRequest {
		return projectList, err
	}

	// parse the response's body
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return projectList, err
	}
	err = json.Unmarshal([]byte(bodyContent), &projectList)

	return projectList, err
}

// GetProject() returns details of a project with given parameters
// from protocol scheme JSON
// Ref: https://www.redmine.org/projects/redmine/wiki/Rest_Projects#Showing-a-project
func (c *Client) GetProject(projectIdOrName interface{}, parameters string) (Project, error) {
	// variable to store return value
	project := Project{}

	// set up request
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/projects/%v.%v?include=%v", c.url, projectIdOrName, c.format, parameters), nil)

	if err != nil {
		return project, err
	}
	// add headers to the request
	req.Header.Add("Content-Type", "application/"+c.format)
	req.Header.Add("X-Redmine-API-Key", c.key)
	// send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return project, err
	}
	defer resp.Body.Close()

	// return error if status code is not OK
	if resp.StatusCode >= http.StatusBadRequest {
		return project, err
	}

	// parse the response's body
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return project, err
	}
	projectWrapper := ProjectWrapper{}
	err = json.Unmarshal([]byte(bodyContent), &projectWrapper)

	return projectWrapper.Project, err
}

// ArchiveProject() archives the project of given id or identifier
// from protocol scheme JSON
// Ref: https://www.redmine.org/projects/redmine/wiki/Rest_Projects#Archiving-a-project
func (c *Client) ArchiveProject(projectIdOrName interface{}) error {

	// set up request
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%v/projects/%v/archive.%v", c.url, projectIdOrName, c.format), nil)

	if err != nil {
		return err
	}
	// add headers to the request
	req.Header.Add("Content-Type", "application/"+c.format)
	req.Header.Add("X-Redmine-API-Key", c.key)
	// send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// parse the response's body
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(bodyContent))

	return err
}

type ProjectList struct {
	Projects   []Project `json:"projects"`
	TotalCount int       `json:"total_count"`
	Offset     int       `json:"offset"`
	Limit      int       `json:"limit"`
}

type ProjectWrapper struct {
	Project Project `json:"project"`
}

type Project struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	Identifier     string         `json:"identifier"`
	Description    string         `json:"description"`
	Status         int            `json:"status"`
	IsPublic       bool           `json:"is_public"`
	InheritMembers bool           `json:"inherit_members"`
	CustomFields   []CustomFields `json:"custom_fields"`
	CreatedOn      time.Time      `json:"created_on"`
	UpdatedOn      time.Time      `json:"updated_on"`
}

type CustomFields struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Value    interface{} `json:"value"`
	Multiple bool        `json:"multiple,omitempty"`
}
