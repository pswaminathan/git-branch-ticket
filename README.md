# git-branch-ticket

Get the summary of a JIRA ticket when your branch is named after it.

## Usage

1. Set up a credentials file at `~/.config/git/jira.json` (and make it `0600`):
	`{"username": "myusername", "password": "mypassword", "base": "https://issues.mediamath.com/jira/"}`
2. Run `git branch-ticket`
3. Amaze

## TODO

1. Proper credentials storage (using `pass` or `keyring`)
2. Multi-branch support
3. Editing branch description with ticket summary
