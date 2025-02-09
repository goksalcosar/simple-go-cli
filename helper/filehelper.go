package helper

import (
	"encoding/json"
	"fmt"
	"go-cli/schema"
	"os"
)

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func WriteTodoItem(fileName  string, todo schema.TodoItem) error {
	dir := "data"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create data directory: %w", err)
		}
	}

	filePath := fmt.Sprintf("%s/%s", dir, fileName)

	var todoItems []schema.TodoItem

	if FileExists(filePath) {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("unable to read file: %w", err)
		}

		err = json.Unmarshal(data, &todoItems)
		if err != nil {
			return fmt.Errorf("json unmarshalling error: %w", err)
		}
	} else {
		todoItems = []schema.TodoItem{}
	}

	todoItems = append(todoItems, todo)

	updatedData, err := json.MarshalIndent(todoItems, "", "  ")
	if err != nil {
		return fmt.Errorf("json marshalling error: %w", err)
	}

	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("unable to write to file: %w", err)
	}

	return nil
}

func ReadTodoItems(fileName string) ([]schema.TodoItem, error) {
	dir := "data"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("data directory does not exist, file cannot be found: %s/%s", dir, fileName)
	}

	filePath := fmt.Sprintf("%s/%s", dir, fileName)

	if !FileExists(filePath) {
		return nil, fmt.Errorf("file does not exist: %s", filePath)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %w", err)
	}

	var todoItems []schema.TodoItem
	err = json.Unmarshal(data, &todoItems)
	if err != nil {
		return nil, fmt.Errorf("json unmarshalling error: %w", err)
	}

	for _, todo := range todoItems {
		fmt.Printf("ID: %s\n", todo.ID)
		fmt.Printf("Title: %s\n", todo.Title)
		fmt.Printf("Description: %s\n", todo.Description)
		fmt.Printf("Date: %s\n", todo.Date)
		fmt.Printf("Complete: %v\n", todo.Complete)
		fmt.Println("-------------")
	}

	return todoItems, nil
}

func MarkTodoComplete(fileName, todoID string) error {
	dir := "data"
	filePath := fmt.Sprintf("%s/%s", dir, fileName)

	if !FileExists(filePath) {
		return fmt.Errorf("file does not exist: %s", filePath)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("unable to read file: %w", err)
	}

	var todoItems []schema.TodoItem
	err = json.Unmarshal(data, &todoItems)
	if err != nil {
		return fmt.Errorf("json unmarshalling error: %w", err)
	}

	var itemFound bool
	for i, todo := range todoItems {
		if todo.ID == todoID {
			todoItems[i].Complete = true
			itemFound = true
			// Print updated item
			fmt.Println("Todo item updated:")
			fmt.Printf("ID: %s\n", todoItems[i].ID)
			fmt.Printf("Title: %s\n", todoItems[i].Title)
			fmt.Printf("Description: %s\n", todoItems[i].Description)
			fmt.Printf("Date: %s\n", todoItems[i].Date)
			fmt.Printf("Complete: %v\n", todoItems[i].Complete)
			break
		}
	}

	if !itemFound {
		return fmt.Errorf("no todo item found with ID: %s", todoID)
	}

	updatedData, err := json.MarshalIndent(todoItems, "", "  ")
	if err != nil {
		return fmt.Errorf("json marshalling error: %w", err)
	}

	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("unable to write to file: %w", err)
	}

	return nil
}