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
	"time"

	"github.com/spf13/cobra"
	"nononsensecode.com/go-concurrency-patterns/functions"
)

// conurrentSearchTimeoutCmd represents the conurrentSearchTimeout command
var conurrentSearchTimeoutCmd = &cobra.Command{
	Use:   "conurrentSearchTimeout",
	Short: "Concurrent search with timeout",
	Long: `This will accepts a string as a query and will do a search in 3 categories,
	 but will not wait every results to complete. It will exit after a certain time`,
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())
		start := time.Now()
		results := functions.GoogleSyncWithTimeout("golang")
		elapsed := time.Since(start)
		fmt.Printf("Results: %v\n", results)
		fmt.Printf("Time used: %v\n", elapsed)
	},
}

func init() {
	rootCmd.AddCommand(conurrentSearchTimeoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// conurrentSearchTimeoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// conurrentSearchTimeoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
