package gotime

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(NewParseCmd())
}

// NewParseCmd creates a command for parsing Unix time into a more human-friendly format.
func NewParseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parse",
		Args:  cobra.MinimumNArgs(1),
		Short: "Parse epoch time",
		Long: `Parses an epoch time value in seconds to a specified time zone. The
default time zone is local.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var epochSec, err = strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse input: %v", err)
			}
			var localTime = time.Unix(epochSec, 0)
			if len(args) == 2 {
				zone := args[1]
				loc, e := time.LoadLocation(zone)
				if e != nil {
					return fmt.Errorf("%s is not a supported time zone\n", zone)
				}
				localTime = localTime.In(loc)
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), localTime)
			return err
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	return cmd
}
