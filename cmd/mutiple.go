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

// mutipleCmd represents the mutiple command
var multipleCmd = &cobra.Command{
	Use:   "multiple",
	Short: "Calling multiple boring functions",
	Long: `It will initiate 2 boring goroutines and will listen
	for their messages`,
	Run: func(cmd *cobra.Command, args []string) {
		msg1, err := cmd.Flags().GetString("msg1")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		msg2, err := cmd.Flags().GetString("msg2")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		c1 := functions.BoringReturnsChannel(msg1)
		c2 := functions.BoringReturnsChannel(msg2)
		for i := 0; i < 5; i++ {
			fmt.Println(<-c1)
			fmt.Println(<-c2)
		}
		fmt.Println("You both are boring; I'm leaving")
	},
}

func init() {
	boringCmd.AddCommand(multipleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mutipleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mutipleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	multipleCmd.Flags().String("msg1", "Joe", "First message to print")
	multipleCmd.Flags().String("msg2", "Ann", "Second message to print")
}
