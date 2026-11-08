package tokens

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minsecretkeysize = 32

type JWTMaker struct {
	secretKey string
}

func NewJwtMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minsecretkeysize {
		return nil, fmt.Errorf("invalid secretkey size must be at least:%v",minsecretkeysize)
	}
	return &JWTMaker{secretKey: secretKey},nil
}

func(maker *JWTMaker) CreateToken(UserID int64, Role string, duration time.Duration)(string,error){
	payload,err:=NewPayloadToken(UserID,Role,duration)
	if err!=nil{
		return "",err
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":payload.UserID,
		"role":payload.Role,
		"iat":payload.IssuedAt.Unix(),
		"exp":payload.ExpiredAt.Unix(),
	})

	return token.SignedString([]byte(maker.secretKey))

}


func (maker *JWTMaker) VerifyToken(tokenString string) (*Payload, error) {
    claims := jwt.MapClaims{}

    parsedToken, err := jwt.ParseWithClaims(
        tokenString,
        claims,
        func(token *jwt.Token) (interface{}, error) {
            return []byte(maker.secretKey), nil
        },
        jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
        jwt.WithExpirationRequired(),
        jwt.WithIssuedAt(),
        jwt.WithJSONNumber(),
    )

    if err != nil {
        if errors.Is(err, jwt.ErrTokenExpired) {
            return nil, ErrExpiredToken
        }
        return nil, ErrInvalidToken
    }

    if !parsedToken.Valid {
        return nil, ErrInvalidToken
    }

    role, ok := claims["role"].(string)
    if !ok || role == "" {
        return nil, ErrInvalidToken
    }

    userIDNumber, ok := claims["user_id"].(json.Number)
    if !ok {
        return nil, ErrInvalidToken
    }

    userID, err := userIDNumber.Int64()
    if err != nil || userID <= 0 {
        return nil, ErrInvalidToken
    }

    exp, err := claims.GetExpirationTime()
    if err != nil || exp == nil {
        return nil, ErrInvalidToken
    }

    return &Payload{
        UserID:    userID,
        Role:     role,
        ExpiredAt: exp.Time,
    }, nil
}

