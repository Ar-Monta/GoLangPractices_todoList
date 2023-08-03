package persistence

import (
	"database/sql"
	"errors"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/app/domain"
	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id string) (*domain.User, error) {
	query := "SELECT id, first_name, last_name FROM users WHERE id = ?"

	row := r.db.QueryRow(query, id)

	user := &domain.User{}
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No user found with the given ID
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetAll() ([]*domain.User, error) {
	query := "SELECT id, first_name, last_name FROM users"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*domain.User{}
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Create(user *domain.User) error {
	query := "INSERT INTO users (first_name, last_name) VALUES (?, ?)"

	_, err := r.db.Exec(query, user.FirstName, user.LastName)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(user *domain.User) error {
	query := "UPDATE users SET first_name = ?, last_name = ? WHERE id = ?"

	_, err := r.db.Exec(query, user.FirstName, user.LastName, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateVerify(user *domain.User) error {
	query := "UPDATE users SET verified_at = ? WHERE id = ?"

	_, err := r.db.Exec(query, user.VerifiedAt, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id string) error {
	query := "DELETE FROM users WHERE id = ?"

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
