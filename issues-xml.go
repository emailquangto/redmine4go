package redmine4go

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// GetIssueListOfProjectXML() returns a raw list of issues (including value of total count, offset, limit)
// in a project
// for the default settings (parameters) from protocol scheme XML
func (c *Client) GetIssueListOfProjectXML(projectId string) (IssueListXML, error) {
	// variable to store return value
	issueList := IssueListXML{}

	// set up request
	req, err := http.NewRequest(http.MethodGet, c.url+"/issues."+c.format+"?project_id="+projectId, nil)
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
	//defer resp.Body.Close()

	// return error if status code is not OK
	if resp.StatusCode >= http.StatusBadRequest {
		return issueList, err
	}

	// parse the response's body
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return issueList, err
	}
	err = xml.Unmarshal([]byte(bodyContent), &issueList)

	return issueList, err
}

// GetIssuesOfProjectXML() returns a list of issues only
// in a project
// for the default settings (parameters) from protocol scheme XML
func (c *Client) GetIssuesOfProjectXML(projectId string) ([]IssueXML, error) {
	// variable to store return value
	issues := []IssueXML{}

	issueList, err := c.GetIssueListOfProjectXML(projectId)
	if err != nil {
		return issues, err
	}
	issues = issueList.Issues

	return issues, err
}

type IssueListXML struct {
	XMLName    xml.Name   `xml:"issues"`
	Text       string     `xml:",chardata"`
	TotalCount string     `xml:"total_count,attr"`
	Offset     string     `xml:"offset,attr"`
	Limit      string     `xml:"limit,attr"`
	Type       string     `xml:"type,attr"`
	Issues     []IssueXML `xml:"issue"`
}

type IssueXML struct {
	Text           string       `xml:",chardata"`
	ID             string       `xml:"id"`
	Project        BriefInfoXML `xml:"project"`
	Tracker        BriefInfoXML `xml:"tracker"`
	Status         BriefInfoXML `xml:"status"`
	Priority       BriefInfoXML `xml:"priority"`
	Author         BriefInfoXML `xml:"author"`
	AssignedTo     BriefInfoXML `xml:"assigned_to"`
	Parent         BriefInfoXML `xml:"parent"`
	Subject        string       `xml:"subject"`
	Description    string       `xml:"description"`
	StartDate      string       `xml:"start_date"`
	DueDate        string       `xml:"due_date"`
	DoneRatio      string       `xml:"done_ratio"`
	IsPrivate      string       `xml:"is_private"`
	EstimatedHours string       `xml:"estimated_hours"`
	CreatedOn      string       `xml:"created_on"`
	UpdatedOn      string       `xml:"updated_on"`
	ClosedOn       string       `xml:"closed_on"`
}

type BriefInfoXML struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}
