package domain

import "time"

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	createdAt   time.Time
	updatedAt   time.Time
}
