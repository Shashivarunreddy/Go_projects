package cmd

import (
	"fmt"
	"strconv"

	"github.com/Shashivarunreddy/Go_projects/todo_cli/tasks"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Task ID must be a number")
			return
		}
		if err := tasks.DeleteTask("tasks.csv", id); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Task deleted successfully!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
