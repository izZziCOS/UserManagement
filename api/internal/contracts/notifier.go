// Notifier defines notification related communication
package contracts

import (
	"context"

	"github.com/izzzicos/UserManagement/api/internal/models"
)

// Notifier defines the interface for sending notifications
type Notifier interface {
	// NotifyUserCreated publishes a notification when a user is created
	NotifyUserCreated(ctx context.Context, user *models.User) error
	// NotifyUserUpdated publishes a notification when a user is updated
	NotifyUserUpdated(ctx context.Context, oldUser, newUser *models.User) error
	// NotifyUserDeleted publishes a notification when a user is deleted
	NotifyUserDeleted(ctx context.Context, userID string) error
	// StartTestConsumer initiates a test message consumer (for development only)
	StartTestConsumer()
	// Close cleanly shuts down the notifier and releases resources
	Close() error
}
