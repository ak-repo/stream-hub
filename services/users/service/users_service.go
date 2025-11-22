package service

import (
	"context"

	"github.com/ak-repo/stream-hub/pkg/errors"
	"github.com/ak-repo/stream-hub/services/common/models"
	"github.com/ak-repo/stream-hub/services/common/repo"
	"github.com/jackc/pgx/v5"
)

type UsersService struct {
	repo *repo.CommonRepository
}

func NewUsersService(repo *repo.CommonRepository) *UsersService {
	return &UsersService{repo: repo}
}

// getUserByField retrieves a user by any column (e.g., "id" or "email")
func (s *UsersService) getUserByField(ctx context.Context, field string, value any) (*models.User, error) {
	query := `
		SELECT id, username, email, password_hash, role, email_verified, created_at
		FROM users
		WHERE ` + field + ` = $1
	`

	row := s.repo.QueryRow(ctx, query, value)

	user := &models.User{}
	err := row.Scan(
		&user.ID, &user.Username, &user.Email,
		&user.PasswordHash, &user.Role,
		&user.EmailVerified, &user.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New(errors.CodeNotFound, "user not found", nil)
		}
		return nil, errors.New(errors.CodeInternal, "failed to scan user", err)
	}

	return user, nil
}

// FindByEmail returns a user by email
func (s *UsersService) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.getUserByField(ctx, "email", email)
}

// FindById returns a user by ID
func (s *UsersService) FindById(ctx context.Context, id string) (*models.User, error) {
	return s.getUserByField(ctx, "id", id)
}

// FindAllUsers returns all users
func (s *UsersService) FindAllUsers(ctx context.Context) ([]*models.User, error) {
	rows, err := s.repo.Query(ctx, `SELECT id, username, email, password_hash, role, email_verified, created_at FROM users`)
	if err != nil {
		return nil, errors.New(errors.CodeInternal, "failed to query users", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		u := &models.User{}
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.Role, &u.EmailVerified, &u.CreatedAt); err != nil {
			return nil, errors.New(errors.CodeInternal, "failed to scan user", err)
		}
		users = append(users, u)
	}

	return users, nil
}

// BlockUser marks a user as blocked
func (s *UsersService) BlockUser(ctx context.Context, userID, targetUserID string) error {
	res, err := s.repo.Exec(ctx, `
		UPDATE users SET is_blocked = TRUE WHERE id = $1
	`, targetUserID)
	if err != nil {
		return errors.New(errors.CodeInternal, "failed to block user", err)
	}
	if res.RowsAffected() == 0 {
		return errors.New(errors.CodeNotFound, "target user not found", nil)
	}
	return nil
}

// UnblockUser unblocks a previously blocked user
func (s *UsersService) UnblockUser(ctx context.Context, userID, targetUserID string) error {
	res, err := s.repo.Exec(ctx, `
		UPDATE users SET is_blocked = FALSE WHERE id = $1
	`, targetUserID)
	if err != nil {
		return errors.New(errors.CodeInternal, "failed to unblock user", err)
	}
	if res.RowsAffected() == 0 {
		return errors.New(errors.CodeNotFound, "target user not found", nil)
	}
	return nil
}

// BanUser sets a user's role to banned
func (s *UsersService) BanUser(ctx context.Context, targetUserID string) error {
	res, err := s.repo.Exec(ctx, `
		UPDATE users SET role = 'banned' WHERE id = $1
	`, targetUserID)
	if err != nil {
		return errors.New(errors.CodeInternal, "failed to ban user", err)
	}
	if res.RowsAffected() == 0 {
		return errors.New(errors.CodeNotFound, "target user not found", nil)
	}
	return nil
}

// UnbanUser restores a user's role to normal
func (s *UsersService) UnbanUser(ctx context.Context, targetUserID string) error {
	res, err := s.repo.Exec(ctx, `
		UPDATE users SET role = 'user' WHERE id = $1
	`, targetUserID)
	if err != nil {
		return errors.New(errors.CodeInternal, "failed to unban user", err)
	}
	if res.RowsAffected() == 0 {
		return errors.New(errors.CodeNotFound, "target user not found", nil)
	}
	return nil
}
