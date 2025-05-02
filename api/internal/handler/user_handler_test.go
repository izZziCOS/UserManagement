package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/izzzicos/UserManagement/api/internal/contracts"
	"github.com/izzzicos/UserManagement/api/internal/mocks"
	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler_CreateUser(t *testing.T) {
	mockService := new(mocks.UserService)
	handler := NewUserHandler(mockService)
	router := gin.Default()
	handler.RegisterRoutes(router)

	createReq := map[string]interface{}{
		"first_name": "John",
		"last_name":  "Doe",
		"nickname":   "johndoe",
		"email":      "john@example.com",
		"password":   "password123",
		"country":    "US",
	}

	expectedUser := &models.UserResponse{
		ID:        "test-id",
		FirstName: "John",
		LastName:  "Doe",
		Nickname:  "johndoe",
		Email:     "john@example.com",
		Country:   "US",
	}

	mockService.On("CreateUser", mock.Anything, mock.MatchedBy(func(req contracts.CreateUserRequest) bool {
		return req.FirstName == "John" &&
			req.LastName == "Doe" &&
			req.Nickname == "johndoe" &&
			req.Email == "john@example.com" &&
			req.Country == "US"
	})).Return(expectedUser, nil)

	body, _ := json.Marshal(createReq)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	
	assert.Equal(t, expectedUser.ID, response.ID)
	assert.Equal(t, expectedUser.FirstName, response.FirstName)
	assert.Equal(t, expectedUser.LastName, response.LastName)
	assert.Equal(t, expectedUser.Nickname, response.Nickname)
	
	mockService.AssertExpectations(t)
}