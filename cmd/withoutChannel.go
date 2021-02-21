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
	"time"

	"github.com/spf13/cobra"
)

// withoutChannelCmd represents the withoutChannel command
var withoutChannelCmd = &cobra.Command{
	Use:   "withoutChannel",
	Short: "A basic boring function",
	Long: `This calls the boring function without a channel`,
	Run: func(cmd *cobra.Command, args []string) {
		random, err := cmd.Flags().GetBool("random")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		msg, err := cmd.Flags().GetString("msg")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		boring(msg, random)
	},
}

func boring(msg string, random bool) {
	for i :=0; ; i++ {
		fmt.Printf("%s %d\n", msg, i)
		if random {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		} else {
			time.Sleep(time.Second)
		}
	}
}

func init() {
	boringCmd.AddCommand(withoutChannelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// withoutChannelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// withoutChannelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	withoutChannelCmd.Flags().String("msg", "Kaushik", "Message to print")

	withoutChannelCmd.Flags().Bool("random", false, "Should use randomness")
}
