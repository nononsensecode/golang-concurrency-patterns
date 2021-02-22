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
	"math/rand"
	"os"

	"github.com/spf13/cobra"
	"nononsensecode.com/go-concurrency-patterns/functions"
)

// cleanupOnQuitCmd represents the cleanupOnQuit command
var cleanupOnQuitCmd = &cobra.Command{
	Use:   "cleanupOnQuit",
	Short: "Does cleanup on quit",
	Long: `This will do any work pending, that has to be done before exiting`,
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := cmd.Flags().GetString("msg")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %q\n", err)
			os.Exit(1)
		}

		quit := make(chan string)
		c := functions.BoringCleanup(msg, quit)
		for i := rand.Intn(10); i >= 0; i-- {
			fmt.Println(<-c)
		}

		quit <- "Bye"
		fmt.Printf("Joue say: %q\n", <-quit)
	},
}

func init() {
	boringCmd.AddCommand(cleanupOnQuitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanupOnQuitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanupOnQuitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cleanupOnQuitCmd.Flags().String("msg", "Joe", "Message to print")
}
