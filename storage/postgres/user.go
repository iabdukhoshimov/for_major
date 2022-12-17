package postgres

import (
	"context"
	"database/sql"
	"editory/editory_user_service/genproto/user_service"
	"editory/editory_user_service/storage"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

// NewSmsRepo ...
func NewUserRepo(db *pgxpool.Pool) storage.UserRepoI {
	return &userRepo{db: db}
}

// Create User...
func (r *userRepo) CreateUser(ctx context.Context, req *user_service.Customer) (resp *user_service.CustomerResponse, err error) {

	resp = &user_service.CustomerResponse{
		Customer: &user_service.Customer{},
	}

	var id string

	query := `INSERT INTO users (
		name,
		phone_number,
		email,
		address,
		work_place,
		date_of_birth,
		gender,
		phone_verified,
		active,
		avatar,
		created_at
		) 
		VALUES (
			$1, 
			$2,
			$3, 
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			NOW()
			) RETURNING id`

	err = r.db.QueryRow(ctx, query,

		req.Name,
		req.PhoneNumber,
		req.Email,
		req.Address,
		req.WorkPlace,
		time.Now(),
		req.Gender,
		req.PhoneVerified,
		req.Active,
		req.Avatar).Scan(&id)
	if err != nil {
		fmt.Println("err>>", err)
		return resp, err
	}

	resp.Customer.Id = id

	return
}

// Get User...
func (r *userRepo) GetUser(ctx context.Context, req *user_service.GetCustomerRequest) (resp *user_service.CustomerResponse, err error) {

	resp = &user_service.CustomerResponse{
		Customer: &user_service.Customer{},
	}

	var (
		updatedAt sql.NullString
		createdAt sql.NullString
		deletedAt sql.NullString
		subQuery  string
		id        string
	)

	// u.user_id defines relationship between
	query := `SELECT
		u.id,
		u.name,
		u.phone_number,
		u.email,
		u.address,
		u.work_place,
		u.avatar,
		COALESCE(to_char(u.date_of_birth, 'MM-DD-YYYY HH24:MI:SS'), ''),
		u.created_at,
		u.updated_at,
		u.deleted_at,
		u.gender,
		u.phone_verified,
		u.active
		FROM users u`
	if req.Id != "" {
		subQuery = ` 
			WHERE u.id = $1 AND u.deleted_at is null`
		id = req.Id
	} else if req.PhoneNumber != "" {
		subQuery = ` 
			WHERE u.phone_number = $1 AND u.deleted_at is null`
		id = req.PhoneNumber
	}
	query = query + subQuery

	err = r.db.QueryRow(ctx, query, id).Scan(
		&resp.Customer.Id,
		&resp.Customer.Name,
		&resp.Customer.PhoneNumber,
		&resp.Customer.Email,
		&resp.Customer.Address,
		&resp.Customer.WorkPlace,
		&resp.Customer.Avatar,
		&resp.Customer.DateOfBirth,
		&createdAt,
		&updatedAt,
		&deletedAt,
		&resp.Customer.Gender,
		&resp.Customer.PhoneVerified,
		&resp.Customer.Active,
	)
	if createdAt.Valid {
		resp.Customer.CreatedAt = createdAt.String
	}
	if updatedAt.Valid {
		resp.Customer.UpdatedAt = updatedAt.String
	}
	if deletedAt.Valid {
		resp.Customer.DeletedAt = deletedAt.String
	}
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Update User...
func (r *userRepo) UpdateUser(ctx context.Context, req *user_service.Customer) (resp *user_service.CustomerResponse, err error) {

	resp = &user_service.CustomerResponse{}

	// I write this query to update user and set updated_at to current timestamp
	query := `UPDATE users SET
		name = $1,
		email = $2,
		address = $3,
		work_place = $4,
		date_of_birth = $5,
		gender = $6,
		phone_verified = $7,
		active = $8,
		avatar = $9,
		updated_at = NOW()
		WHERE id = $10
		`

	_, err = r.db.Exec(ctx, query,
		req.Name,
		req.Email,
		req.Address,
		req.WorkPlace,
		req.DateOfBirth,
		req.Gender,
		req.PhoneVerified,
		req.Active,
		req.Avatar,
		req.Id,
	)
	if err != nil {
		return resp, err
	}

	return
}

// Delete User...
func (r *userRepo) DeleteUser(ctx context.Context, req *user_service.DeleteRequest) (resp *user_service.Empty, err error) {

	resp = &user_service.Empty{}

	// I use soft delete and set deleted_at to current timestamp to keep data integrity
	query := `UPDATE users SET deleted_at = NOW() WHERE id = $1`

	_, err = r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return resp, err
	}

	return
}

func (r *userRepo) ExistsCustomer(ctx context.Context, req *user_service.ExistsCustomerRequest) (resp *user_service.ExistsCustomerResponse, err error) {

	var (
		existsUser string = ""
	)
	resp = &user_service.ExistsCustomerResponse{}
	resp.Exists = false
	fmt.Println(" req.PhoneNumber ", req.PhoneNumber)

	if len(req.PhoneNumber) > 0 {
		row := r.db.QueryRow(ctx, `
		SELECT id 
		FROM users
		WHERE phone_number =$1`, req.PhoneNumber)
		err := row.Scan(&existsUser)
		if err != nil {
			if err.Error() == "no rows in result set" {
				return resp, nil
			}
			return nil, err
		}

		// this is condition
		if len(existsUser) > 0 {
			resp.Exists = true
			return resp, nil
		}
	}

	return resp, nil

}

func (r *userRepo) EditCustomerPhoneNumber(ctx context.Context, req *user_service.EditCustomerPhoneNumberRequest) (resp *user_service.Empty, err error) {

	if req.GetPhoneNumber() != "" {
		_, err := r.db.Exec(ctx,
			`UPDATE users
			SET
				phone_number=$1
			WHERE id=$2 AND active=true`,
			req.GetPhoneNumber(),
			req.GetUserId(),
		)
		if err != nil {
			return nil, err
		}
	}

	return
}
