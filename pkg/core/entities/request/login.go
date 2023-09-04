package request

type Login struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}
