package main

import (
	"log"

	"github.com/ak-repo/stream-hub/api/authpb"
	"github.com/ak-repo/stream-hub/config"
	"github.com/ak-repo/stream-hub/pkg/helper"
	"github.com/ak-repo/stream-hub/pkg/logger"
	"github.com/ak-repo/stream-hub/services/gateway/clients"
	"github.com/ak-repo/stream-hub/services/gateway/routes"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	// ---- Load Config ----
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load failed: %v", err)
	}

	// ---- Local Development Overrides ----
	helper.OverrideLocal(cfg)

	// ---- Logger ----
	zlog := logger.New(cfg.Logging.Level, cfg.Logging.Format)
	defer zlog.Sync()

	// ---- Fiber App ----
	app := fiber.New()

	// ---- Create gRPC Client Container ----
	clientContainer := clients.NewContainer()

	// Initialize Auth Client
	initClient(zlog, clientContainer,
		cfg.Services.Auth.Host,
		cfg.Services.Auth.Port,
		func(conn *grpc.ClientConn) interface{} { return authpb.NewAuthServiceClient(conn) },
		&clientContainer.Auth,
	)

	// Clean up gRPC connections on exit
	defer clientContainer.CloseAll()

	// ---- Register Routes ----
	routes.New(app, zlog, cfg, clientContainer)

	// ---- Start Gateway HTTP Server ----
	addr := cfg.Services.Gateway.Host + ":" + cfg.Services.Gateway.Port
	zlog.Info("gateway started", zap.String("addr", addr))

	if err := app.Listen(addr); err != nil {
		zlog.Fatal("gateway startup failed", zap.Error(err))
	}
}

// initClient initializes a gRPC client in a clean, reusable way
func initClient(
	zlog *zap.Logger,
	container *clients.Clients,
	host, port string,
	factory func(*grpc.ClientConn) interface{},
	target interface{},
) {
	cli, conn, err := clients.NewClient(host, port, factory)
	if err != nil {
		zlog.Fatal("gRPC client initialization failed", zap.Error(err))
	}

	// Assign the concrete type via pointer to interface
	switch t := target.(type) {
	case *authpb.AuthServiceClient:
		*t = cli.(authpb.AuthServiceClient)
	default:
		zlog.Fatal("unsupported client type")
	}

	container.AddConn(conn)
}
