package gotime

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(NewVersionCmd())
}

// NewVersionCmd creates the 'version' command which displays the version of the CLI tool.
func NewVersionCmd() *cobra.Command {
	var ReleaseNotes = false

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Shows CLI version",
		Long: `Shows the version of the gotime CLI tool.

Details about changes can be found at: https://github.com/urianchang/LearnGo/blob/master/internal/gotime/CHANGELOG.md.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprintln(cmd.OutOrStdout(), Version)
			if err != nil {
				return err
			}
			if ReleaseNotes {
				err = printReleaseChanges(cmd.OutOrStdout(), Version, CHANGELOG)
				if err != nil {
					return err
				}
			}
			return nil
		},
		SilenceUsage: true,
	}
	cmd.Flags().BoolVar(&ReleaseNotes, "verbose", false, "Print out release notes")
	return cmd
}

const (
	ReleaseHeaderPrefix = "## Release: "
	SectionHeaderPrefix = "### "
)

func printReleaseChanges(writer io.Writer, releaseVersion string, changelog []byte) error {
	shouldPrint := false
	reader := bytes.NewReader(changelog)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ReleaseHeaderPrefix) {
			trimmed := strings.TrimPrefix(line, ReleaseHeaderPrefix)
			versionInfo := strings.Split(trimmed, " ")
			if versionInfo[0] == releaseVersion {
				shouldPrint = true
				continue
			} else {
				shouldPrint = false
			}
		}

		if shouldPrint {
			_, err := fmt.Fprintln(writer, cleanLine(line))
			if err != nil {
				return fmt.Errorf("unable to print line: %v", err)
			}
		}
	}
	return nil
}

func embolden(words string) string {
	return "\033[1m" + words + "\033[0m"
}

// MarkdownLinkTextRegex will match and extract the link text of a markdown link.
// For example, from "[foo](bar)", "foo" is returned.
var MarkdownLinkTextRegex = regexp.MustCompile(`\[(.+?)]\(.+?\)`)

func cleanLine(line string) string {
	if strings.HasPrefix(line, SectionHeaderPrefix) {
		return embolden(strings.TrimPrefix(line, SectionHeaderPrefix))
	}
	return MarkdownLinkTextRegex.ReplaceAllString(line, "$1")
}
