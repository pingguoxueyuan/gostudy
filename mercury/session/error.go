package session

import "errors"

var (
	ErrSessionNotExist      = errors.New("session not exists")
	ErrKeyNotExistInSession = errors.New("key not exists in session")
)
