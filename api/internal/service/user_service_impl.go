package service

import (
	"context"
	"log"
	"time"

	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/izzzicos/UserManagement/api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo     repository.UserRepository
	notifier Notifier
}

func NewUserService(repo repository.UserRepository, notifier Notifier) UserService {
	return &userService{
		repo:     repo,
		notifier: notifier,
	}
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

	go func(u *models.User) {
		notifyCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.notifier.NotifyUserCreated(notifyCtx, u); err != nil {
			log.Printf("Failed to send notification for user %s: %v", u.ID, err)
		} else {
			log.Printf("Successfully notified about user creation: %s", u.ID)
		}
	}(user) // Pass user as parameter to avoid race conditions

	response := user.ToResponse()
	return &response, nil
}

func (s *userService) UpdateUser(ctx context.Context, id string, req UpdateUserRequest) (*models.UserResponse, error) {
	oldUser, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	updatedUser := *oldUser

	if req.FirstName != "" {
		updatedUser.FirstName = req.FirstName
	}
	if req.LastName != "" {
		updatedUser.LastName = req.LastName
	}
	if req.Nickname != "" {
		updatedUser.Nickname = req.Nickname
	}
	if req.Email != "" {
		updatedUser.Email = req.Email
	}
	if req.Country != "" {
		updatedUser.Country = req.Country
	}
	updatedUser.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, &updatedUser); err != nil {
		return nil, err
	}

	go func(oldU, newU models.User) {
		notifyCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.notifier.NotifyUserUpdated(notifyCtx, &oldU, &newU); err != nil {
			log.Printf("Failed to send update notification for user %s: %v", id, err)
		} else {
			log.Printf("Successfully notified about user update: %s", id)
		}
	}(*oldUser, updatedUser)

	response := updatedUser.ToResponse()
	return &response, nil
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	go func() {
		notifyCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.notifier.NotifyUserDeleted(notifyCtx, id); err != nil {
			log.Printf("Failed to send notification for user deletion %s: %v", user.ID, err)
		} else {
			log.Printf("Successfully notified about user deletion: %s", user.ID)
		}
	}()

	return nil
}

func (s *userService) GetUsers(ctx context.Context, filter UserFilter, pagination Pagination) ([]models.UserResponse, error) {
	repoFilter := repository.UserFilter{
		FirstName: filter.FirstName,
		LastName:  filter.LastName,
		Nickname:  filter.Nickname,
		Email:     filter.Email,
		Country:   filter.Country,
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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
