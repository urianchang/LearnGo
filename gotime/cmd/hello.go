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
	"github.com/spf13/cobra"
	"os/user"
	"time"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Says hello",
	Long:  `Greeting is based on the time of day`,
	Run: func(cmd *cobra.Command, args []string) {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		hour := time.Now().Hour()
		greeting := "Good evening"
		if hour >= 6 && hour < 12 {
			greeting = "Good morning"
		}
		if hour >= 12 && hour < 16 {
			greeting = "Good afternoon"
		}
		fmt.Printf("%s, %s\n", greeting, user.Name)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
