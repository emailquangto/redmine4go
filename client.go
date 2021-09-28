package redmine4go

import (
	"net/http"
)

// A Client stores the client information
// and implement all the api functions
// to communicate with redmine
type Client struct {
	url        string
	key        string
	format     string
	httpClient *http.Client
}

// CreateClient returns a client
// with the given credential
// for the given redmine domain
func CreateClient(url, key, format string) (c *Client) {
	c = &Client{}
	c.url = url
	c.key = key
	c.format = format
	c.httpClient = &http.Client{}

	return
}
