/*
Copyright © 2025 Ben Fleuty <github.com/benfleuty>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Sets a task's done status to true",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("complete called")

		id := convertToId(args[0])
		outputTask(id)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

func convertToId(idString string) int {
	id, err := strconv.Atoi(idString)
	if err != nil {
		// fmt.Printf("Error: Invalid ID < %s >", idString)
		log.Fatal(err)
	}
	validateId(id)
	return id
}

func validateId(id int) {
	idx := id - 1
	if idx < 0 || idx >= len(tasks) {
		msg := fmt.Sprintf("Error: Task not found with ID < %d >", id)
		fmt.Println(msg)
		log.Fatal(msg)
	}
}

func truncateTaskDesc(desc string, limit int) string {
	if len(desc)+3 > limit {
		return desc[:limit] + "..."
	}
	return desc
}

func outputTask(id int) {
	idx := id - 1
	task := tasks[idx]
	task.Done = true
	tasks[idx] = task
	fmt.Printf("Marked task #%d \"%s\" as completed:\n", id, truncateTaskDesc(task.Description, 20))
}
