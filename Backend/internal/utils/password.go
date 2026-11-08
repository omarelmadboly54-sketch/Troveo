package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(Password string) (HashPassword string, err error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(Password),bcrypt.DefaultCost)
	if err!=nil{
		return "",fmt.Errorf("failed to hash password:%v",err)
	}

	return string(byte),nil
}

func VerifyPassword(hashpassword,password string)error{
	return bcrypt.CompareHashAndPassword([]byte(hashpassword),[]byte(password))
}



