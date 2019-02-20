package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// IssuesURL ...
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult ...
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue ...
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User ...
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
