/*
Copyright © 2025 Ben Fleuty <github.com/benfleuty>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/benfleuty/goprojects/todoapp/data"
	"github.com/benfleuty/goprojects/todoapp/model"
	"github.com/spf13/cobra"
)

var (
	showAllTasks bool
	tasks        []model.Task
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists your tasks.",
	Long: `Lists your tasks.
		Only shows defaults tasks by default.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var b strings.Builder
		fmt.Fprintf(&b, "%s\t%-20s\t%-12s\t%s\n", "ID", "Description", "Created", "Done?")
		for i, task := range tasks {
			if task.Done && !showAllTasks {
				continue
			}
			fmt.Fprintf(&b, "%d\t%-20s\t%d\t%t\n", i+1, task.Description, task.Created, task.Done)
		}
		fmt.Println(b.String())
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&showAllTasks, "all", "a", false, "Shows all tasks.")

	var taskReader data.TaskReader = &data.CSVReader{FilePath: "db.csv"}
	tasks = taskReader.GetAll()
}
