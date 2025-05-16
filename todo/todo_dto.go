package todo

import "fmt"

type TodoDTO struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
}

func NewTodoDTO(id int, description, status, dueDate string) TodoDTO {
	return TodoDTO{
		ID:          id,
		Description: description,
		Status:      status,
		DueDate:     dueDate,
	}
}

func (t TodoDTO) String() string {
	return fmt.Sprintf(
		"TodoDTO[ID=%d, Description=%s, Status=%s, DueDate=%s]",
		t.ID,
		t.Description,
		t.Status,
		t.DueDate,
	)
}
