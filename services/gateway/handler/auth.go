package handler

import (
	"context"
	"time"

	"github.com/ak-repo/stream-hub/api/authpb"
	"github.com/ak-repo/stream-hub/pkg/errors"
	"github.com/ak-repo/stream-hub/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AuthHandler struct {
	client     authpb.AuthServiceClient
	zlog       *zap.Logger
	jwtManager *jwt.JWTManager
}

func NewAuthHandler(cli authpb.AuthServiceClient, log *zap.Logger, jwt *jwt.JWTManager) *AuthHandler {
	return &AuthHandler{client: cli, zlog: log, jwtManager: jwt}
}

func withGRPCTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}

func (h *AuthHandler) Register(ctx *fiber.Ctx) error {
	req := &authpb.RegisterRequest{}
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	gc, cancel := withGRPCTimeout()
	defer cancel()
	resp, err := h.client.Register(gc, req)
	if err != nil {
		h.zlog.Error("grpc register", zap.Error(err))
		code, body := errors.GRPCToFiber(err)
		return ctx.Status(code).JSON(body)
	}
	return ctx.JSON(resp)
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	req := &authpb.LoginRequest{}
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	gc, cancel := withGRPCTimeout()
	defer cancel()
	resp, err := h.client.Login(gc, req)
	if err != nil {
		h.zlog.Error("grpc login", zap.Error(err))
		code, body := errors.GRPCToFiber(err)
		return ctx.Status(code).JSON(body)
	}
	access, aExp, err := h.jwtManager.GenerateAccessToken(resp.User.Id, resp.User.Email)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "token creation failed"})
	}
	refresh, rExp, err := h.jwtManager.GenerateRefreshToken(resp.User.Id)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "token creation failed"})
	}
	ctx.Cookie(&fiber.Cookie{Name: "refresh", Value: refresh, Path: "/", Expires: rExp, HTTPOnly: true})
	return ctx.JSON(fiber.Map{"message": "login successful", "token": access, "exp": aExp})
}

func (h *AuthHandler) SendMagicLink(ctx *fiber.Ctx) error {
	req := &authpb.SendMagicLinkRequest{}
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	gc, cancel := withGRPCTimeout()
	defer cancel()
	resp, err := h.client.SendMagicLink(gc, req)
	if err != nil {
		h.zlog.Error("grpc magic link", zap.Error(err))
		code, body := errors.GRPCToFiber(err)
		return ctx.Status(code).JSON(body)
	}
	return ctx.JSON(resp)
}

func (h *AuthHandler) VerifyMagicLink(ctx *fiber.Ctx) error {
	req := &authpb.VerifyMagicLinkRequest{
		Email: ctx.Query("email"),
		Token: ctx.Query("token"),
	}
	if req.Email == "" || req.Token == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "missing email or token"})
	}
	gc, cancel := withGRPCTimeout()
	defer cancel()
	resp, err := h.client.VerifyMagicLink(gc, req)
	if err != nil {
		h.zlog.Error("grpc verify magic link", zap.Error(err))
		code, body := errors.GRPCToFiber(err)
		return ctx.Status(code).JSON(body)
	}
	return ctx.JSON(resp)
}
