package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mori5600/gotodo/common"
	"github.com/mori5600/gotodo/db"
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
	// Initialize the logger
	logger := logging.GetLogger()
	logger.Info("Starting Todo application")

	// Initialize the database
	conn, err := db.InitDB()
	if err != nil {
		logger.Error("Error initializing database", "error", err)
		return
	}
	defer conn.Close()

	repo := todo.NewSQLiteTodoRepository(conn)
	todoApplicationService := todo.NewTodoApplicationService(repo)

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
			todos, err := todoApplicationService.List()
			if err != nil {
				fmt.Println("Error listing todos:", err)
				continue
			}
			for _, t := range todos {
				fmt.Println(t.String())
			}
			fmt.Println("==================")
			continue
		case "2":
			fmt.Printf("Enter your todo description: ")
			if err := logErrorReadingInput(scanner); err != nil {
				break
			}
			description := scanner.Text()

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

			dto, err := todoApplicationService.Create(description, dueDate)
			if err != nil {
				fmt.Println("Error creating todo:", err)
				continue
			}

			fmt.Println("Todo added:", dto.String())
			continue
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

			todos, err := todoApplicationService.List()
			if err != nil {
				fmt.Println("Error listing todos:", err)
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
			// fmt.Printf("Current status: %s\n", convertNumToStatus(todoItem.Status))
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

			fmt.Printf("Current due date: %s\n", todoItem.DueDate)
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
				todoItem.DueDate = newDueDate.Format(common.TIME_FORMAT)
			}
			todoStatus, err := todo.IntToStatus(status)
			if err != nil {
				fmt.Println("Error converting status:", err)
				continue
			}

			d, err := time.Parse(common.TIME_FORMAT, todoItem.DueDate)
			if err != nil {
				fmt.Println("Error parsing due date:", err)
				continue
			}
			updatedDto, err := todoApplicationService.Update(todoItem.ID, todoItem.Description, todoStatus, d)
			if err != nil {
				fmt.Println("Error updating todo:", err)
				continue
			}
			fmt.Println("Todo updated:", updatedDto.String())
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
