package redmine4go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// GetIssueListOfProject() returns a raw list of issues (including value of total count, offset, limit)
// in a project
// for the default settings (parameters) from protocol scheme JSON
func (c *Client) GetIssueListOfProject(projectId int, para *IssueListParameter, filter *IssueListFilter) (IssueList, error) {
	// variable to store return value
	issueList := IssueList{}

	// set up request
	query := generateIssueListQuery(para, filter)
	req, err := http.NewRequest(http.MethodGet, c.url+"/issues."+c.format+"?project_id="+strconv.Itoa(projectId)+query, nil)
	if err != nil {
		return issueList, err
	}
	// add headers to the request
	req.Header.Add("Content-Type", "application/"+c.format)
	req.Header.Add("X-Redmine-API-Key", c.key)
	// send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return issueList, err
	}
	defer resp.Body.Close()

	// return error if status code is not OK
	if resp.StatusCode >= http.StatusBadRequest {
		return issueList, err
	}

	// parse the response's body
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return issueList, err
	}
	err = json.Unmarshal([]byte(bodyContent), &issueList)

	return issueList, err
}

// GetIssuesOfProject() returns a list of issues only
// in a project
// for the default settings (parameters) from protocol scheme JSON
func (c *Client) GetIssuesOfProject(projectId int, para *IssueListParameter, filter *IssueListFilter) ([]Issue, error) {
	// variable to store return value
	issues := []Issue{}

	issueList, err := c.GetIssueListOfProject(projectId, para, filter)
	if err != nil {
		return issues, err
	}
	issues = issueList.Issues

	return issues, err
}

func generateIssueListQuery(para *IssueListParameter, filter *IssueListFilter) string {
	if para == nil {
		return ""
	}

	query := ""
	if para.Offset != nil {
		query += fmt.Sprintf("&offset=%v", para.Offset)
	}
	if para.Limit != nil {
		query += fmt.Sprintf("&limit=%v", para.Limit)
	}
	if para.Sort != "" {
		query += fmt.Sprintf("&sort=%v", para.Sort)
	}
	if para.Include != "" {
		query += fmt.Sprintf("&status_id=%v", para.Include)
	}

	if filter == nil {
		return query
	}

	if filter.IssueId != nil {
		query += fmt.Sprintf("&issue_id=%v", filter.IssueId)
	}
	if filter.ProjectId != nil {
		query += fmt.Sprintf("&project_id=%v", filter.ProjectId)
	}
	if filter.SubprojectId != nil {
		query += fmt.Sprintf("&subproject_id=%v", filter.SubprojectId)
	}
	if filter.TrackerId != nil {
		query += fmt.Sprintf("&tracker_id=%v", filter.TrackerId)
	}
	if filter.StatusId != nil {
		query += fmt.Sprintf("&status_id=%v", filter.StatusId)
	}
	if filter.AssignedToId != nil {
		query += fmt.Sprintf("&assigned_to_id=%v", filter.AssignedToId)
	}
	if filter.ParentId != nil {
		query += fmt.Sprintf("&parent_id=%v", filter.ParentId)
	}

	return query
}

type IssueListParameter struct {
	Offset  interface{}
	Limit   interface{}
	Sort    string
	Include string
}

type IssueListFilter struct {
	IssueId      interface{}
	ProjectId    interface{}
	SubprojectId interface{}
	TrackerId    interface{}
	StatusId     interface{}
	AssignedToId interface{}
	ParentId     interface{}
}

type IssueList struct {
	Issues     []Issue `json:"issues"`
	TotalCount int     `json:"total_count"`
	Offset     int     `json:"offset"`
	Limit      int     `json:"limit"`
}

type Issue struct {
	ID             int         `json:"id"`
	Project        BriefInfo   `json:"project"`
	Tracker        BriefInfo   `json:"tracker"`
	Status         BriefInfo   `json:"status"`
	Priority       BriefInfo   `json:"priority"`
	Author         BriefInfo   `json:"author"`
	AssignedTo     BriefInfo   `json:"assigned_to"`
	Parent         Parent      `json:"parent,omitempty"`
	Subject        string      `json:"subject"`
	Description    string      `json:"description"`
	StartDate      string      `json:"start_date"`
	DueDate        string      `json:"due_date"`
	DoneRatio      int         `json:"done_ratio"`
	IsPrivate      bool        `json:"is_private"`
	EstimatedHours interface{} `json:"estimated_hours"`
	CreatedOn      time.Time   `json:"created_on"`
	UpdatedOn      time.Time   `json:"updated_on"`
	ClosedOn       interface{} `json:"closed_on"`
}

type BriefInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Parent struct {
	ID int `json:"id"`
}
