/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/scottcrawford03/felix/internal/service"
	"github.com/spf13/cobra"
)

// fixitCmd represents the version command
var fixitCmd = &cobra.Command{
	Use:   "fixit",
	Short: "inits a new golang project",
	Long:  `fixit inits a new golang project with a file structure.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.Init(); err != nil {
			fmt.Printf("Something went wrong: %s", err.Error())

			return
		}
		fmt.Println("All done!")
	},
}

func init() {
	rootCmd.AddCommand(fixitCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fixitCmd.PersistentFlags().String("foo", "", "A help for foo")f

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fixitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
