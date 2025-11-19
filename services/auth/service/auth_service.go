package service

import (
	"context"

	"github.com/ak-repo/stream-hub/services/common/repo"
)

type AuthService struct {
	repo *repo.CommonRepository
}

func NewAuthService(repo *repo.CommonRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(ctx context.Context, username, email, password string) (string, error) {

	query := "INSERT INTO users (username,email,password_hash,role)VALUES ($1,$2,$3,$4)RETURNING id"
	row := s.repo.QueryRow(ctx, query, username, email, password, "user")
	var userID string
	if err := row.Scan(&userID); err != nil {
		return "", err
	}

	return userID, nil

}
