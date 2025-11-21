package service

import (
	"context"
	"errors"

	"github.com/ak-repo/stream-hub/pkg/utils"
	"github.com/ak-repo/stream-hub/services/common/models"
	"github.com/ak-repo/stream-hub/services/common/repo"
)

type AuthService struct {
	repo *repo.CommonRepository
}

func NewAuthService(repo *repo.CommonRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(ctx context.Context, username, email, password string) (*models.User, error) {
	if err := utils.ValidateRegister(username, email, password); err != nil {
		return nil, err
	}

	// Check if email already exists
	existing := &models.User{}
	err := s.repo.QueryRow(ctx, "SELECT id FROM users WHERE email=$1", email).Scan(&existing.ID)
	if err == nil {
		return nil, errors.New("email already in use")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Insert into DB
	query := "INSERT INTO users (username,email,password_hash,role) VALUES ($1,$2,$3,$4) RETURNING id, username, email, role"
	row := s.repo.QueryRow(ctx, query, username, email, hashedPassword, "user")

	// Scan DB row into model
	user := &models.User{}
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Role); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*models.User, error) {

	// Validate input
	if err := utils.ValidateLogin(email, password); err != nil {
		return nil, err
	}

	// Get user from DB
	query := `SELECT id,username,email,password_hash,role FROM users WHERE email=$1`
	row := s.repo.QueryRow(ctx, query, email)
	user := &models.User{}
	var storedHash string

	err := row.Scan(&user.ID, &user.Username, &user.Email, &storedHash, &user.Role)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	//Compare password
	if ok := utils.ComparePassword(storedHash, password); !ok {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
