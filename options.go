package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

var (
	opts options
)

type options struct {
	Version     bool     `short:"v" long:"version" description:"Show version number and exit"`
	AllBranches bool     `short:"a" long:"all-branches" description:"Retrieve ticket information for all JIRA-ticket-formatted branches"`
	Branches    []string `short:"b" long:"branch" description:"Retrieve ticket information for named branches"`
}

func parseFlags() {
	_, err := flags.Parse(&opts)
	if err != nil {
		e := err.(*flags.Error)
		if e.Type != flags.ErrHelp {
			fmt.Println(err)
		}
		os.Exit(1)
	}
	if opts.Version {
		fmt.Println("git-branch-ticket " + version)
		os.Exit(0)
	}
}
