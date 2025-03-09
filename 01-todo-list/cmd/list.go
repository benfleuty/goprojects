/*
Copyright © 2025 Ben Fleuty <github.com/benfleuty>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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

	taskRecords, err := readCsv("db.csv")
	if err != nil {
		log.Fatal(err)
	}

	tasks = parseTasks(taskRecords)
}

func parseTasks(records [][]string) []model.Task {
	if len(records) == 0 {
		return []model.Task{}
	}

	var tasks []model.Task

	for _, record := range records {
		if len(record) < 4 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}
		description := record[1]
		created, err := strconv.Atoi(record[2])
		if err != nil {
			created = -1
		}
		done, err := strconv.ParseBool(record[3])
		if err != nil {
			done = false
		}

		var task model.Task
		task.ID = id
		task.Description = description
		task.Created = created
		task.Done = done
		tasks = append(tasks, task)
	}

	return tasks
}

func readCsv(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
