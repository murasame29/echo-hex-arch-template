package response

import "time"

type ListTodo struct {
	TodoID     string    `json:"todo_id"`
	Title      string    `json:"title"`
	CreatedAt  time.Time `json:"created_at"`
	IsComplete bool      `json:"is_complete"`
}

type Todo struct {
	TodoID      string    `json:"todo_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsComplete  bool      `json:"is_complete"`
}
