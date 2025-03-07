/*
Copyright © 2025 Ben Fleuty <github.com/benfleuty>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		for i, task := range tasks {
			fmt.Println(strconv.Itoa(i+1) + ": " + task)
		}
	},
}

var tasks []string

func init() {
	rootCmd.AddCommand(listCmd)
	tasks = []string{"clean kitchen", "make bed", "learn go"}
}
