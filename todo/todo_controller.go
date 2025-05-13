package todo

type TodoController interface {
	CreateTodo(todo TodoCreate) (int, error)
}

type TodoControllerImpl struct {
	uc CreateTodoUseCase
}

func NewTodoController(uc CreateTodoUseCase) TodoController {
	return &TodoControllerImpl{
		uc: uc,
	}
}

func (c *TodoControllerImpl) CreateTodo(todo TodoCreate) (int, error) {
	id, err := c.uc.Execute(todo)
	if err != nil {
		return 0, err
	}
	return id, nil
}
