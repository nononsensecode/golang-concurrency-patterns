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

// channelOuputCmd represents the channelOuput command
var channelOutputCmd = &cobra.Command{
	Use:   "channelOutput",
	Short: "Will output a channel",
	Long: `This function will receive a message and outputs a channel
	which in turn stream messages with count`,
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := cmd.Flags().GetString("msg")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		c := functions.BoringReturnsChannel(msg)
		for i := 0; i < 5; i++ {
			fmt.Printf("You say: %q\n", <-c)
		}
		fmt.Println("You're boring; I'm leaving")
	},
}

func init() {
	withChannelCmd.AddCommand(channelOutputCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// channelOuputCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// channelOuputCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	channelOutputCmd.Flags().String("msg", "Kaushik", "Message to print")
}
