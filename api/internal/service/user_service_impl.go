package service

import (
	"context"
	"time"

	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/izzzicos/UserManagement/api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req CreateUserRequest) (*models.UserResponse, error) {
    hashedPassword, err := hashPassword(req.Password)
    if err != nil {
        return nil, err
    }

    user := models.NewUser()
    user.FirstName = req.FirstName
    user.LastName = req.LastName
    user.Nickname = req.Nickname
    user.Password = hashedPassword
    user.Email = req.Email
    user.Country = req.Country

    if err := s.repo.Create(ctx, user); err != nil {
        return nil, err
    }

    response := user.ToResponse()
    return &response, nil
}

func (s *userService) UpdateUser(ctx context.Context, id string, req UpdateUserRequest) (*models.UserResponse, error) {
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

    response := user.ToResponse()
    return &response, nil
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}

func (s *userService) GetUsers(ctx context.Context, filter UserFilter, pagination Pagination) ([]models.UserResponse, error) {
    repoFilter := repository.UserFilter{
        Country: filter.Country,
    }
    repoPagination := repository.Pagination{
        Page:  pagination.Page,
        Limit: pagination.Limit,
    }

    users, err := s.repo.FindAll(ctx, repoFilter, repoPagination)
    if err != nil {
        return nil, err
    }

    responses := make([]models.UserResponse, len(users))
    for i, user := range users {
        responses[i] = user.ToResponse()
    }

    return responses, nil
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*models.UserResponse, error) {
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }

    response := user.ToResponse()
    return &response, nil
}

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}