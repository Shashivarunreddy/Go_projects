package cmd

import (
	"fmt"

	"github.com/Shashivarunreddy/Go_projects/todo_cli/tasks"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task Description]",
	Short: "Add a task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		if err := tasks.AddTask("tasks.csv", description); err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task added Successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
