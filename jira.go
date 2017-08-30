package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/andygrunwald/go-jira"
)

var (
	client = &http.Client{Timeout: 10 * time.Second}
)

func getTicket(key string, creds credentials) (*jira.Issue, error) {
	jiraClient, err := jira.NewClient(client, creds.Base)
	if err != nil {
		return nil, err
	}
	jiraClient.Authentication.SetBasicAuth(creds.Username, creds.Password)

	issue, _, err := jiraClient.Issue.Get(key, nil)
	return issue, err
}

func getTickets(keys []string, creds credentials) ([]jira.Issue, error) {
	jiraClient, err := jira.NewClient(client, creds.Base)
	if err != nil {
		return nil, err
	}
	jiraClient.Authentication.SetBasicAuth(creds.Username, creds.Password)
	jql := fmt.Sprintf("key in (%s)", strings.Join(keys, ","))
	issues, _, err := jiraClient.Issue.Search(jql, nil)
	return issues, err
}
