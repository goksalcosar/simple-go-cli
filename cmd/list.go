package cmd

import (
	"go-cli/helper"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list", // CLI command
	Short: "short description", // CLI short description
	Long:  `long description`, // CLI long description
	Run: func(cmd *cobra.Command, args []string) { // CLI command logic
		helper.ReadTodoItems("todo-list.json")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
