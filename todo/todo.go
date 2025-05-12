package todo

import (
	"fmt"
	"time"

	"github.com/mori5600/gotodo/common"
)

type Todo struct {
	ID          int
	Description string
	Status      int
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
