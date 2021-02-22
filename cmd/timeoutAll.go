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
	"time"

	"github.com/spf13/cobra"
	"nononsensecode.com/go-concurrency-patterns/functions"
)

// timeoutAllCmd represents the timeoutAll command
var timeoutAllCmd = &cobra.Command{
	Use:   "timeoutAll",
	Short: "Timeout the whole process",
	Long: `Timeout the whole process a certain time`,
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := cmd.Flags().GetString("msg")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %q\n", err)
			os.Exit(1)
		}

		c := functions.BoringReturnsChannel(msg)
		timeout := time.After(1 * time.Second)
		for {
			select {
			case s := <-c:
				fmt.Printf("You say: %q\n", s)
			case <-timeout:
				fmt.Println("You're too slow. I'm leaving!")
				return
			}
		}
	},
}

func init() {
	boringCmd.AddCommand(timeoutAllCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timeoutAllCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timeoutAllCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	timeoutAllCmd.Flags().String("msg", "Joe", "Message to print")
}
