package mocks

import (
	"context"

	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/izzzicos/UserManagement/api/internal/repository"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) Create(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *UserRepository) Update(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *UserRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *UserRepository) FindAll(ctx context.Context, filter repository.UserFilter, pagination repository.Pagination) ([]models.User, error) {
	args := m.Called(ctx, filter, pagination)

	users, _ := args.Get(0).([]models.User)
	return users, args.Error(1)
}

func (m *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	args := m.Called(ctx, id)

	users, _ := args.Get(0).(*models.User)
	return users, args.Error(1)
}
