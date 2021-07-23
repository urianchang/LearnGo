// Copyright © 2019 Urian
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
