package domain

import "errors"

var (
	ErrUserDuplicate       = errors.New("user already exists")
	ErrEmptyGuid           = errors.New("guid is empty")
	ErrUserNotFound        = errors.New("user not found")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)
