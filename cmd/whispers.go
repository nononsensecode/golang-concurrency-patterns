/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/spf13/cobra"
	"nononsensecode.com/go-concurrency-patterns/functions"
)

// whispersCmd represents the whispers command
var whispersCmd = &cobra.Command{
	Use:   "whispers",
	Short: "Chinese whispers game",
	Long: `This will implement a chinese whispers game`,
	Run: func(cmd *cobra.Command, args []string) {
		size, err := cmd.Flags().GetInt("size")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %q\n", err)
			os.Exit(1)
		}

		leftMost := make(chan int)
		left := leftMost
		var right chan int

		for i := 0; i < size; i++ {
			right = make(chan int)
			go functions.PassWhisper(left, right)
			left = right
		}

		right <- 1
		fmt.Printf("Last count: %d\n", <-leftMost)
	},
}

func init() {
	rootCmd.AddCommand(whispersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// whispersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// whispersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	whispersCmd.Flags().Int("size", 10000, "Size of the team")
}
