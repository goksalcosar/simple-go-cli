package cmd

import (
	"fmt"
	"go-cli/helper"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete", // CLI command
	Short: "short description", // CLI short description
	Long:  `long description`, // CLI long description
	Run: func(cmd *cobra.Command, args []string) { // CLI command logic
		todo, err := helper.CompleteToFlags(cmd)

		if err != nil {
			fmt.Println("Error: ", err)

			return
		}

		helper.MarkTodoComplete("todo-list.json", todo.ID)
	},
}

func init() {
	helper.BuildToComplate(completeCmd)

	rootCmd.AddCommand(completeCmd)
}
