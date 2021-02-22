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

// channelInputCmd represents the channelInput command
var channelInputCmd = &cobra.Command{
	Use:   "channelInput",
	Short: "Using a channel input",
	Long: `We will pass a channel as input to the boring function
	and boring function will write the message with counter to that
	channel`,
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := cmd.Flags().GetString("msg")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		c := make(chan string)
		go functions.BoringWithChannelInput(msg, c)
		for i := 0; i < 5; i++ {
			fmt.Printf("You say: %q\n", <-c)
		}
		fmt.Println("You're boring; I'm leaving!")
	},
}

func init() {
	withChannelCmd.AddCommand(channelInputCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// channelInputCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// channelInputCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	channelInputCmd.Flags().String("msg", "Kaushik", "Message to print")
}
