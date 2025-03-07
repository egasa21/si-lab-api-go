package service

import (
	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/repository"
)

type StudentClassEnrollmentService interface {
	EnrollStudent(classID, studentID int) error
	GetEnrollmentsByStudentID(studentID int) ([]model.StudentClassEnrollment, error)
	GetEnrollmentsByClassID(classID int) ([]model.StudentClassEnrollment, error)
	UnenrollStudent(enrollmentID int) error
}

type studentClassEnrollmentService struct {
	repo repository.StudentClassEnrollmentRepository
}

// NewStudentClassEnrollmentService creates a new instance of the service
func NewStudentClassEnrollmentService(repo repository.StudentClassEnrollmentRepository) StudentClassEnrollmentService {
	return &studentClassEnrollmentService{repo: repo}
}

// EnrollStudent enrolls a student in a class
func (s *studentClassEnrollmentService) EnrollStudent(classID, studentID int) error {
	return s.repo.EnrollStudent(classID, studentID)
}

// GetEnrollmentsByStudentID retrieves all class enrollments for a student
func (s *studentClassEnrollmentService) GetEnrollmentsByStudentID(studentID int) ([]model.StudentClassEnrollment, error) {
	return s.repo.GetEnrollmentsByStudentID(studentID)
}

// GetEnrollmentsByClassID retrieves all students in a specific class
func (s *studentClassEnrollmentService) GetEnrollmentsByClassID(classID int) ([]model.StudentClassEnrollment, error) {
	return s.repo.GetEnrollmentsByClassID(classID)
}

// UnenrollStudent removes a student from a class
func (s *studentClassEnrollmentService) UnenrollStudent(enrollmentID int) error {
	return s.repo.DeleteEnrollment(enrollmentID)
}
