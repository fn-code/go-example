// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		val, err := cmd.PersistentFlags().GetBool("toggle")
		if err != nil {
			return err
		}
		fmt.Println("Status : ", val)
		str, err := cmd.PersistentFlags().GetString("list")
		if err != nil {
			return err
		}
		fmt.Println("List : ", str)
		fmt.Println("run called ", args)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.
	runCmd.PersistentFlags().BoolP("toggle", "t", false, "Help message for toggle")
	runCmd.PersistentFlags().StringP("list", "l", "all", "show list data with 'all' or 'success' ")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

}
