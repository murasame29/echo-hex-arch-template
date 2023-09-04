package request

type TodoURI struct {
	TodoID string `param:"todo_id"`
}

type ListTodo struct {
	PageSize   int    `query:"page_size"`
	PageID     int    `query:"page_id"`
	IsComplete bool   `query:"is_complete"`
	UserID     string `query:"user_id"`
}

type CreateTodo struct {
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DeleteAllTodo struct {
	UserID string `query:"user_id"`
}

type UpdateTodo struct {
	TodoURI
	CreateTodo
	IsComplete bool `json:"is_complete"`
}
