// Provides mocks for Notifier
package mocks

import (
	"context"

	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/stretchr/testify/mock"
)

// Notifier is a mock implementation of contracts.Notifier interface
type Notifier struct {
	mock.Mock
}

// NotifyUserCreated mocks the user creation notification
func (m *Notifier) NotifyUserCreated(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// NotifyUserUpdated mocks the user update notification
func (m *Notifier) NotifyUserUpdated(ctx context.Context, oldUser, newUser *models.User) error {
	args := m.Called(ctx, oldUser, newUser)
	return args.Error(0)
}

// NotifyUserDeleted mocks the user deletion notification
func (m *Notifier) NotifyUserDeleted(ctx context.Context, userID string) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

// StartTestConsumer mocks the test consumer initialization
func (m *Notifier) StartTestConsumer() {
	m.Called()
}

// Close mocks the notifier cleanup
func (m *Notifier) Close() error {
	args := m.Called()
	return args.Error(0)
}
