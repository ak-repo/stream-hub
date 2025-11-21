package models

type User struct {
	ID        string `db:"id"`
	Username  string `db:"name"`
	Email     string `db:"email"`
	Password  string `db:"password"` // hashed
	Role      string `db:"role"`     // user/admin
	CreatedAt string `db:"created_at"`
}
