package main

import (
	"fmt"
	"log"
	"regexp"
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

func currentBranch(creds credentials) {
	branchName, err := getCurrentBranch()
	if err != nil {
		log.Fatal(err)
	}
	if match := regex.MatchString(branchName); !match {
		log.Fatalf("Branch name must match the JIRA format. This branch: %v\n", branchName)
	}

	issue, err := getTicket(branchName, creds)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
}

func allBranches(creds credentials) {
	branches, err := getAllBranches()
	if err != nil {
		log.Fatal(err)
	}
	for i, b := range branches {
		if !regex.MatchString(b) {
			branches = append(branches[:i], branches[i+1:]...)
		}
	}
	if len(branches) == 0 {
		log.Fatal("no branches match JIRA format")
	}
	issues, err := getTickets(branches, creds)
	if err != nil {
		log.Fatal(err)
	}
	for _, issue := range issues {
		fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	}
}

func main() {
	parseFlags()

	creds, err := loadCredentials()
	if err != nil {
		log.Fatal(err)
	}

	if opts.AllBranches {
		allBranches(creds)
	} else {
		currentBranch(creds)
	}
}
