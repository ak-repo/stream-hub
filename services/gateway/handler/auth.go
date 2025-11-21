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

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	h.zlog.Info("login called")
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
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}

	h.zlog.Info("login request", zap.String("email", req.Email))

	// JWT
	accessToken, aExp, err := h.jwtManager.GenerateAccessToken(resp.User.Id, resp.User.Email)
	if err != nil {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": errors.ErrTokenGenerate.Error()})
	}
	refrToken, rExp, err := h.jwtManager.GenerateRefreshToken(resp.User.Id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": errors.ErrTokenGenerate.Error()})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "refresh",
		Value:   refrToken,
		Path:    "/",
		Domain:  "",
		Expires: rExp,
	})

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "login successful", "token": accessToken, "exp": aExp})
}
