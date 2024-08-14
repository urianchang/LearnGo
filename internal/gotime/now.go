package gotime

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var currentTime time.Time

func init() {
	currentTime = time.Now()
	rootCmd.AddCommand(NewNowCmd())
}

// NewNowCmd creates a command for getting the current time.
func NewNowCmd() *cobra.Command {
	var getEpoch bool

	cmd := &cobra.Command{
		Use:   "now",
		Args:  cobra.MaximumNArgs(1),
		Short: "Get the current time",
		Long: `Get the current time in a specified IANA time zone. The default time
zone is local.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var output interface{} = currentTime
			if getEpoch {
				output = currentTime.Unix()
			}
			if len(args) == 1 {
				zone := args[0]
				loc, err := time.LoadLocation(zone)
				if err != nil {
					return fmt.Errorf("%s is not a supported time zone", zone)
				}
				output = currentTime.In(loc)
			}
			_, err := fmt.Fprintln(cmd.OutOrStdout(), output)
			return err
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.PersistentFlags().BoolVarP(&getEpoch, "epoch", "e", false, "Returns the epoch time in seconds")
	return cmd
}
