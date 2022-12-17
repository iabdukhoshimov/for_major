package grpc

import (
	"editory/editory_user_service/config"
	"editory/editory_user_service/genproto/user_service"
	"editory/editory_user_service/grpc/service"
	"editory/editory_user_service/pkg/logger"
	"editory/editory_user_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	user_service.RegisterCustomerServiceServer(grpcServer, service.NewUserService(cfg, log, strg))

	reflection.Register(grpcServer)
	return
}
