package todo

import (
	"fmt"
	"time"

	"github.com/mori5600/gotodo/common"
)

type Todo struct {
	ID          int
	Description string
	Status      TodoStatus
	DueDate     time.Time
}

func NewTodo(id int, description string, dueDate time.Time) Todo {
	return Todo{
		ID:          id,
		Description: description,
		Status:      0,
		DueDate:     dueDate,
	}
}

func (t Todo) String() string {
	return fmt.Sprintf(
		"Todo[ID=%d, Description=%s, Status=%d, DueDate=%s]",
		t.ID,
		t.Description,
		t.Status,
		t.DueDate.Format(common.TIME_FORMAT),
	)
}

type TodoCreate struct {
	Description string
	DueDate     time.Time
}

func NewTodoCreate(description string, dueDate time.Time) (TodoCreate, error) {
	if description == "" {
		return TodoCreate{}, fmt.Errorf("description cannot be empty")
	}
	return TodoCreate{
		Description: description,
		DueDate:     dueDate,
	}, nil
}

func (t TodoCreate) String() string {
	return fmt.Sprintf(
		"TodoCreate[Description=%s, DueDate=%s]",
		t.Description,
		t.DueDate.Format(common.TIME_FORMAT),
	)
}

type TodoUpdate struct {
	ID          int
	Description string
	Status      TodoStatus
	DueDate     time.Time
}

func NewTodoUpdate(id int, description string, status TodoStatus, dueDate time.Time) (TodoUpdate, error) {
	if description == "" {
		return TodoUpdate{}, fmt.Errorf("description cannot be empty")
	}
	return TodoUpdate{
		ID:          id,
		Description: description,
		Status:      status,
		DueDate:     dueDate,
	}, nil
}
