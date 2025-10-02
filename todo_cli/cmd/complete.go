package cmd

import (
	"fmt"
	"strconv"

	"github.com/Shashivarunreddy/Go_projects/todo_cli/tasks"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [task Id]",
	Short: "marks tasks as completed",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Task Id must be number")
			return
		}

		if err := tasks.CompleteTask("tasks.csv", id); err != nil {
			fmt.Errorf("Error : ", err)
			return
		}

		fmt.Println("Task marked as completed")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
