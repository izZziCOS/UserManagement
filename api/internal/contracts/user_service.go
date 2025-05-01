package contracts

import (
	"context"

	"github.com/izzzicos/UserManagement/api/internal/models"
)

type UserService interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*models.UserResponse, error)
	UpdateUser(ctx context.Context, id string, req UpdateUserRequest) (*models.UserResponse, error)
	DeleteUser(ctx context.Context, id string) error
	GetUsers(ctx context.Context, filter UserFilter, pagination Pagination) ([]models.UserResponse, error)
}

type UserFilter struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Nickname  *string `json:"nickname,omitempty"`
	Email     *string `json:"email,omitempty"`
	Country   *string `json:"country,omitempty"`
}

type Pagination struct {
	Page  int
	Limit int
}

type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Nickname  string `json:"nickname" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
	Email     string `json:"email" validate:"required,email"`
	Country   string `json:"country" validate:"required"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}
