/*
Copyright © 2025 Ben Fleuty <github.com/benfleuty>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var description string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if args[0] == "" {
			return errors.New("description cannot be empty")
		}
		description = args[0]
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("add called with task: ")
		fmt.Println(description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
