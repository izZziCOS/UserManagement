package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/izzzicos/UserManagement/api/internal/service/mocks"
	"github.com/stretchr/testify/mock"
)

func setupRouter(service *mocks.UserService) *gin.Engine {
	handler := NewUserHandler(service)
	router := gin.Default()
	handler.RegisterRoutes(router)
	return router
}

func TestUserHandler_CreateUser(t *testing.T) {
	mockService := new(mocks.UserService)
	router := setupRouter(mockService)

	createReq := map[string]interface{}{
		"first_name": "John",
		"last_name":  "Doe",
		"email":      "john@example.com",
		"password":   "password123",
		"country":    "US",
	}

	expectedUser := &models.UserResponse{
		ID:        "test-id",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}

	mockService.On("CreateUser", mock.Anything, mock.AnythingOfType("service.CreateUserRequest")).
		Return(expectedUser, nil)

	body, _ := json.Marshal(createReq)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.UserResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, expectedUser.ID, response.ID)
	assert.Equal(t, expectedUser.FirstName, response.FirstName)
	mockService.AssertExpectations(t)
}

// Add similar tests for other handler methods
