package main

import (
	"context"
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"editory/editory_user_service/config"
	"editory/editory_user_service/grpc"
	"editory/editory_user_service/pkg/logger"
	"editory/editory_user_service/storage/postgres"
)

func main() {
	godotenv.Load(".env")
	cfg := config.Load()

	fmt.Printf("%+v\n", cfg)

	loggerLevel := logger.LevelDebug

	switch cfg.Project.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger(cfg.Project.ServiceName, loggerLevel)
	log.Info("Server is starting to run...")
	defer logger.Cleanup(log)

	pgStore, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}
	defer pgStore.CloseDB()

	grpcServer := grpc.SetUpServer(cfg, log, pgStore)

	lis, err := net.Listen("tcp", cfg.Project.RPCPort)
	if err != nil {
		log.Panic("net.Listen", logger.Error(err))
	}

	log.Info("GRPC: Server being started...", logger.String("port", cfg.Project.RPCPort))

	if err := grpcServer.Serve(lis); err != nil {
		log.Panic("grpcServer.Serve", logger.Error(err))

	}
}
