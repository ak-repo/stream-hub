package errors

import "errors"

var (

	// JWT
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)
