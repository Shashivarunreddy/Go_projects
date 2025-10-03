package tasks

import (
	"encoding/csv"
	"fmt"
	"github.com/mergestat/timediff"
	"os"
	"strconv"
	"text/tabwriter"
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

func SaveTasks(filename string, tasks []Task) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, task := range tasks {
		record := []string{
			strconv.Itoa(task.ID),
			task.Description,
			task.CreatedAt.Format(time.RFC3339),
			strconv.FormatBool(task.Completed),
		}

		if err := writer.Write(record); err != nil {
			return err
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}

func AddTask(filename, description string) error {
	tasks, err := LoadTasks(filename)
	if err != nil {
		return err
	}
	newId := 1
	for _, t := range tasks {
		if t.ID >= newId {
			newId = t.ID + 1
		}
	}
	newTask := Task{
		ID:          newId,
		Description: description,
		CreatedAt:   time.Now(),
		Completed:   false,
	}

	tasks = append(tasks, newTask)

	return SaveTasks(filename, tasks)
}

func ListTasks(filename string, all bool) error {
	tasks, err := LoadTasks(filename)

	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("now tasks found")
		return nil
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	if all {
		fmt.Fprintln(writer, "ID\tTask\tCreated\tDone")
	} else {
		fmt.Fprintln(writer, "ID\tTask\tCreated")
	}

	for _, t := range tasks {
		if !all && t.Completed {
			continue
		}

		createdStr := timediff.TimeDiff(t.CreatedAt)

		if all {
			fmt.Fprintf(writer, "%d\t%s\t%s\t%v\n", t.ID, t.Description, createdStr, t.Completed)
		} else {
			fmt.Fprintf(writer, "%d\t%s\t%s\n", t.ID, t.Description, createdStr)
		}
	}
	writer.Flush()
	return nil
}

func CompleteTask(filename string, id int) error {
	tasks, err := LoadTasks(filename)

	if err != nil {
		return err
	}

	found := false

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		fmt.Println("The task with %d not found", id)
	}

	return SaveTasks(filename, tasks)
}

func DeleteTask(filename string, id int) error {
	tasks, err := LoadTasks(filename)

	if err != nil {
		return err
	}

	newTasks := []Task{}

	found := false

	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}

		newTasks = append(newTasks, t)
	}

	if !found {
		return fmt.Errorf(" task with id %d not found", id)
	}

	return SaveTasks(filename, newTasks)
}
