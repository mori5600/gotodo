package todo

import (
	"time"

	"github.com/mori5600/gotodo/common"
)

type TodoApplicationService interface {
	Create(description string, dueDate time.Time) (TodoDTO, error)
	List() ([]TodoDTO, error)
	Update(id int, description string, status TodoStatus, dueDate time.Time) (TodoDTO, error)
}

type TodoApplicationServiceImpl struct {
	repo TodoRepository
}

func NewTodoApplicationService(repo TodoRepository) TodoApplicationService {
	return &TodoApplicationServiceImpl{
		repo: repo,
	}
}

func (s *TodoApplicationServiceImpl) Create(description string, dueDate time.Time) (TodoDTO, error) {
	todoCreate, err := NewTodoCreate(description, dueDate)
	if err != nil {
		return TodoDTO{}, err
	}

	// Call the repository to create the todo
	todo, err := s.repo.CreateTodo(todoCreate)
	if err != nil {
		return TodoDTO{}, err
	}

	// Convert the created todo to DTO
	todoDTO := TodoDTO{
		ID:          todo.ID,
		Description: todo.Description,
		Status:      StatusToString(todo.Status),
		DueDate:     common.TimeToString(todo.DueDate),
	}

	return todoDTO, nil
}

func (s *TodoApplicationServiceImpl) List() ([]TodoDTO, error) {
	todos, err := s.repo.GetAllTodos()
	if err != nil {
		return nil, err
	}

	var todoDTOs []TodoDTO
	for _, todo := range todos {
		todoDTO := TodoDTO{
			ID:          todo.ID,
			Description: todo.Description,
			Status:      StatusToString(todo.Status),
			DueDate:     common.TimeToString(todo.DueDate),
		}
		todoDTOs = append(todoDTOs, todoDTO)
	}

	return todoDTOs, nil
}

func (s *TodoApplicationServiceImpl) Update(id int, description string, status TodoStatus, dueDate time.Time) (TodoDTO, error) {
	todoUpdate, err := NewTodoUpdate(id, description, status, dueDate)
	if err != nil {
		return TodoDTO{}, err
	}
	// Call the repository to update the todo
	updatedTodo, err := s.repo.UpdateTodo(todoUpdate)
	if err != nil {
		return TodoDTO{}, err
	}
	return TodoDTO{
		ID:          updatedTodo.ID,
		Description: updatedTodo.Description,
		Status:      StatusToString(updatedTodo.Status),
		DueDate:     common.TimeToString(updatedTodo.DueDate),
	}, nil
}
