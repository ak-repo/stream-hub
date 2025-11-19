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

// Login
func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {

	return &authpb.LoginResponse{Id: "100"}, nil
}

// Register
func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {

	id, err := s.service.Register(ctx, req.Name, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &authpb.RegisterResponse{Id: id}, nil
}
func (s *AuthServer) FindByEmail(ctx context.Context, req *authpb.FindByEmailRequest) (*authpb.FindUserResponse, error) {

	return &authpb.FindUserResponse{}, nil
}
func (s *AuthServer) FIndById(ctx context.Context, req *authpb.FindByIdRequest) (*authpb.FindUserResponse, error) {

	return &authpb.FindUserResponse{}, nil
}

func (s *AuthServer) FindUsers(ctx context.Context, req *authpb.Empty) (*authpb.FindAllUsersResponse, error) {

	return &authpb.FindAllUsersResponse{}, nil
}
