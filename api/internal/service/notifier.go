package service

import (
	"context"

	"github.com/izzzicos/UserManagement/api/internal/models"
)

// Notifier defines the interface for sending notifications
type Notifier interface {
    NotifyUserCreated(ctx context.Context, user *models.User) error
    NotifyUserUpdated(ctx context.Context, oldUser, newUser *models.User) error
    NotifyUserDeleted(ctx context.Context, userID string) error
	StartTestConsumer()
	Close() error
}