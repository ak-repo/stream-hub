

package models

import "time"

type User struct {
	ID            string    `db:"id"`            // unique identifier
	Username      string    `db:"username"`      // mapped to 'name' in gRPC
	Email         string    `db:"email"`         // email address
	PasswordHash  string    `db:"password_hash"` // hashed password
	Role          string    `db:"role"`          // user / admin / banned
	EmailVerified bool      `db:"email_verified"`// email confirmation status
	IsBlocked     bool      `db:"is_blocked"`    // user blocked by another user
	CreatedAt     time.Time `db:"created_at"`    // creation timestamp
}
