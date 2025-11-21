package models

import "time"

type User struct {
	ID            string    `db:"id"`
	Username      string    `db:"name"`
	Email         string    `db:"email"`
	PasswordHash  string    `db:"password_hash"`
	Role          string    `db:"role"` // user/admin
	EmailVerified bool      `db:"email_verified"`
	CreatedAt     time.Time `db:"created_at"`
}
