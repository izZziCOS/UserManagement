// Package contracts defines the interfaces and data structures
// for communication between application layers
package contracts

import (
	"context"

	"github.com/izzzicos/UserManagement/api/internal/models"
)

// UserService defines operations for user management
type UserService interface {
	// CreateUser registers a new user with the system
	CreateUser(ctx context.Context, req CreateUserRequest) (*models.UserResponse, error)
	// UpdateUser modifies an existing user's information
	UpdateUser(ctx context.Context, id string, req UpdateUserRequest) (*models.UserResponse, error)
	// DeleteUser removes a user from the system
	DeleteUser(ctx context.Context, id string) error
	// GetUsers retrieves a filtered and paginated list of users
	GetUsers(ctx context.Context, filter UserFilter, pagination Pagination) ([]models.UserResponse, error)
}

// UserFilter contains optional criteria for querying users.
type UserFilter struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Nickname  *string `json:"nickname,omitempty"`
	Email     *string `json:"email,omitempty"`
	Country   *string `json:"country,omitempty"`
}

// Pagination controls result set size for queries
type Pagination struct {
	Page  int
	Limit int
}

// CreateUserRequest defines the input for user creation
type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Nickname  string `json:"nickname" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
	Email     string `json:"email" validate:"required,email"`
	Country   string `json:"country" validate:"required"`
}

// UpdateUserRequest defines the input for user updates
type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}
