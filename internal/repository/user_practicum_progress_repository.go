package repository

import (
	"database/sql"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/rs/zerolog/log"
)

type UserPracticumProgressRepository interface {
	CreateProgress(progress *model.UserPracticumProgress) error
	GetProgressByUserAndPracticum(userID, practicumID int) (*model.UserPracticumProgress, error)
	UpdateProgress(progress *model.UserPracticumProgress) error
	DeleteProgress(id int) error
}

type userPracticumProgressRepository struct {
	db *sql.DB
}

func NewUserPracticumProgressRepository(db *sql.DB) UserPracticumProgressRepository {
	return &userPracticumProgressRepository{db: db}
}

func (r *userPracticumProgressRepository) CreateProgress(progress *model.UserPracticumProgress) error {
	query := `
		INSERT INTO user_practicum_progress (id_user, id_practicum, progress, completed_at, last_updated_at)
		VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP)
		RETURNING id
	`
	err := r.db.QueryRow(query, progress.UserID, progress.PracticumID, progress.Progress, progress.CompletedAt).
		Scan(&progress.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create user practicum progress")
		return err
	}
	return nil
}

func (r *userPracticumProgressRepository) GetProgressByUserAndPracticum(userID, practicumID int) (*model.UserPracticumProgress, error) {
	query := `
		SELECT id, id_user, id_practicum, progress, completed_at, last_updated_at
		FROM user_practicum_progress
		WHERE id_user = $1 AND id_practicum = $2
	`
	var progress model.UserPracticumProgress
	err := r.db.QueryRow(query, userID, practicumID).
		Scan(&progress.ID, &progress.UserID, &progress.PracticumID, &progress.Progress, &progress.CompletedAt, &progress.LastUpdated)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch user practicum progress")
		return nil, err
	}
	return &progress, nil
}

func (r *userPracticumProgressRepository) UpdateProgress(progress *model.UserPracticumProgress) error {
	query := `
		UPDATE user_practicum_progress
		SET progress = $1, completed_at = $2, last_updated_at = CURRENT_TIMESTAMP
		WHERE id = $3
	`
	_, err := r.db.Exec(query, progress.Progress, progress.CompletedAt, progress.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update user practicum progress")
		return err
	}
	return nil
}

func (r *userPracticumProgressRepository) DeleteProgress(id int) error {
	query := `DELETE FROM user_practicum_progress WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete user practicum progress")
		return err
	}
	return nil
}
