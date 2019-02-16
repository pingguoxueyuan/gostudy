package db

import "errors"

var (
	ErrUserExists        = errors.New("username is exist")
	ErrUserNotExists     = errors.New("username not exist")
	ErrUserPasswordWrong = errors.New("username or password not right")
	ErrRecordExists      = errors.New("record exist")
)
