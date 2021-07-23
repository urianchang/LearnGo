package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var getEpoch bool

// nowCmd represents the now command
var nowCmd = &cobra.Command{
	Use:   "now",
	Args: cobra.MaximumNArgs(1),
	Short: "Get the current time",
	Long: `Get the current time in a specified IANA time zone. The default time 
zone is local.`,
	Run: func(cmd *cobra.Command, args []string) {
		if getEpoch {
			fmt.Println(time.Now().Unix())
			return
		}

		if len(args) == 1 {
			zone := args[0]
			loc, err := time.LoadLocation(zone)
			if err != nil {
				fmt.Printf("%s is not a supported time zone\n", zone)
			} else {
				fmt.Println(time.Now().In(loc))
			}
		} else {
			fmt.Println(time.Now())
		}
	},
}

func init() {
	rootCmd.AddCommand(nowCmd)
	nowCmd.PersistentFlags().BoolVarP(&getEpoch, "epoch", "e", false,"Returns the epoch time in seconds")
}
