package cmd

import (
	"fmt"
	"go-cli/helper"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add", // CLI command
	Short: "short description", // CLI short description
	Long:  `long description`, // CLI long description
	Run: func(cmd *cobra.Command, args []string) { // CLI command logic
		todo, err := helper.AddTodoFlags(cmd)

		if err != nil {
			fmt.Println("Error: ", err)

			return
		}

		helper.WriteTodoItem("todo-list.json", *todo)
	},
}
	
func init() {
	helper.BuildToAdd(addCmd)

	rootCmd.AddCommand(addCmd)
}
