// Package models defines the core data structures and transformations
package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user entity in the system
type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Nickname  string    `json:"nickname" validate:"required"`
	Password  string    `json:"-" validate:"required,min=8"`
	Email     string    `json:"email" validate:"required,email"`
	Country   string    `json:"country" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserResponse represents the API response format for user data
type UserResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToResponse converts a User model to its API response representation
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Nickname:  u.Nickname,
		Email:     u.Email,
		Country:   u.Country,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// NewUser creates a new User instance with initialized fields
func NewUser() *User {
	return &User{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
