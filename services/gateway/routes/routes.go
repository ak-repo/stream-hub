package routes

import (
	"github.com/ak-repo/stream-hub/api/authpb"
	"github.com/ak-repo/stream-hub/config"
	"github.com/ak-repo/stream-hub/services/gateway/clients"
	"github.com/ak-repo/stream-hub/services/gateway/handler"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func New(app *fiber.App, zlog *zap.Logger, cfg *config.Config, clients *clients.Clients) {

	api := app.Group("/api/v1")
	authRoutes(api, zlog, cfg, clients.Auth)

}

func authRoutes(r fiber.Router, zlog *zap.Logger, cfg *config.Config, authClient authpb.AuthServiceClient) {

	auth := handler.NewAuthHandler(authClient, zlog)

	r.Post("/login", auth.Login)
	r.Post("/register", auth.Register)

}
