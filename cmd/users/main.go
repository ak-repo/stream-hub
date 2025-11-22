package main

import (
	"context"
	"log"
	"net"

	"github.com/ak-repo/stream-hub/api/userspb"
	"github.com/ak-repo/stream-hub/config"
	"github.com/ak-repo/stream-hub/pkg/db"
	"github.com/ak-repo/stream-hub/pkg/logger"
	"github.com/ak-repo/stream-hub/services/common/interceptors"
	"github.com/ak-repo/stream-hub/services/common/repo"
	"github.com/ak-repo/stream-hub/services/users/server"
	"github.com/ak-repo/stream-hub/services/users/service"
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
		zlog.Fatal("failed to connect DB: ", zap.Error(err))
	}
	defer pgDB.Close()
	commonRepo := repo.NewCommonRepository(pgDB)

	service := service.NewUsersService(commonRepo)
	server := server.NewUserServer(cfg, service)

	addr := ":" + cfg.Services.Users.Port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		zlog.Fatal("listen faild- users", zap.Error(err))
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.AppErrorInterceptor()))

	userspb.RegisterUserServiceServer(grpcServer, server)

	zlog.Info("users service started", zap.String("addr", addr))

	if err := grpcServer.Serve(lis); err != nil {
		zlog.Fatal("grpc users server failed :", zap.Error(err))
	}
}
