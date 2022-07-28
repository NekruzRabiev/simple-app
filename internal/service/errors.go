package service

import "errors"

var (
	ErrEmailOrPassword = errors.New("email or password is incorrect")
	ErrInvalidPassword = errors.New("invalid password")
	ErrUserExists      = errors.New("user's already exists")
)
