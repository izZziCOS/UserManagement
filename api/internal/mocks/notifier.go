package mocks

import (
	"context"

	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/stretchr/testify/mock"
)

type Notifier struct {
	mock.Mock
}

func (m *Notifier) NotifyUserCreated(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *Notifier) NotifyUserUpdated(ctx context.Context, oldUser, newUser *models.User) error {
	args := m.Called(ctx, oldUser, newUser)
	return args.Error(0)
}

func (m *Notifier) NotifyUserDeleted(ctx context.Context, userID string) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func (m *Notifier) StartTestConsumer() {
	m.Called()
}

func (m *Notifier) Close() error {
	args := m.Called()
	return args.Error(0)
}
