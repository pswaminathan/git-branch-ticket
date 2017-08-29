package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"time"

	"github.com/andygrunwald/go-jira"
)

const (
	// TicketRegex is the regular expression used to match a JIRA ticket.
	// Understood to be a sequence of capital letters, followed by a hyphen,
	// followed by a sequence of numbers.
	TicketRegex = "[A-Z]+-[\\d]+"
)

var (
	regex = regexp.MustCompile(TicketRegex)
)

func getCurrentBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--no-color")
	out := bytes.NewBuffer(nil)
	cmd.Stdout = out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	curBranchBytes, err := out.ReadBytes('\n')
	if err != nil {
		return "", err
	}
	return string(bytes.Trim(curBranchBytes, "* \n")), nil
}

func getTicket(key string, creds credentials) (*jira.Issue, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	jiraClient, err := jira.NewClient(client, creds.Base)
	if err != nil {
		return nil, err
	}
	jiraClient.Authentication.SetBasicAuth(creds.Username, creds.Password)

	issue, _, err := jiraClient.Issue.Get(key, nil)
	return issue, err
}

func main() {
	creds, err := loadCredentials()
	if err != nil {
		log.Fatal(err)
	}
	branchName, err := getCurrentBranch()
	if err != nil {
		log.Fatal(err)
	}
	if match := regex.MatchString(branchName); !match {
		log.Fatalf("Branch name must match the JIRA format of %v. This branch: %v\n", TicketRegex, branchName)
	}

	issue, err := getTicket(branchName, creds)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
}
