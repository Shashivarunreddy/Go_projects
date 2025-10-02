package tasks

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	Completed   bool
}

func LoadTasks(filename string) ([]Task, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return []Task{}, nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tasks []Task
	for _, record := range records {
		if len(record) < 4 {
			continue
		}

		id, _ := strconv.Atoi(record[0])
		decs := record[1]
		createdAt, _ := time.Parse(time.RFC3339, record[2])
		completed, _ := strconv.ParseBool(record[3])

		tasks = append(tasks, Task{
			ID:          id,
			Description: decs,
			CreatedAt:   createdAt,
			Completed:   completed,
		})

	}

	return tasks, nil

}
