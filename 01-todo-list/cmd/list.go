/*
Copyright © 2025 Ben Fleuty <github.com/benfleuty>
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type Task struct {
	Description string
	Created     int64
	Done        bool
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		var builder strings.Builder
		out := fmt.Sprintf("%s\t%-20s\t%-12s\t%s\n", "ID", "Description", "Created", "Done?")
		builder.WriteString(out)
		for i, task := range tasks {
			out := fmt.Sprintf("%d\t%-20s\t%d\t%t\n", i+1, task.Description, task.Created, task.Done)
			builder.WriteString(out)
		}
		fmt.Println(builder.String())
	},
}

var tasks []Task

func init() {
	rootCmd.AddCommand(listCmd)
	tasks = append(tasks,
		Task{Description: "Clean the kitchen", Created: time.Now().Unix(), Done: false},
		Task{Description: "Make the bed", Created: time.Now().Unix() - (3600 * 6), Done: true},
		Task{Description: "Learn more Go!", Created: time.Now().Unix() - (3600 * 3), Done: false},
		Task{Description: "Feed the Gopher", Created: time.Now().Unix() - 3600, Done: true},
	)
}
