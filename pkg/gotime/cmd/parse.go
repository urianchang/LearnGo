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
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(parseCmd)
}

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Args: cobra.MinimumNArgs(1),
	Short: "Parse epoch time",
	Long: `Parses an epoch time value in seconds to a specified time zone. The
default time zone is local.`,
	Run: func(cmd *cobra.Command, args []string) {
		var epochSec, err = strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Unable to parse input")
		}

		var localTime = time.Unix(epochSec, 0)
		if len(args) == 2 {
			zone := args[1]
			loc, err := time.LoadLocation(zone)
			if err != nil {
				fmt.Printf("%s is not a supported time zone\n", zone)
			} else {
				fmt.Println(localTime.In(loc))
			}
		} else {
			fmt.Println(localTime)
		}
	},
}


