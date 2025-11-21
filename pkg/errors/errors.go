package errors

import "errors"

var (

	// JWT
	ErrInvalidToken  = errors.New("invalid token")
	ErrExpiredToken  = errors.New("token expired")
	ErrTokenGenerate = errors.New("token generation failed")

	// AUTH
	ErrEmailAllre = errors.New("email already in use")
)

// errors.New("invalid email or password")
// errors.New("invalid email or password")
