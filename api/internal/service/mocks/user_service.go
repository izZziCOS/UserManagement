package mocks

import (
	"context"

	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/izzzicos/UserManagement/api/internal/service"
	"github.com/stretchr/testify/mock"
)

type UserService struct {
	mock.Mock
}

func (m *UserService) CreateUser(ctx context.Context, req service.CreateUserRequest) (*models.UserResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*models.UserResponse), args.Error(1)
}

func (m *UserService) UpdateUser(ctx context.Context, id string, req service.UpdateUserRequest) (*models.UserResponse, error) {
	args := m.Called(ctx, id, req)
	return args.Get(0).(*models.UserResponse), args.Error(1)
}

func (m *UserService) DeleteUser(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *UserService) GetUsers(ctx context.Context, filter service.UserFilter, pagination service.Pagination) ([]models.UserResponse, error) {
	args := m.Called(ctx, filter, pagination)
	return args.Get(0).([]models.UserResponse), args.Error(1)
}

func (m *UserService) GetUserByID(ctx context.Context, id string) (*models.UserResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.UserResponse), args.Error(1)
}
