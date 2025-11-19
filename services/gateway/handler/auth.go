package handler

import (
	"context"
	"time"

	"github.com/ak-repo/stream-hub/api/authpb"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AuthHandler struct {
	client authpb.AuthServiceClient
	zlog   *zap.Logger
}

func NewAuthHandler(cli authpb.AuthServiceClient, log *zap.Logger) *AuthHandler {
	return &AuthHandler{client: cli, zlog: log}
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	req := new(authpb.LoginRequest)

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON"})
	}

	// Timeout-based context
	grpcCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := h.client.Login(grpcCtx, req)
	if err != nil {
		h.zlog.Error("gRPC login failed", zap.Error(err))
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "auth service unavailable"})
	}

	h.zlog.Info("login request", zap.String("email", req.Email))

	return ctx.JSON(resp)
}

func (h *AuthHandler) Register(ctx *fiber.Ctx) error {
	req := new(authpb.RegisterRequest)

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON"})
	}

	// Timeout-based context
	grpcCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := h.client.Register(grpcCtx, req)
	if err != nil {
		h.zlog.Error("gRPC register failed", zap.Error(err))
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "auth service unavailable"})
	}

	h.zlog.Info("login request", zap.String("email", req.Email))

	return ctx.JSON(resp)
}
