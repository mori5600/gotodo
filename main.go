package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mori5600/gotodo/common"
	"github.com/mori5600/gotodo/logging"
	"github.com/mori5600/gotodo/todo"
)

type TodoStatus int

const (
	NotStarted TodoStatus = iota
	InProgress
	Done
)

func (s TodoStatus) String() string {
	switch s {
	case NotStarted:
		return "Not Started"
	case InProgress:
		return "In Progress"
	case Done:
		return "Done"
	default:
		return "Unknown"
	}
}

func convertNumToStatus(num int) TodoStatus {
	switch num {
	case 0:
		return NotStarted
	case 1:
		return InProgress
	case 2:
		return Done
	default:
		return NotStarted
	}
}

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
	var todos []todo.Todo
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Select an option:")
		fmt.Println("0. Exit")
		fmt.Println("1. List Todos")
		fmt.Println("2. Add Todo")
		fmt.Println("3. Update Todo")
		fmt.Print("Enter your choice: ")

		if !scanner.Scan() {
			logErrorReadingInput(scanner)
			continue
		}

		choice := scanner.Text()

		switch choice {
		case "0":
			fmt.Println("Exiting...")
			return
		case "1":
			fmt.Println("=== Todo List ===")
			for _, t := range todos {
				msg := fmt.Sprintf(
					"Todo[ID=%d, Description=%s, Status=%s, DueDate=%s]",
					t.ID,
					t.Description,
					convertNumToStatus(t.Status),
					t.DueDate.Format(common.TIME_FORMAT),
				)
				fmt.Println(msg)
			}
			fmt.Println("==================")
			continue
		case "2":
			fmt.Printf("Enter your todo description: ")
			if err := logErrorReadingInput(scanner); err != nil {
				break
			}
			description := scanner.Text()

			todoID := len(todos) + 1

			fmt.Printf("Enter due date (YYYY-MM-DD HH:MM): ")
			if err := logErrorReadingInput(scanner); err != nil {
				break
			}
			dueDateInput := scanner.Text()
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
		case "3":
			fmt.Printf("Enter the ID of the todo to update: ")
			if err := logErrorReadingInput(scanner); err != nil {
				break
			}
			idInput := scanner.Text()
			id, err := strconv.Atoi(idInput)
			if err != nil {
				fmt.Println("Error parsing ID:", err)
				continue
			}
			if id < 1 || id > len(todos) {
				fmt.Println("Invalid ID. Please try again.")
				continue
			}
			todoItem := todos[id-1]
			fmt.Printf("Current description: %s\n", todoItem.Description)
			fmt.Printf("Enter new description (leave blank to keep current): ")
			if err := logErrorReadingInput(scanner); err != nil {
				break
			}
			newDescription := scanner.Text()
			if newDescription != "" {
				todoItem.Description = newDescription
			}
			fmt.Printf("Current status: %s\n", convertNumToStatus(todoItem.Status))
			fmt.Printf("Enter new status (0: Not Started, 1: In Progress, 2: Done): ")
			if err := logErrorReadingInput(scanner); err != nil {
				break
			}
			statusInput := scanner.Text()
			status, err := strconv.Atoi(statusInput)
			if err != nil {
				fmt.Println("Error parsing status:", err)
				continue
			}
			if status < 0 || status > 2 {
				fmt.Println("Invalid status. Please try again.")
				continue
			}
			todoItem.Status = status
			fmt.Printf("Current due date: %s\n", todoItem.DueDate.Format(common.TIME_FORMAT))
			fmt.Printf("Enter new due date (leave blank to keep current): ")
			if err := logErrorReadingInput(scanner); err != nil {
				break
			}
			newDueDateInput := scanner.Text()
			if newDueDateInput != "" {
				newDueDate, err := time.Parse(common.TIME_FORMAT, newDueDateInput)

				if err != nil {
					fmt.Println("Error parsing due date:", err)
					continue
				}
				todoItem.DueDate = newDueDate
			}
			todos[id-1] = todoItem
			fmt.Println("Todo updated:", todoItem)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
