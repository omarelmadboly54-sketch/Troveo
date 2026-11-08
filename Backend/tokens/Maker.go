package tokens

import "time"

type Maker interface {
	CreateToken(UserID int64, Role string, duration time.Duration)(string,error)
	VerifyToken(tokenString string) (*Payload, error)
}
