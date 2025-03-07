/*
Copyright © 2025 Ben Fleuty <github.com/benfleuty>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		fmt.Println(args)
		parseDesc()
	},
}

func parseDesc() {
	if description == "" {
		fmt.Println("You need to provide a description for your task!")
		return
	}

	fmt.Println("Added task: \" " + description + "\".")
}

var description string

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&description, "description", "d", "", "A description of your task.")
}
