package service

import (
	"errors"
	"fmt"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/repository"
	"github.com/egasa21/si-lab-api-go/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user *model.User, roles []string) error
	Login(email, password string) (*auth.TokenDetails, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(user *model.User, roles []string) error {
	// Check if the email is already registered
	existingUser, _ := s.repo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email already registered")
	}

	if user.Password == "" {
		return errors.New("password cannot be empty")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Create the user and get the ID
	if err := s.repo.Register(user); err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	// Assign default role of "student" if no roles are provided
	if len(roles) == 0 {
		roles = []string{"student"}
	}

	// Assign roles to the user
	for _, role := range roles {
		// Get the role ID from the role name
		roleID, err := s.getRoleIDByName(role)
		if err != nil {
			return fmt.Errorf("failed to get role ID for role %s: %w", role, err)
		}

		// Insert the role into the user_roles table
		if err := s.repo.AddRoleToUser(user.IDUser, roleID); err != nil {
			return fmt.Errorf("failed to assign role %s to user: %w", role, err)
		}
	}

	return nil
}

func (s *authService) Login(email, password string) (*auth.TokenDetails, error) {
	// Get user by email
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println("Stored Password:", user.Password)
		fmt.Println("Input Password:", password)
		fmt.Printf("Password comparison failed: %v\n", err)
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT (access and refresh tokens)
	tokens, err := auth.GenerateJWT(user.IDUser, user.Roles)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	return tokens, nil
}


func (s *authService) getRoleIDByName(name string) (int, error) {
	roleMapping := map[string]int{
		"admin":                1,
		"student":              2,
		"lecturer":             3,
		"laboratory_assistant": 4,
	}
	id, exists := roleMapping[name]
	if !exists {
		return 0, errors.New("invalid role")
	}
	return id, nil
}
