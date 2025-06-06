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
		log.Fatalf("Error generating an ID for \"%v\": %v\n", desc, err)
	}

	done := false

	task := model.Task{
		ID:          id,
		Description: *desc,
		Created:     time.Now(),
		Done:        done,
	}

	f, err := os.OpenFile(r.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o640)
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer f.Close()

	idStr := strconv.Itoa(task.ID)
	descStr := task.Description
	createdStr := strconv.FormatInt(task.Created.Unix(), 10)
	doneStr := strconv.FormatBool(task.Done)
	record := []string{
		idStr,
		descStr,
		createdStr,
		doneStr,
	}

	fmt.Println("Record: ", record)

	w := csv.NewWriter(f)
	err = w.Write(record)
	if err != nil {
		log.Fatalf("Error writing task %v to file '%s': %v\n", record, r.FilePath, err)
	}

	w.Flush()

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
		createdUnixTimeStr := record[2]
		createdUnixTimeInt, err := strconv.ParseInt(createdUnixTimeStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		created := time.Unix(createdUnixTimeInt, 0)
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
