package todo

type CreateTodoUseCase interface {
	Execute(todo TodoCreate) (int, error)
}

type CreateTodoUseCaseImpl struct {
	repo TodoRepository
}

func NewCreateTodoUseCase(repo TodoRepository) CreateTodoUseCase {
	return &CreateTodoUseCaseImpl{
		repo: repo,
	}
}

func (uc *CreateTodoUseCaseImpl) Execute(todo TodoCreate) (int, error) {
	id, err := uc.repo.CreateTodo(todo)
	if err != nil {
		return 0, err
	}
	return id, nil
}
