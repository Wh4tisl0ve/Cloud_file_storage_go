package domain

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidUsername   = errors.New("username must be between 4 and 50 characters")
	ErrInvalidPassword   = errors.New("password must be between 8 and 20 characters")
)
