package token

import "time"

type Maker interface {
	CreateToken(userPayload UserPayload, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
