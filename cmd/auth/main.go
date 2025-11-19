package main

import (
	"context"
	"log"
	"net"

	"github.com/ak-repo/stream-hub/api/authpb"
	"github.com/ak-repo/stream-hub/config"
	"github.com/ak-repo/stream-hub/pkg/db"
	"github.com/ak-repo/stream-hub/pkg/logger"
	"github.com/ak-repo/stream-hub/services/auth/server"
	"github.com/ak-repo/stream-hub/services/auth/service"
	"github.com/ak-repo/stream-hub/services/common/repo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load failed: %v", err)
	}

	zlog := logger.New(cfg.Logging.Level, cfg.Logging.Format)
	defer zlog.Sync()

	//db
	pgDB, err := db.NewPostgresDB(context.Background(), cfg)
	if err != nil {
		zlog.Fatal("failed to connect db:", zap.Error(err))
	}
	defer pgDB.Close()
	commonRepo := repo.NewCommonRepository(pgDB)

	service := service.NewAuthService(commonRepo)
	server := server.NewAuthServer(service)

	addr := ":" + cfg.Services.Auth.Port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		zlog.Fatal("listen failed", zap.Error(err))
	}

	grpcServer := grpc.NewServer()

	authpb.RegisterAuthServiceServer(grpcServer, server)

	zlog.Info("auth service started", zap.String("addr", addr))

	if err := grpcServer.Serve(lis); err != nil {
		zlog.Fatal("grpc server failed", zap.Error(err))
	}
}
