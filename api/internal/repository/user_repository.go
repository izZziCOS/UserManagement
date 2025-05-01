// Package repository provides data access operations for user entities
// It defines the database interactions using GORM as the ORM layer
package repository

import (
	"context"

	"github.com/izzzicos/UserManagement/api/internal/models"
	"gorm.io/gorm"
)

// UserFilter contains optional filters for querying users
// Each field is a pointer to allow for nil values (unfiltered)
type UserFilter struct {
	FirstName *string
	LastName  *string
	Nickname  *string
	Country   *string
	Email     *string
}

// Pagination controls the result set size and offset for queries
type Pagination struct {
	Page  int
	Limit int
}

// UserRepository defines the interface for all user data operations
type UserRepository interface {
	// Create creates a new user in the database
	Create(ctx context.Context, user *models.User) error
	// Update modifies an existing user record
	Update(ctx context.Context, user *models.User) error
	// Delete removes a user by ID
	Delete(ctx context.Context, id string) error
	// FindAll retrieves users matching the filter with pagination
	FindAll(ctx context.Context, filter UserFilter, pagination Pagination) ([]models.User, error)
	// FindByID retrieves a single user by their ID
	FindByID(ctx context.Context, id string) (*models.User, error)
}

// userRepository is the GORM implementation of UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository instance with the given GORM DB
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create inserts a new user record into the database
func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// Update modifies an existing user record in the database
func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// Delete removes a user record by ID
func (r *userRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, "id = ?", id).Error
}

// FindAll retrieves a paginated list of users matching the filter criteria
func (r *userRepository) FindAll(ctx context.Context, filter UserFilter, pagination Pagination) ([]models.User, error) {
	var users []models.User
	query := r.db.WithContext(ctx).Model(&models.User{})

	if filter.FirstName != nil {
		query = query.Where("first_name ILIKE ?", "%"+*filter.FirstName+"%")
	}
	if filter.LastName != nil {
		query = query.Where("last_name ILIKE ?", "%"+*filter.LastName+"%")
	}
	if filter.Nickname != nil {
		query = query.Where("nickname = ?", *filter.Nickname)
	}
	if filter.Email != nil {
		query = query.Where("email = ?", *filter.Email)
	}
	if filter.Country != nil {
		query = query.Where("country = ?", *filter.Country)
	}

	if pagination.Page < 1 {
		pagination.Page = 1
	}
	if pagination.Limit < 1 {
		pagination.Limit = 10
	}
	offset := (pagination.Page - 1) * pagination.Limit
	err := query.Offset(offset).Limit(pagination.Limit).Find(&users).Error
	return users, err
}

// FindByID retrieves a single user by their unique ID
func (r *userRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error
	return &user, err
}
