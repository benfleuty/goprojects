package data

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/benfleuty/goprojects/todoapp/model"
)

type CSVReader struct {
	FilePath string
}

func (r *CSVReader) GetAll() []model.Task {
	records, err := readCsv(r)
	if err != nil {
		log.Fatalf("There was an error reading the CSV file %s!\n%v", r.FilePath, err)
	}

	return parseTasks(records)
}

func (r *CSVReader) WriteTask(desc *string) (model.Task, error) {
	id, err := getNextId(r)
	if err != nil {
		log.Fatalf("Error generating an ID for \"%v\": %v", desc, err)
	}

	now := time.Now().Unix()
	isDone := false

	task := model.Task{
		ID:          id,
		Description: *desc,
		Created:     int(now),
		Done:        isDone,
	}

	f, err := os.Open(r.FilePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()

	record := []string{
		strconv.Itoa(task.ID),
		task.Description,
		strconv.Itoa(task.Created),
		strconv.FormatBool(task.Done),
	}

	fmt.Printf("Record: %v", record)

	w := csv.NewWriter(f)
	err = w.Write(record)
	if err != nil {
		log.Fatalf("Error writing task %v to file '%s': %v", record, r.FilePath, err)
	}

	if err != nil {
		fmt.Printf("%v\n", task)
		log.Fatalf("Error writing to file: %v", err)
	}

	return task, nil
}

func getNextId(r *CSVReader) (int, error) {
	records, err := readCsv(r)
	if err != nil {
		return -1, err
	}
	biggestId := 0
	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}
		if id > biggestId {
			biggestId = id
		}
	}
	return biggestId + 1, nil
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

func readCsv(r *CSVReader) ([][]string, error) {
	f, err := os.Open(r.FilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
