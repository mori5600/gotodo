package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mori5600/gotodo/common"
	"github.com/mori5600/gotodo/logging"
	"github.com/mori5600/gotodo/todo"
)

func logErrorReadingInput(sccaner *bufio.Scanner) error {
	logger := logging.GetLogger()
	if !sccaner.Scan() {
		err := sccaner.Err()
		logger.Error("Error reading input", "error", err)
		return err
	}
	return nil
}

func main() {
	var todos []*todo.Todo

	sccaner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Select an option:")
		fmt.Println("0. Exit")
		fmt.Println("1. List Todos")
		fmt.Println("2. Add Todo")
		fmt.Println("3. Update Todo")
		fmt.Print("Enter your choice: ")

		if err := logErrorReadingInput(sccaner); err != nil {
			break
		}

		fmt.Printf("Enter your todo description: ")
		if err := logErrorReadingInput(sccaner); err != nil {
			break
		}

		description := sccaner.Text()
		todoID := len(todos) + 1
		fmt.Printf("Enter due date (YYYY-MM-DD HH:MM): ")
		if err := logErrorReadingInput(sccaner); err != nil {
			break
		}

		dueDateInput := sccaner.Text()
		dueDate, err := time.Parse(common.TIME_FORMAT, dueDateInput)
		if err != nil {
			fmt.Println("Error parsing due date:", err)
			continue
		}
		todoItem := todo.NewTodo(todoID, description, dueDate)
		todos = append(todos, todoItem)
		fmt.Println("Todo added:", todoItem)
		fmt.Println("Todos:")
		for _, t := range todos {
			fmt.Println(t.String())
		}
	}
}
