package todo

import (
	"fmt"
	"time"
)

type Todo struct {
	ID          int
	Description string
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewTodo(id int, description string) *Todo {
	return &Todo{
		ID:          id,
		Description: description,
		Status:      0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Todo) String() string {
	return fmt.Sprintf(
		"ID: %d, Description: %s, Status: %d, CreatedAt: %s, UpdatedAt: %s, DeletedAt: %s",
		t.ID,
		t.Description,
		t.Status,
		t.CreatedAt,
		t.UpdatedAt,
		t.DeletedAt,
	)
}
