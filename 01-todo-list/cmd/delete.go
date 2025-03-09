/*
Copyright © 2025 Ben Fleuty <github.com/benfleuty>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/benfleuty/goprojects/todoapp/model"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from your TODO list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		if !taskIdExists(id) {
			log.Fatalf("%d is not a valid task ID!", id)
		}

		task := tasks[id-1]
		deleteTask(tasks, id-1)
		fmt.Printf("Deleted task #%d: %s\n", id, task.Description)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func taskIdExists(id int) bool {
	return id >= 1 && id <= len(tasks)
}

func deleteTask(taskList []model.Task, idx int) []model.Task {
	return append(taskList[idx:], taskList[:idx+1]...)
}
