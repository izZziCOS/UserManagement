package repository

import (
	"context"
	"testing"

	"github.com/izzzicos/UserManagement/api/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	return gormDB, mock
}

func TestUserRepository_Create(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := NewUserRepository(gormDB)

	user := &models.User{
		ID:        "test-id",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(user.ID, user.FirstName, user.LastName, user.Nickname, 
			user.Password, user.Email, user.Country, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Create(context.Background(), user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// Add similar tests for other repository methods