package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/izzzicos/UserManagement/api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*models.User, error)
	UpdateUser(ctx context.Context, id string, req UpdateUserRequest) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
	GetUsers(ctx context.Context, filter repository.UserFilter, pagination repository.Pagination) ([]models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}

type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Nickname  string `json:"nickname" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
	Email     string `json:"email" validate:"required,email"`
	Country   string `json:"country" validate:"required"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req CreateUserRequest) (*models.User, error) {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:        uuid.New().String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Nickname:  req.Nickname,
		Password:  hashedPassword,
		Email:     req.Email,
		Country:   req.Country,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, id string, req UpdateUserRequest) (*models.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Country != "" {
		user.Country = req.Country
	}
	user.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *userService) GetUsers(ctx context.Context, filter repository.UserFilter, pagination repository.Pagination) ([]models.User, error) {
	return s.repo.FindAll(ctx, filter, pagination)
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return s.repo.FindByID(ctx, id)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}