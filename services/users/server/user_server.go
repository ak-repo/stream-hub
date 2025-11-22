package server

import (
	"context"

	"github.com/ak-repo/stream-hub/api/userspb"
	"github.com/ak-repo/stream-hub/config"
	"github.com/ak-repo/stream-hub/pkg/helper"
	"github.com/ak-repo/stream-hub/services/users/service"
)

type UserServer struct {
	userspb.UnimplementedUserServiceServer
	cfg          *config.Config
	usersService *service.UsersService
}

func NewUserServer(cfg *config.Config, service *service.UsersService) *UserServer {
	return &UserServer{
		cfg:          cfg,
		usersService: service,
	}
}

// FindByEmail finds a user by email
func (s *UserServer) FindByEmail(ctx context.Context, req *userspb.FindByEmailRequest) (*userspb.FindUserResponse, error) {
	user, err := s.usersService.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &userspb.FindUserResponse{User: &userspb.User{
		Id:        user.ID,
		Name:      user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: helper.TimeToString(user.CreatedAt),
		IsBlocked: user.IsBlocked,
	}}, nil
}

// FindById finds a user by ID
func (s *UserServer) FindById(ctx context.Context, req *userspb.FindByIdRequest) (*userspb.FindUserResponse, error) {
	user, err := s.usersService.FindById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &userspb.FindUserResponse{User: &userspb.User{
		Id:        user.ID,
		Name:      user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: helper.TimeToString(user.CreatedAt),
		IsBlocked: user.IsBlocked,
	}}, nil
}

// FindAllUsers returns all users
func (s *UserServer) FindAllUsers(ctx context.Context, req *userspb.Empty) (*userspb.FindAllUsersResponse, error) {
	users, err := s.usersService.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	resp := &userspb.FindAllUsersResponse{}
	for _, u := range users {
		resp.Users = append(resp.Users, &userspb.User{
			Id:        u.ID,
			Name:      u.Username,
			Email:     u.Email,
			Role:      u.Role,
			CreatedAt: helper.TimeToString(u.CreatedAt),
			IsBlocked: u.IsBlocked,
		})
	}
	return resp, nil
}

// BlockUser blocks a target user
func (s *UserServer) BlockUser(ctx context.Context, req *userspb.BlockUserRequest) (*userspb.BlockUserResponse, error) {
	if err := s.usersService.BlockUser(ctx, req.UserId, req.TargetUserId); err != nil {
		return nil, err
	}
	return &userspb.BlockUserResponse{Message: "user blocked successfully"}, nil
}

// UnblockUser unblocks a target user
func (s *UserServer) UnblockUser(ctx context.Context, req *userspb.BlockUserRequest) (*userspb.BlockUserResponse, error) {
	if err := s.usersService.UnblockUser(ctx, req.UserId, req.TargetUserId); err != nil {
		return nil, err
	}
	return &userspb.BlockUserResponse{Message: "user unblocked successfully"}, nil
}

// BanUser bans a target user
func (s *UserServer) BanUser(ctx context.Context, req *userspb.BanUserRequest) (*userspb.BanUserResponse, error) {
	if err := s.usersService.BanUser(ctx, req.TargetUserId); err != nil {
		return nil, err
	}
	return &userspb.BanUserResponse{Message: "user banned successfully"}, nil
}

// UnbanUser unbans a target user
func (s *UserServer) UnbanUser(ctx context.Context, req *userspb.BanUserRequest) (*userspb.BanUserResponse, error) {
	if err := s.usersService.UnbanUser(ctx, req.TargetUserId); err != nil {
		return nil, err
	}
	return &userspb.BanUserResponse{Message: "user unbanned successfully"}, nil
}
