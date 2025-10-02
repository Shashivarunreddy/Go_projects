package main

import (
	"fmt"
	"github.com/Shashivarunreddy/Go_projects/todo_cli/tasks"
)

func main() {

	filename := "tasks.csv"

	// Add a new task
	err := tasks.AddTask(filename, "Learn Go CLI project")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Task added successfully!")

	// Reload tasks to verify
	allTasks, _ := tasks.LoadTasks(filename)
	fmt.Println("All Tasks:", allTasks)

	tasks.ListTasks(filename, true)

	tasks.ListTasks(filename, false)
	err1 := tasks.CompleteTask(filename, 4)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Task marked as completed!")
	}

	// Delete a task
	err2 := tasks.DeleteTask(filename, 5)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Task deleted!")
	}

}
