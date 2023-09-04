package token

import "time"

type Maker interface {
	CreateToken(userID string, email string, duration time.Duration) (string, error)
	Verifytoken(token string) (*Payload, error)
}
