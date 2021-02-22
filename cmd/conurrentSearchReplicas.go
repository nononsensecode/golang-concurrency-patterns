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

// conurrentSearchReplicasCmd represents the conurrentSearchReplicas command
var conurrentSearchReplicasCmd = &cobra.Command{
	Use:   "conurrentSearchReplicas",
	Short: "Does concurrent search using replicas",
	Long: `This function will do a concurrent search in all types of search (web, image, video) by
	spawning 2 instances of each search and by selecting the first result from these instances`,
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())
		start := time.Now()
		results := functions.GoogleWithReplicas("golang")
		elapsed := time.Since(start)
		fmt.Printf("Results: %v\n", results)
		fmt.Printf("Time used: %v\n", elapsed)
	},
}

func init() {
	rootCmd.AddCommand(conurrentSearchReplicasCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// conurrentSearchReplicasCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// conurrentSearchReplicasCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
