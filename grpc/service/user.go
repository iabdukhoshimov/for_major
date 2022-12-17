package service

import (
	"context"
	"editory/editory_user_service/config"
	"editory/editory_user_service/genproto/user_service"
	"editory/editory_user_service/pkg/logger"
	"editory/editory_user_service/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SendService struct...
type userService struct {
	strg storage.StorageI
	cfg  config.Config
	log  logger.LoggerI

	user_service.UnimplementedCustomerServiceServer
}

// NewSendService ...
func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *userService {
	return &userService{
		strg: strg,
		cfg:  cfg,
		log:  log,
	}
}

//Create User...
func (u *userService) CreateUser(ctx context.Context, req *user_service.Customer) (resp *user_service.CustomerResponse, err error) {

	u.log.Info("CreateUser-->>", logger.Any("req", req))

	resp, err = u.strg.User().CreateUser(ctx, req)
	if err != nil {
		u.log.Error("!!!CreateCustomer->Customer->CreateCustomer--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())

	}

	return
}

//Get User...
func (u *userService) GetUser(ctx context.Context, req *user_service.GetCustomerRequest) (resp *user_service.CustomerResponse, err error) {

	u.log.Info("GetUser-->>", logger.Any("req", req))

	resp, err = u.strg.User().GetUser(ctx, req)
	if err != nil {
		u.log.Error("!!!GetUser->User->GetUser--->", logger.Error(err))
		return nil, err
	}
	return
}

//Update User...
func (u *userService) UpdateUser(ctx context.Context, req *user_service.Customer) (resp *user_service.CustomerResponse, err error) {

	u.log.Info("UpdateUser-->>", logger.Any("req", req))

	resp, err = u.strg.User().UpdateUser(ctx, req)
	if err != nil {
		u.log.Error("!!!UpdateUser->User->UpdateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

//Delete User...
func (u *userService) DeleteUser(ctx context.Context, req *user_service.DeleteRequest) (resp *user_service.Empty, err error) {

	u.log.Info("DeleteUser-->>", logger.Any("req", req))

	resp, err = u.strg.User().DeleteUser(ctx, req)
	if err != nil {
		u.log.Error("!!!DeleteUser->User->DeleteUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (u *userService) ExistsCustomer(ctx context.Context, req *user_service.ExistsCustomerRequest) (resp *user_service.ExistsCustomerResponse, err error) {

	u.log.Info("ExistsCustomer-->>", logger.Any("req", req))

	resp, err = u.strg.User().ExistsCustomer(ctx, req)
	if err != nil {
		u.log.Error("!!!ExistsCustomer->User->ExistsCustomer--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

//CloseDB ...
func (u *userService) CloseDB() {
	u.strg.CloseDB()
}
