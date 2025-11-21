package dto

// UserResponseDTO is used to send user info in responses
type UserResponseDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// AuthResponseDTO is used for responses after login or registration
type LoginResponseDTO struct {
	User  UserResponseDTO `json:"user"`
	Token string          `json:"token,omitempty"` // JWT token, optional
}

// AuthResponseDTO is used for responses after login or registration
type RegisterResponseDTO struct {
	User    UserResponseDTO `json:"user"`
	Message string          `json:"message,omitempty"` // optional
}
