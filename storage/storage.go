package storage

import (
	"context"
	"editory/editory_user_service/genproto/user_service"
)

type StorageI interface {
	CloseDB()
	User() UserRepoI
}

// SmstorageI ...
type UserRepoI interface {
	CreateUser(ctx context.Context, req *user_service.Customer) (resp *user_service.CustomerResponse, err error)
	GetUser(ctx context.Context, req *user_service.GetCustomerRequest) (resp *user_service.CustomerResponse, err error)
	UpdateUser(ctx context.Context, req *user_service.Customer) (resp *user_service.CustomerResponse, err error)
	DeleteUser(ctx context.Context, req *user_service.DeleteRequest) (resp *user_service.Empty, err error)
	ExistsCustomer(ctx context.Context, req *user_service.ExistsCustomerRequest) (resp *user_service.ExistsCustomerResponse, err error)
	EditCustomerPhoneNumber(ctx context.Context, req *user_service.EditCustomerPhoneNumberRequest) (resp *user_service.Empty, err error)
}
