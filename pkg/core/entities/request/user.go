package request

type UserURI struct {
	UserID string `param:"user_id"`
}

type ListUser struct {
	PageID   int `query:"page_id"`
	PageSize int `query:"page_size"`
}

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	UserURI
	CreateUser
}
