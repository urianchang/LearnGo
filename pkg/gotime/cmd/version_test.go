package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestChangelog = `
# CHANGELOG
## Release: 1.0.0+dev ([link_text](link_url))
### Summary
Some cool things are getting added.

### Features
* New thing 1 ([Jira_Ticket](link_to_Jira)), [#GitHub_PR](link_to_GitHub_PR)

### Bug Fixes
* Bug fix [#GitHub_PR](link_to_GitHub_PR)

## Release: 1.0.0 ([link_text](link_url))
### Summary
Woot! Version 1.0.0 achieved!

### Features
* New thing 1 ([Jira_Ticket](link_to_Jira)), [#GitHub_PR](link_to_GitHub_PR)
* New thing 2 ([Jira_Ticket](link_to_Jira)), [#GitHub_PR](link_to_GitHub_PR)

### Bug Fixes
No bug fixes.

## Release: 0.9.0 ([link_text](link_url))
### Summary
Bug fix release.

### Features
No new features.

### Bug Fixes
* Bug fix [#GitHub_PR](link_to_GitHub_PR)
`

// NewVersionCmd tests
func TestExecute_Version_Short(t *testing.T) {
	Version = "1.2.3+dev"
	buf := new(bytes.Buffer)
	cmd := NewVersionCmd()
	cmd.SetOut(buf)
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.EqualValues(t, Version+"\n", buf.String())
}

func TestExecute_Version_Verbose(t *testing.T) {
	Version = "0.9.0"
	CHANGELOG = []byte(TestChangelog)
	buf := new(bytes.Buffer)

	cmd := NewVersionCmd()
	cmd.SetOut(buf)
	cmd.SetArgs([]string{
		"--verbose",
	})
	err := cmd.Execute()
	assert.NoError(t, err)
	expected := fmt.Sprintf(`0.9.0
%s
Bug fix release.

%s
No new features.

%s
* Bug fix #GitHub_PR
`, embolden("Summary"), embolden("Features"), embolden("Bug Fixes"))
	assert.EqualValues(t, expected, buf.String())
}

func TestExecute_Version_NoMatchingReleaseNotes(t *testing.T) {
	Version = "no-release-notes"
	CHANGELOG = []byte(TestChangelog)
	buf := new(bytes.Buffer)

	cmd := NewVersionCmd()
	cmd.SetOut(buf)
	cmd.SetArgs([]string{
		"--verbose",
	})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.EqualValues(t, Version+"\n", buf.String())
}

// printReleaseChanges tests
func TestPrintReleaseChanges(t *testing.T) {
	buf := new(bytes.Buffer)
	err := printReleaseChanges(buf, "1.0.0", []byte(TestChangelog))
	assert.NoError(t, err)
	expected := fmt.Sprintf(`%s
Woot! Version 1.0.0 achieved!

%s
* New thing 1 (Jira_Ticket), #GitHub_PR
* New thing 2 (Jira_Ticket), #GitHub_PR

%s
No bug fixes.

`, embolden("Summary"), embolden("Features"), embolden("Bug Fixes"))
	assert.EqualValues(t, expected, buf.String())
}

// Sanity checking the embolden function that wraps provided text
func TestEmbolden(t *testing.T) {
	assert.EqualValues(t, "\033[1mHello World!\033[0m", embolden("Hello World!"))
}
