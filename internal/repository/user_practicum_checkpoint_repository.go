package repository

import (
	"database/sql"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/rs/zerolog/log"
)

type UserPracticumCheckpointRepository interface {
	CreateCheckpoint(checkpoint *model.UserPracticumCheckpoint) error
	GetCheckpointByUserAndPracticum(userID, practicumID int) (*model.UserPracticumCheckpoint, error)
	UpdateCheckpoint(checkpoint *model.UserPracticumCheckpoint) error
	DeleteCheckpoint(id int) error
}

type userPracticumCheckpointRepository struct {
	db *sql.DB
}

func NewUserPracticumCheckpointRepository(db *sql.DB) UserPracticumCheckpointRepository {
	return &userPracticumCheckpointRepository{db: db}
}

// CreateCheckpoint inserts a new checkpoint into the database
func (r *userPracticumCheckpointRepository) CreateCheckpoint(checkpoint *model.UserPracticumCheckpoint) error {
	query := `
		INSERT INTO user_practicum_checkpoint (id_user, id_practicum, id_module, id_content)
		VALUES ($1, $2, $3, $4)
		RETURNING id, updated_at
	`
	err := r.db.QueryRow(query, checkpoint.UserID, checkpoint.PracticumID, checkpoint.ModuleID, checkpoint.ContentID).
		Scan(&checkpoint.ID, &checkpoint.UpdatedAt)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create user practicum checkpoint")
		return err
	}
	return nil
}

// GetCheckpointByUserAndPracticum fetches the checkpoint for a specific user and practicum
func (r *userPracticumCheckpointRepository) GetCheckpointByUserAndPracticum(userID, practicumID int) (*model.UserPracticumCheckpoint, error) {
	query := `
		SELECT id, id_user, id_practicum, id_module, id_content, updated_at
		FROM user_practicum_checkpoint
		WHERE id_user = $1 AND id_practicum = $2
	`
	row := r.db.QueryRow(query, userID, practicumID)

	var checkpoint model.UserPracticumCheckpoint
	err := row.Scan(&checkpoint.ID, &checkpoint.UserID, &checkpoint.PracticumID, &checkpoint.ModuleID, &checkpoint.ContentID, &checkpoint.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No checkpoint found
		}
		log.Error().Err(err).Msg("Failed to fetch user practicum checkpoint")
		return nil, err
	}
	return &checkpoint, nil
}

// UpdateCheckpoint updates an existing checkpoint record
func (r *userPracticumCheckpointRepository) UpdateCheckpoint(checkpoint *model.UserPracticumCheckpoint) error {
	query := `
		UPDATE user_practicum_checkpoint
		SET id_module = $1, id_content = $2, updated_at = CURRENT_TIMESTAMP
		WHERE id = $3
	`
	_, err := r.db.Exec(query, checkpoint.ModuleID, checkpoint.ContentID, checkpoint.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update user practicum checkpoint")
		return err
	}
	return nil
}

// DeleteCheckpoint deletes a checkpoint by ID
func (r *userPracticumCheckpointRepository) DeleteCheckpoint(id int) error {
	query := `DELETE FROM user_practicum_checkpoint WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete user practicum checkpoint")
		return err
	}
	return nil
}
