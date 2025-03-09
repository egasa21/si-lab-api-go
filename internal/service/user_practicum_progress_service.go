package service

import (
	"errors"
	"time"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/repository"
)

type UserPracticumProgressService interface {
	CreateProgress(progress *model.UserPracticumProgress) error
	GetProgress(userID, practicumID int) (*model.UserPracticumProgress, error)
	UpdateProgress(progress *model.UserPracticumProgress) error
	MarkAsCompleted(userID, practicumID int) error
	DeleteProgress(id int) error
}

type userPracticumProgressService struct {
	repo repository.UserPracticumProgressRepository
}

func NewUserPracticumProgressService(repo repository.UserPracticumProgressRepository) UserPracticumProgressService {
	return &userPracticumProgressService{repo: repo}
}

func (s *userPracticumProgressService) CreateProgress(progress *model.UserPracticumProgress) error {
	// Ensure progress is within valid range (0-100)
	if progress.Progress < 0 || progress.Progress > 100 {
		return errors.New("progress must be between 0 and 100")
	}

	// If progress is 100, set completion time
	if progress.Progress == 100 {
		now := time.Now()
		progress.CompletedAt = &now
	}

	return s.repo.CreateProgress(progress)
}

func (s *userPracticumProgressService) GetProgress(userID, practicumID int) (*model.UserPracticumProgress, error) {
	return s.repo.GetProgressByUserAndPracticum(userID, practicumID)
}

func (s *userPracticumProgressService) UpdateProgress(progress *model.UserPracticumProgress) error {
	// Ensure progress is within valid range
	if progress.Progress < 0 || progress.Progress > 100 {
		return errors.New("progress must be between 0 and 100")
	}

	// If progress is 100, mark as completed
	if progress.Progress == 100 {
		now := time.Now()
		progress.CompletedAt = &now
	} else {
		progress.CompletedAt = nil
	}

	return s.repo.UpdateProgress(progress)
}

func (s *userPracticumProgressService) MarkAsCompleted(userID, practicumID int) error {
	progress, err := s.repo.GetProgressByUserAndPracticum(userID, practicumID)
	if err != nil {
		return err
	}

	now := time.Now()
	progress.Progress = 100
	progress.CompletedAt = &now

	return s.repo.UpdateProgress(progress)
}

func (s *userPracticumProgressService) DeleteProgress(id int) error {
	return s.repo.DeleteProgress(id)
}
