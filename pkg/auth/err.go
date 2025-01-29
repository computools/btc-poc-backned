package auth

import "errors"

var (
	ErrLoginClient      = errors.New("login client error")
	ErrCreateUserFailed = errors.New("failed to create user")
	ErrTokenNotActive   = errors.New("token not active")
)
