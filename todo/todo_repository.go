package todo

import (
	"database/sql"
	"fmt"
)

type TodoRepository interface {
	GetAllTodos() ([]Todo, error)
	GetTodoByID(id int) (Todo, error)
	CreateTodo(todo TodoCreate) (int, error)
	UpdateTodo(todo Todo) error
	DeleteTodo(id int) error
	GetTodosByStatus(status int) ([]Todo, error)
}

type SQLiteTodoRepository struct {
	db *sql.DB
}

func NewSQLiteTodoRepository(db *sql.DB) TodoRepository {
	return &SQLiteTodoRepository{db: db}
}

func (r *SQLiteTodoRepository) GetAllTodos() ([]Todo, error) {
	rows, err := r.db.Query("SELECT id, description, status, due_date FROM todos WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Description, &todo.Status, &todo.DueDate); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *SQLiteTodoRepository) GetTodoByID(id int) (Todo, error) {
	row := r.db.QueryRow("SELECT id, description, status, due_date FROM todos WHERE id = ? AND deleted_at IS NULL", id)

	var todo Todo
	if err := row.Scan(&todo.ID, &todo.Description, &todo.Status, &todo.DueDate); err != nil {
		if err == sql.ErrNoRows {
			return Todo{}, fmt.Errorf("todo with ID %d not found", id)
		}
		return Todo{}, err
	}

	return todo, nil
}

func (r *SQLiteTodoRepository) CreateTodo(todo TodoCreate) (int, error) {
	const initialStatus = 0

	result, err := r.db.Exec(
		"INSERT INTO todos (description, status, due_date) VALUES (?, ?, ?)",
		todo.Description, initialStatus, todo.DueDate,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *SQLiteTodoRepository) UpdateTodo(todo Todo) error {
	_, err := r.db.Exec("UPDATE todos SET description = ?, status = ?, due_date = ? WHERE id = ?", todo.Description, todo.Status, todo.DueDate, todo.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteTodoRepository) DeleteTodo(id int) error {
	_, err := r.db.Exec("UPDATE todos SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteTodoRepository) GetTodosByStatus(status int) ([]Todo, error) {
	rows, err := r.db.Query("SELECT id, description, status, due_date FROM todos WHERE status = ? AND deleted_at IS NULL", status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Description, &todo.Status, &todo.DueDate); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
