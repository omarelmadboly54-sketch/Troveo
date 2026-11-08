package apperrors

import "errors"

var (
	//Users
	ErrUserAlreadyExists = errors.New("User already exists")
	ErrInvalidCredentials=errors.New("Invalid credentials")
	ErrUserNotFound=errors.New("User not found")
)
