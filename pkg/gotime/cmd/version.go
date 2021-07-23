package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number of gotime",
	Long:  `Prints the version number of gotime`,
	RunE: func(cmd *cobra.Command, args []string) error {
		version := Version{0, 0, 1, "dev"}
		_, err := fmt.Fprintln(cmd.OutOrStdout(), version.Stringify())
		if err != nil {
			return err
		}
		return nil
	},
}

type Version struct {
	Major int    `json:"major"`
	Minor int    `json:"minor"`
	Patch int    `json:"patch"`
	Note  string `json:"note"`
}

func (v *Version) Stringify() string {
	s := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.Note != "" {
		s += "+" + v.Note
	}
	return s
}
