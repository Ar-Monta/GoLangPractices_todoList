package user

import "github.com/ArMo-Team/GoLangPractices_todoList/internal/app/domain"

type UserRepository interface {
	GetByID(id string) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	UpdateVerify(user *domain.User) error
	Delete(id string) error
}
