package service

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
    Country string
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