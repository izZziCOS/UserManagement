package service

import (
	"context"
	"testing"

	"github.com/izzzicos/UserManagement/api/internal/contracts"
	"github.com/izzzicos/UserManagement/api/internal/mocks"
	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockNotifier := new(mocks.Notifier)
	service := NewUserService(mockRepo, mockNotifier)

	req := contracts.CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "password123",
		Country:   "US",
	}

	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.User")).
		Return(nil).
		Run(func(args mock.Arguments) {
			user := args.Get(1).(*models.User)
			assert.Equal(t, req.FirstName, user.FirstName)
			assert.Equal(t, req.LastName, user.LastName)
			assert.Equal(t, req.Email, user.Email)
			assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)))
		})

	mockNotifier.On("NotifyUserCreated", mock.Anything, mock.AnythingOfType("*models.User")).
		Return(nil)

	_, err := service.CreateUser(context.Background(), req)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUsers(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockNotifier := new(mocks.Notifier)
	service := NewUserService(mockRepo, mockNotifier)

	firstName := "John"
	country := "US"

	reqFilter := contracts.UserFilter{
		FirstName: &firstName,
		Country:   &country,
	}

	reqPagination := contracts.Pagination{
		Page:  1,
		Limit: 1,
	}

	mockRepo.On("FindAll",
		mock.Anything,
		mock.AnythingOfType("repository.UserFilter"),
		mock.AnythingOfType("repository.Pagination"),
	).Return([]models.User{}, nil)

	responses, err := service.GetUsers(context.Background(), reqFilter, reqPagination)

	assert.NoError(t, err)
	assert.NotNil(t, responses)
	mockRepo.AssertExpectations(t)
}
