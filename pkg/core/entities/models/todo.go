package models

import "time"

type Todo struct {
	TodoID      string    `json:"todo_id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsComplete  bool      `json:"is_complete"`
}
