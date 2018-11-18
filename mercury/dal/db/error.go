package db

import "errors"

var (
	ErrUserExists = errors.New("username is exist")
)
