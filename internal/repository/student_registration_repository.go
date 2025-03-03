package repository

import (
	"database/sql"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/rs/zerolog/log"
)

type StudentRegistrationRepository interface {
	RegisterStudent(registration *model.StudentRegistration) error
	GetRegistrationsByStudentID(studentID int) ([]model.StudentRegistration, error)
	GetRegistrationsByPracticumID(practicumID int) ([]model.StudentRegistration, error)
	DeleteRegistration(id int) error
}

type studentRegistrationRepository struct {
	db *sql.DB
}

func NewStudentRegistrationRepository(db *sql.DB) StudentRegistrationRepository {
	return &studentRegistrationRepository{db: db}
}

func (r *studentRegistrationRepository) RegisterStudent(registration *model.StudentRegistration) error {
	query := `
		INSERT INTO student_registration (student_id, practicum_id)
		VALUES ($1, $2)
		RETURNING id_student_registration, created_at, updated_at
	`
	err := r.db.QueryRow(query, registration.StudentID, registration.PracticumID).
		Scan(&registration.IDStudentRegistration, &registration.CreatedAt, &registration.UpdatedAt)
	if err != nil {
		log.Error().Err(err).Msg("Failed to register student")
		return err
	}
	return nil
}

func (r *studentRegistrationRepository) GetRegistrationsByStudentID(studentID int) ([]model.StudentRegistration, error) {
	query := `
		SELECT id_student_registration, student_id, practicum_id, created_at, updated_at
		FROM student_registration
		WHERE student_id = $1
	`
	rows, err := r.db.Query(query, studentID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch registrations by student ID")
		return nil, err
	}
	defer rows.Close()

	var registrations []model.StudentRegistration
	for rows.Next() {
		var reg model.StudentRegistration
		if err := rows.Scan(&reg.IDStudentRegistration, &reg.StudentID, &reg.PracticumID, &reg.CreatedAt, &reg.UpdatedAt); err != nil {
			return nil, err
		}
		registrations = append(registrations, reg)
	}

	return registrations, nil
}

func (r *studentRegistrationRepository) GetRegistrationsByPracticumID(practicumID int) ([]model.StudentRegistration, error) {
	query := `
		SELECT id_student_registration, student_id, practicum_id, created_at, updated_at
		FROM student_registration
		WHERE practicum_id = $1
	`
	rows, err := r.db.Query(query, practicumID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch registrations by practicum ID")
		return nil, err
	}
	defer rows.Close()

	var registrations []model.StudentRegistration
	for rows.Next() {
		var reg model.StudentRegistration
		if err := rows.Scan(&reg.IDStudentRegistration, &reg.StudentID, &reg.PracticumID, &reg.CreatedAt, &reg.UpdatedAt); err != nil {
			return nil, err
		}
		registrations = append(registrations, reg)
	}

	return registrations, nil
}

func (r *studentRegistrationRepository) DeleteRegistration(id int) error {
	query := "DELETE FROM student_registration WHERE id_student_registration = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete student registration")
		return err
	}
	return nil
}
