package service

import (
	"context"

	"github.com/ak-repo/stream-hub/pkg/errors"
	"github.com/ak-repo/stream-hub/pkg/utils"
	"github.com/ak-repo/stream-hub/services/common/models"
	"github.com/ak-repo/stream-hub/services/common/repo"
	"github.com/jackc/pgx/v5"
)

type AuthService struct {
	repo *repo.CommonRepository
}

func NewAuthService(repo *repo.CommonRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(ctx context.Context, username, email, password string) error {
	// validate
	if err := utils.ValidateRegister(username, email, password); err != nil {
		return errors.New(errors.CodeInvalidInput, "invalid user input", err)
	}

	// check existing email
	var existingID string
	err := s.repo.QueryRow(ctx, "SELECT id FROM users WHERE email=$1", email).Scan(&existingID)
	if err == nil {
		// email is already taken
		return errors.New(errors.CodeConflict, "email already in use", nil)
	}

	// hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return errors.New(errors.CodeInternal, "failed to hash password", err)
	}

	// create new user
	row := s.repo.QueryRow(ctx, `
		INSERT INTO users (username, email, password_hash, role, email_verified)
		VALUES ($1, $2, $3, $4, FALSE)
		RETURNING id, username, email, role
	`, username, email, hashedPassword, "user")

	user := &models.User{}
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Role); err != nil {
		return errors.New(errors.CodeInternal, "failed to scan created user", err)
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*models.User, error) {
	if err := utils.ValidateLogin(email, password); err != nil {
		return nil, errors.New(errors.CodeInvalidInput, "invalid login input", err)
	}

	user, err := s.GetUserByEmail(ctx, email)
	if err != nil {
		// user not found → unauthorized
		return nil, errors.New(errors.CodeUnauthorized, "invalid email or password", nil)
	}

	// password mismatch → unauthorized
	if !utils.ComparePassword(user.PasswordHash, password) {
		return nil, errors.New(errors.CodeUnauthorized, "invalid email or password", nil)
	}

	user.PasswordHash = ""
	return user, nil
}

func (s *AuthService) VerifyMagicLink(ctx context.Context, email string) (*models.User, error) {
	query := `UPDATE users SET email_verified = TRUE WHERE email = $1`

	res, err := s.repo.Exec(ctx, query, email)
	if err != nil {
		return nil, errors.New(errors.CodeInternal, "database update error", err)
	}

	if res.RowsAffected() == 0 {
		return nil, errors.New(errors.CodeNotFound, "no user found with this email", nil)
	}

	// fetch updated user
	user, err := s.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = ""
	return user, nil
}

func (s *AuthService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	row := s.repo.QueryRow(ctx, `
		SELECT id, username, email, password_hash, role, email_verified, created_at
		FROM users
		WHERE email = $1
	`, email)

	user := &models.User{}
	err := row.Scan(
		&user.ID, &user.Username, &user.Email,
		&user.PasswordHash, &user.Role,
		&user.EmailVerified, &user.CreatedAt,
	)

	if err != nil {
		// no rows → not found
		if err == pgx.ErrNoRows {
			return nil, errors.New(errors.CodeNotFound, "user not found", nil)
		}

		// other scan/db errors
		return nil, errors.New(errors.CodeInternal, "failed to scan user", err)
	}

	return user, nil
}
