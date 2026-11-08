package tokens

import (
	"errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken=errors.New("token is expired")
)

type Payload struct {
	UserID			int64		`json:"user_id"`
	Role			string		`json:"role"`
	IssuedAt		time.Time	`json:"issued_at"`
	ExpiredAt		time.Time	`json:"expired_at"`
}

func NewPayloadToken(UserID int64,Role string,duration time.Duration)(*Payload,error){
	paylaod:=&Payload{
		UserID: UserID,
		Role: Role,
		IssuedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return paylaod,nil
}

func(p *Payload) Valid()error{
	if time.Now().After(p.ExpiredAt){
		return ErrExpiredToken
	}
	return nil
}

