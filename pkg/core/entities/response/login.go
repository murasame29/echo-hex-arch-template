package response

type Login struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}
