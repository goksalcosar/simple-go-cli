package helper

import (
	"fmt"
	"go-cli/schema"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var id, title, description, date string

func BuildToAdd(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&title, "title", "t", "", "The title of the todo item")
	cmd.Flags().StringVarP(&description, "description", "d", "", "The description of the todo item")
	cmd.Flags().StringVarP(&date, "date", "D", "", "The due date for the todo item")

	return
}

func BuildToComplate(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&id, "id", "i", "", "The id of the todo item")
}

func AddTodoFlags(cmd *cobra.Command) (*schema.TodoItem, error) {
	if title == "" || description == "" || date == "" {
		return nil, fmt.Errorf("Error: All fields (title, description, and date) are required.")
	}

	todo := schema.TodoItem{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Date:        date,
		Complete:    false,
	}

	return &todo, nil
}

func CompleteToFlags(cmd *cobra.Command) (*schema.TodoItem, error) {
	if id == "" {
		return nil, fmt.Errorf("Error: Id field is required.")
	}

	todo := schema.TodoItem{
		ID:          id,
	}

	return &todo, nil
}
