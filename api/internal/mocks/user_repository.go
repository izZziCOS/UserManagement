package mocks

import (
	"context"

	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/izzzicos/UserManagement/api/internal/repository"
	"github.com/stretchr/testify/mock"
)

// UserRepository is a mock implementation of repository.UserRepository interface
type UserRepository struct {
	mock.Mock
}

// Create mocks the user creation
func (m *UserRepository) Create(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// Update mocks updating the user
func (m *UserRepository) Update(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// Delete mocks deletion of user
func (m *UserRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// FindAll mock finding matching users by filters with pagination
func (m *UserRepository) FindAll(ctx context.Context, filter repository.UserFilter, pagination repository.Pagination) ([]models.User, error) {
	args := m.Called(ctx, filter, pagination)

	users, _ := args.Get(0).([]models.User)
	return users, args.Error(1)
}

// FindByID finds specific user by specified ID
func (m *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	args := m.Called(ctx, id)

	users, _ := args.Get(0).(*models.User)
	return users, args.Error(1)
}
