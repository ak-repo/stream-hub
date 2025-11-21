package server

import (
	"context"
	"log"

	"github.com/ak-repo/stream-hub/api/authpb"
	"github.com/ak-repo/stream-hub/config"
	"github.com/ak-repo/stream-hub/pkg/errors"
	"github.com/ak-repo/stream-hub/pkg/helper"
	"github.com/ak-repo/stream-hub/pkg/jwt"
	"github.com/ak-repo/stream-hub/services/auth/service"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer
	service    *service.AuthService
	cfg        *config.Config
	jwtManager *jwt.JWTManager
}

func NewAuthServer(svc *service.AuthService, cfg *config.Config, jwt *jwt.JWTManager) *AuthServer {
	return &AuthServer{service: svc, cfg: cfg, jwtManager: jwt}
}

//
// ───────────────────────────────────────────────
//   Register
// ───────────────────────────────────────────────
//
func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {

	err := s.service.Register(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		return nil, err // interceptor handles ToGRPC
	}

	return &authpb.RegisterResponse{
		LinkSend: true,
		Message:  "User registered successfully",
	}, nil
}

//
// ───────────────────────────────────────────────
//   Login
// ───────────────────────────────────────────────
//
func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {

	user, err := s.service.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err // interceptor handles it
	}

	res := &authpb.AuthUser{
		Id:            user.ID,
		Email:         user.Email,
		Role:          user.Role,
		EmailVerified: user.EmailVerified,
	}

	return &authpb.LoginResponse{User: res}, nil
}

//
//   Send Magic Link
//
func (s *AuthServer) SendMagicLink(ctx context.Context, req *authpb.SendMagicLinkRequest) (*authpb.SendMagicLinkResponse, error) {

	token, out, err := s.jwtManager.GenerateAccessToken("0", req.Email)
	if err != nil {
		return nil, err // interceptor handles
	}

	magicLink := s.cfg.Services.Gateway.Host + ":" +
		s.cfg.Services.Gateway.Port +
		"/verify-link?email=" + req.Email +
		"&token=" + token

	// SendGrid
	from := mail.NewEmail("StreamHub", "ak506lap@gmail.com")
	to := mail.NewEmail("", req.Email)

	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.Subject = "Your Magic Login Link"
	message.SetTemplateID("d-9f3316146d5d46e4ba3efdc8c6ba98c6")

	p := mail.NewPersonalization()
	p.AddTos(to)
	p.SetDynamicTemplateData("magic_link", magicLink)
	message.AddPersonalizations(p)

	client := sendgrid.NewSendClient(s.cfg.SendGrid.Key)
	response, err := client.Send(message)

	if err != nil || response.StatusCode != 202 {
		log.Println("SendGrid Error:", err)
		return nil, errors.New(errors.CodeInternal, "failed to send email", err)
	}

	return &authpb.SendMagicLinkResponse{
		MagicLink: magicLink,
		Expire:    helper.TimeToString(out),
		Message:   "verification link sent to " + req.Email,
	}, nil
}

//
// ───────────────────────────────────────────────
//   Verify Magic Link
// ───────────────────────────────────────────────
//
func (s *AuthServer) VerifyMagicLink(ctx context.Context, req *authpb.VerifyMagicLinkRequest) (*authpb.VerifyMagicLinkResponse, error) {

	claims, err := s.jwtManager.ValidateToken(req.Token)
	if err != nil {
		return nil, err // interceptor handles
	}

	if claims.Email != req.Email {
		return nil, errors.New(errors.CodeUnauthorized, "token email mismatch", nil)
	}

	user, err := s.service.VerifyMagicLink(ctx, claims.Email)
	if err != nil {
		return nil, err
	}

	resp := &authpb.AuthUser{
		Id:            user.ID,
		Email:         user.Email,
		Role:          user.Role,
		EmailVerified: user.EmailVerified,
		CreatedAt:     helper.TimeToString(user.CreatedAt),
	}

	return &authpb.VerifyMagicLinkResponse{User: resp}, nil
}
