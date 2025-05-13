package todo

import "time"

type TodoController interface {
	Create(description string, dueDate time.Time) (int, error)
	List() ([]Todo, error)
}

type TodoControllerImpl struct {
	createTodoUseCase CreateTodoUseCase
	listTodosUseCase  ListTodosUseCase
}

func NewTodoController(createTodoUseCase CreateTodoUseCase, listTodoUseCase ListTodosUseCase) TodoController {
	return &TodoControllerImpl{
		createTodoUseCase: createTodoUseCase,
		listTodosUseCase:  listTodoUseCase,
	}
}

func (c *TodoControllerImpl) Create(description string, dueDate time.Time) (int, error) {
	todo, err := NewTodoCreate(description, dueDate)
	if err != nil {
		return 0, err
	}

	// Call the use case to create the todo
	id, err := c.createTodoUseCase.Execute(todo)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *TodoControllerImpl) List() ([]Todo, error) {
	todos, err := c.listTodosUseCase.Execute()
	if err != nil {
		return nil, err
	}
	return todos, nil
}
