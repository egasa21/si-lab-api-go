package service

import (
	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/repository"
	"github.com/rs/zerolog/log"
)

type UserPracticumCheckpointService interface {
	CreateCheckpoint(checkpoint *model.UserPracticumCheckpoint) error
	GetCheckpointByUserAndPracticum(userID, practicumID int) (*model.UserPracticumCheckpoint, error)
	GetCheckpointByUser(userID int) ([]model.UserPracticumCheckpoint, error)
	UpdateCheckpoint(checkpoint *model.UserPracticumCheckpoint) error
	DeleteCheckpoint(id int) error
}

type userPracticumCheckpointService struct {
	repo repository.UserPracticumCheckpointRepository
}

func NewUserPracticumCheckpointService(repo repository.UserPracticumCheckpointRepository) UserPracticumCheckpointService {
	return &userPracticumCheckpointService{repo: repo}
}

// CreateCheckpoint creates a new user practicum checkpoint
func (s *userPracticumCheckpointService) CreateCheckpoint(checkpoint *model.UserPracticumCheckpoint) error {
	// You could add validation here, for example, checking if the user or practicum exists.
	err := s.repo.CreateCheckpoint(checkpoint)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create user practicum checkpoint")
		return err
	}
	return nil
}

// GetCheckpointByUserAndPracticum retrieves a checkpoint by user and practicum ID
func (s *userPracticumCheckpointService) GetCheckpointByUserAndPracticum(userID, practicumID int) (*model.UserPracticumCheckpoint, error) {
	checkpoint, err := s.repo.GetCheckpointByUserAndPracticum(userID, practicumID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user practicum checkpoint")
		return nil, err
	}
	if checkpoint == nil {
		log.Info().Msg("Checkpoint not found")
	}
	return checkpoint, nil
}

func (s *userPracticumCheckpointService) GetCheckpointByUser(userID int) ([]model.UserPracticumCheckpoint, error) {
	userCheckpoint, err := s.repo.GetCheckpointByUser(userID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user practicum checkpoint")
		return nil, err
	}
	if userCheckpoint == nil {
		log.Info().Msg("Checkpoint not found")
	}
	return userCheckpoint, nil
}

// UpdateCheckpoint updates an existing user practicum checkpoint
func (s *userPracticumCheckpointService) UpdateCheckpoint(checkpoint *model.UserPracticumCheckpoint) error {
	// You could add business logic here, like ensuring the checkpoint exists before updating it.
	err := s.repo.UpdateCheckpoint(checkpoint)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update user practicum checkpoint")
		return err
	}
	return nil
}

// DeleteCheckpoint deletes a user practicum checkpoint by ID
func (s *userPracticumCheckpointService) DeleteCheckpoint(id int) error {
	// You could add additional checks here, such as ensuring the checkpoint exists before deleting it.
	err := s.repo.DeleteCheckpoint(id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete user practicum checkpoint")
		return err
	}
	return nil
}
