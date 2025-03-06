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
		parseName()
	},
}

func parseName() {
	if name == "" {
		fmt.Println("No name set!")
		return
	}

	fmt.Println("Name was set to: " + name)
}

var name string

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Help message for name")
}
