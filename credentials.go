package main

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
)

var (
	confFile string
)

type credentials struct {
	Username string
	Password string
	Base     string
}

func init() {
	usr, _ := user.Current()
	confFile = filepath.Join(usr.HomeDir, ".config", "git", "jira.json")
}

func loadCredentials() (credentials, error) {
	var creds credentials
	f, err := os.Open(confFile)
	if err != nil {
		return creds, err
	}
	err = json.NewDecoder(f).Decode(&creds)
	return creds, err
}
