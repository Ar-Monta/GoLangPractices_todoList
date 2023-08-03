package user

import (
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/app/domain"
	"time"
)

type UserService interface {
	GetUsers() ([]*domain.User, error)
	GetUserByID(id string) (*domain.User, error)
	CreateUser(user *domain.User) error
	UpdateUser(user *domain.User) error
	UpdateUserVerifiedAt(userID string, verifiedAt time.Time) error
	DeleteUser(id string) error
}

type userService struct {
	Repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		Repository: repository,
	}
}

func (s *userService) GetUsers() ([]*domain.User, error) {
	users, err := s.Repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) GetUserByID(id string) (*domain.User, error) {
	user, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil // User not found
	}
	return user, nil
}

func (s *userService) CreateUser(user *domain.User) error {
	err := s.Repository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) UpdateUser(user *domain.User) error {
	err := s.Repository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) UpdateUserVerifiedAt(userID string, verifiedAt time.Time) error {
	// Retrieve the user by ID from the repository
	user, err := s.Repository.GetByID(userID)
	if err != nil {
		return err
	}

	// Update the verified_at status
	user.VerifiedAt = verifiedAt

	// Update the user in the repository
	err = s.Repository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) DeleteUser(id string) error {
	err := s.Repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
