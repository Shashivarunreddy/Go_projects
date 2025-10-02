package cmd

import (
	"fmt"

	"github.com/Shashivarunreddy/Go_projects/todo_cli/tasks"
	"github.com/spf13/cobra"
)

var listAll bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if err := tasks.ListTasks("tasks.csv", listAll); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&listAll, "all", "a", false, "List all the tasks including completed")
}
