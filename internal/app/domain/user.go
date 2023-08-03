package domain

import "time"

type User struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	VerifiedAt time.Time `json:"verified_at"`
	createdAt  time.Time
	updatedAt  time.Time
}
