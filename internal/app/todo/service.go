package todo

import "github.com/ArMo-Team/GoLangPractices_todoList/internal/app/domain"

type TodoService interface {
	GetTodos() ([]*domain.Todo, error)
	GetTodoByID(id string) (*domain.Todo, error)
	CreateTodo(todo *domain.Todo) error
	UpdateTodo(todo *domain.Todo) error
	UpdateTodoCompleted(todoID string, completed bool) error
	DeleteTodo(id string) error
}

type todoService struct {
	Repository TodoRepository
}

func NewTodoService(repository TodoRepository) TodoService {
	return &todoService{
		Repository: repository,
	}
}

func (s *todoService) GetTodos() ([]*domain.Todo, error) {
	todos, err := s.Repository.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *todoService) GetTodoByID(id string) (*domain.Todo, error) {
	todo, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, nil // Todo not found
	}
	return todo, nil
}

func (s *todoService) CreateTodo(todo *domain.Todo) error {
	err := s.Repository.Create(todo)
	if err != nil {
		return err
	}
	return nil
}

func (s *todoService) UpdateTodo(todo *domain.Todo) error {
	err := s.Repository.Update(todo)
	if err != nil {
		return err
	}
	return nil
}

func (s *todoService) UpdateTodoCompleted(todoID string, completed bool) error {
	// Retrieve the todo by ID from the repository
	todo, err := s.Repository.GetByID(todoID)
	if err != nil {
		return err
	}

	// Update the completed status
	todo.Completed = completed

	// Update the todo in the repository
	err = s.Repository.Update(todo)
	if err != nil {
		return err
	}
	return nil
}

func (s *todoService) DeleteTodo(id string) error {
	err := s.Repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
