package todo

import (
	"fmt"
	"time"
)

type Todo struct {
	ID          int
	Description string
	Status      int
	DueDate     time.Time
}

func NewTodo(id int, description string, dueDate time.Time) *Todo {
	return &Todo{
		ID:          id,
		Description: description,
		Status:      0,
		DueDate:     dueDate,
	}
}

func (t *Todo) String() string {
	return fmt.Sprintf(
		"Todo[ID=%d, Description=%s, Status=%d, DueDate=%s]",
		t.ID,
		t.Description,
		t.Status,
		t.DueDate.Format("2006-01-02 15:04:05"),
	)
}
