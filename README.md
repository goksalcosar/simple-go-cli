# Todo CLI Commands

This Todo CLI application allows you to manage your Todo list directly from the terminal.

## Commands

### 1. Add a Todo Item

This command allows you to add a new Todo item. It requires a title, description, and due date.

```bash
go run main.go add --title "Buy Milk" --description "Get milk from the store" --date "2025-02-10"
```

### 2. List All Todo Items

This command displays all the Todo items saved in the data file.

```bash
go run main.go list
```

### 3. Mark Todo as Complete

This command marks a specific Todo item as "Complete" by providing the Todo item ID.

```bash
go run main.go complete --id "id"
```
