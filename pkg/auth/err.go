package auth

import "errors"

var (
	ErrCreateUserFailed = errors.New("failed to create user")
	ErrTokenNotActive   = errors.New("token not active")
)
