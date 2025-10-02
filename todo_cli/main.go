package main

import (
	"fmt"
	"log"

	"github.com/Shashivarunreddy/Go_projects/todo_cli/tasks"
)

func main() {

	tasksList, err := tasks.LoadTasks("data/tasks.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("taskes loaded in csv")
	for _, t := range tasksList {
		fmt.Println(t.ID, t.Description, t.CreatedAt, t.Completed)
	}

}
