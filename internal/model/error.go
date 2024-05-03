package model

import "errors"

var (
	ErrUserNotFound = errors.New("user doesn't exists")
	ErrUserBlocked  = errors.New("user is blocked")
)
