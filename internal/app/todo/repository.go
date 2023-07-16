package todo

import "github.com/ArMo-Team/GoLangPractices_todoList/internal/app/domain"

type TodoRepository interface {
	GetByID(id string) (*domain.Todo, error)
	GetAll() ([]*domain.Todo, error)
	Create(todo *domain.Todo) error
	Update(todo *domain.Todo) error
	Delete(id string) error
}
