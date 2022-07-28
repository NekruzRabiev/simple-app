package service

import "errors"

var (
	ErrEmailOrPassword = errors.New("email or password is incorrect")
	ErrOldPassword     = errors.New("incorrect old password")
	ErrInvalidPassword = errors.New("invalid password")
	ErrUserNotExist    = errors.New("user does not exist")
	ErrUserExists      = errors.New("user's already exists")
)
