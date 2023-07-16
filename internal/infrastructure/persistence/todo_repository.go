package persistence

import (
	"database/sql"
	"errors"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/app/domain"
	_ "github.com/go-sql-driver/mysql"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) GetByID(id string) (*domain.Todo, error) {
	query := "SELECT id, title, description, completed FROM todos WHERE id = ?"

	row := r.db.QueryRow(query, id)

	todo := &domain.Todo{}
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No todo found with the given ID
		}
		return nil, err
	}

	return todo, nil
}

func (r *TodoRepository) GetAll() ([]*domain.Todo, error) {
	query := "SELECT id, title, description, completed FROM todos"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*domain.Todo{}
	for rows.Next() {
		todo := &domain.Todo{}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.Description)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *TodoRepository) Create(todo *domain.Todo) error {
	query := "INSERT INTO todos (title, description) VALUES (?, ?)"

	_, err := r.db.Exec(query, todo.Title, todo.Description)
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) Update(todo *domain.Todo) error {
	query := "UPDATE todos SET title = ?, description = ?, completed = ? WHERE id = ?"

	_, err := r.db.Exec(query, todo.Title, todo.Description, todo.Completed, todo.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) Delete(id string) error {
	query := "DELETE FROM todos WHERE id = ?"

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
