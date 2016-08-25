// Package github provides a GO API for the GitHub issue trakcer.
package github

import "time"

// IssuesURL : The github endpoint for issues
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult : The Search result
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue : A single issue
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	REPOURL   string `json:"repository_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User : a user
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
