package routes

import (
	"github.com/ak-repo/stream-hub/api/authpb"
	"github.com/ak-repo/stream-hub/config"
	"github.com/ak-repo/stream-hub/pkg/jwt"
	"github.com/ak-repo/stream-hub/services/gateway/clients"
	"github.com/ak-repo/stream-hub/services/gateway/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

func New(app *fiber.App, zlog *zap.Logger, cfg *config.Config, clients *clients.Clients) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Authorization",
	}))
	api := app.Group("/api/v1")
	authRoutes(api, zlog, cfg, clients.Auth)

}

func authRoutes(r fiber.Router, zlog *zap.Logger, cfg *config.Config, authClient authpb.AuthServiceClient) {

	jwtMan := jwt.NewJWTManager(cfg.JWT.Secret, cfg.JWT.Expiry, cfg.JWT.Expiry*7)
	auth := handler.NewAuthHandler(authClient, zlog, jwtMan)

	r.Post("/login", auth.Login)
	r.Post("/register", auth.Register)

}
