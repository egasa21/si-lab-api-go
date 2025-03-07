package repository

import (
	"database/sql"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/rs/zerolog/log"
)

type StudentClassEnrollmentRepository interface {
	EnrollStudent(classID, studentID int) error
	GetEnrollmentsByStudentID(studentID int) ([]model.StudentClassEnrollment, error)
	GetEnrollmentsByClassID(classID int) ([]model.StudentClassEnrollment, error)
	DeleteEnrollment(enrollmentID int) error
}

type studentClassEnrollmentRepository struct {
	db *sql.DB
}

// NewStudentClassEnrollmentRepository creates a new instance of the repository
func NewStudentClassEnrollmentRepository(db *sql.DB) StudentClassEnrollmentRepository {
	return &studentClassEnrollmentRepository{db: db}
}

// EnrollStudent enrolls a student in a class
func (r *studentClassEnrollmentRepository) EnrollStudent(classID, studentID int) error {
	query := `
		INSERT INTO student_class_enrollment (class_id, student_id, created_at, updated_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err := r.db.Exec(query, classID, studentID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to enroll student in class")
		return err
	}
	return nil
}

// GetEnrollmentsByStudentID retrieves all class enrollments for a specific student
func (r *studentClassEnrollmentRepository) GetEnrollmentsByStudentID(studentID int) ([]model.StudentClassEnrollment, error) {
	query := `
		SELECT id, class_id, student_id, created_at, updated_at
		FROM student_class_enrollment
		WHERE student_id = $1
	`
	rows, err := r.db.Query(query, studentID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve enrollments for student")
		return nil, err
	}
	defer rows.Close()

	var enrollments []model.StudentClassEnrollment
	for rows.Next() {
		var enrollment model.StudentClassEnrollment
		err := rows.Scan(&enrollment.ID, &enrollment.ClassID, &enrollment.StudentID, &enrollment.CreatedAt, &enrollment.UpdatedAt)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning enrollment record")
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}

	return enrollments, nil
}

// GetEnrollmentsByClassID retrieves all students enrolled in a specific class
func (r *studentClassEnrollmentRepository) GetEnrollmentsByClassID(classID int) ([]model.StudentClassEnrollment, error) {
	query := `
		SELECT id, class_id, student_id, created_at, updated_at
		FROM student_class_enrollment
		WHERE class_id = $1
	`
	rows, err := r.db.Query(query, classID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve enrollments for class")
		return nil, err
	}
	defer rows.Close()

	var enrollments []model.StudentClassEnrollment
	for rows.Next() {
		var enrollment model.StudentClassEnrollment
		err := rows.Scan(&enrollment.ID, &enrollment.ClassID, &enrollment.StudentID, &enrollment.CreatedAt, &enrollment.UpdatedAt)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning enrollment record")
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}

	return enrollments, nil
}

// DeleteEnrollment removes a student from a class
func (r *studentClassEnrollmentRepository) DeleteEnrollment(enrollmentID int) error {
	query := `
		DELETE FROM student_class_enrollment WHERE id = $1
	`
	_, err := r.db.Exec(query, enrollmentID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete enrollment")
		return err
	}
	return nil
}
