package main

import (
	"bytes"
	"os/exec"
)

func getCurrentBranch() (string, error) {
	out, err := getBranchesShell()
	if err != nil {
		return "", err
	}
	curBranchBytes, err := out.ReadBytes('\n')
	if err != nil {
		return "", err
	}
	return string(bytes.Trim(curBranchBytes, "* \n")), nil
}

func getAllBranches() ([]string, error) {
	out, err := getBranchesShell()
	if err != nil {
		return nil, err
	}
	branchBytes := bytes.Split(out.Bytes(), []byte{'\n'})
	var branches []string
	for _, b := range branchBytes {
		if s := string(bytes.Trim(b, "* ")); s != "" {
			branches = append(branches, s)
		}
	}
	return branches, nil
}

func getBranchesShell() (*bytes.Buffer, error) {
	cmd := exec.Command("git", "branch", "--no-color")
	out := bytes.NewBuffer(nil)
	cmd.Stdout = out
	err := cmd.Run()
	return out, err
}
