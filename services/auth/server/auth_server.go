package server

import (
	"context"

	"github.com/ak-repo/stream-hub/api/authpb"
	"github.com/ak-repo/stream-hub/services/auth/service"
)

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer
	service *service.AuthService
}

func NewAuthServer(svc *service.AuthService) *AuthServer {
	return &AuthServer{service: svc}
}

// Register
func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {

	// Call the service to register the user
	user, err := s.service.Register(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	res := &authpb.AuthUser{
		Id:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}

	return &authpb.RegisterResponse{
		User:    res,
		Message: "User registered successfully",
	}, nil
}

// Login
func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {

	user, err := s.service.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	res := &authpb.AuthUser{
		Id:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}

	return &authpb.LoginResponse{
		User: res,
	}, nil
}
