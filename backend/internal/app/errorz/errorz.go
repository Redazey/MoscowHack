package errorz

import (
	"errors"
)

var (
	ErrTokenExpired        = errors.New("jwt token expired")
	ErrUserExists          = errors.New("user already exists")
	ErrValidation          = errors.New("jwt token validation failed")
	ErrUserNotFound        = errors.New("user not found")
	ErrPanicHandle         = errors.New("panic was handled")
	ErrServerIsNotResponse = errors.New("server is not response")
)
