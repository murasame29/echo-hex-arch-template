package token

import "time"

type Maker interface {
	CreateToken(userID string, email string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
