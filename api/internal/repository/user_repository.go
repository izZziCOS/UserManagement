package repository

import (
	"context"

	"github.com/izzzicos/UserManagement/api/internal/models"
	"gorm.io/gorm"
)

type UserFilter struct {
	FirstName *string
	LastName  *string
	Nickname  *string
	Country   *string
	Email     *string
}

type Pagination struct {
	Page  int
	Limit int
}

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, filter UserFilter, pagination Pagination) ([]models.User, error)
	FindByID(ctx context.Context, id string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, "id = ?", id).Error
}

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

func (r *userRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error
	return &user, err
}
