package services

import (
	"errors"

	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/models"
	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService defines methods for user business logic
type UserService interface {
	CreateUser(req *models.CreateUserRequest) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	GetUserByID(id int) (*models.User, error)
	UpdateUser(id int, req *models.UpdateUserRequest) (*models.User, error)
	DeleteUser(id int) error
}

// userService implements UserService
type userService struct {
	userRepo repositories.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser handles user creation business logic
func (s *userService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	// Check if email already exists
	_, err := s.userRepo.GetByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Save to database
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetAllUsers() ([]*models.User, error) {
	return s.userRepo.GetAll()
}

func (s *userService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.GetByID(uint(id))
}

func (s *userService) UpdateUser(id int, req *models.UpdateUserRequest) (*models.User, error) {
	// Fetch existing user
	user, err := s.userRepo.GetByID(uint(id))
	if err != nil {
		return nil, err
	}

	// Update fields only if they are not empty
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		// Check if new email already exists
		existingUser, err := s.userRepo.GetByEmail(req.Email)
		if err == nil && existingUser.ID != id {
			return nil, errors.New("email already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		user.Email = req.Email
	}

	// Save updates
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id int) error {
	// Check if user exists first
	_, err := s.userRepo.GetByID(uint(id))
	if err != nil {
		return err
	}

	// Delete the user
	return s.userRepo.Delete(uint(id))
}
