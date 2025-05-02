// Provides mocks for User Service
package mocks

import (
	"context"

	service "github.com/izzzicos/UserManagement/api/internal/contracts"
	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/stretchr/testify/mock"
)

// UserService is a mock implementation of contracts.UserService interface
type UserService struct {
	mock.Mock
}

// CreateUser mocks user creation
func (m *UserService) CreateUser(ctx context.Context, req service.CreateUserRequest) (*models.UserResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.UserResponse), args.Error(1)
}

// UpdateUser mocks updating the user
func (m *UserService) UpdateUser(ctx context.Context, id string, req service.UpdateUserRequest) (*models.UserResponse, error) {
	args := m.Called(ctx, id, req)
	return args.Get(0).(*models.UserResponse), args.Error(1)
}

// DeleteUser mocks deletion of user
func (m *UserService) DeleteUser(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// GetUser mocks retrieving users with filters and pagination
func (m *UserService) GetUsers(ctx context.Context, filter service.UserFilter, pagination service.Pagination) ([]models.UserResponse, error) {
	args := m.Called(ctx, filter, pagination)
	return args.Get(0).([]models.UserResponse), args.Error(1)
}

// GetUserByID gets specific user with the id specified
func (m *UserService) GetUserByID(ctx context.Context, id string) (*models.UserResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.UserResponse), args.Error(1)
}
