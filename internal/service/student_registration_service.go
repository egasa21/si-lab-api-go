package service

import (
	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/repository"
)

type StudentRegistrationService interface {
	RegisterStudent(registration *model.StudentRegistration) error
	GetRegistrationsByStudentID(studentID int) ([]model.StudentRegistration, error)
	GetRegistrationsByPracticumID(practicumID int) ([]model.StudentRegistration, error)
	DeleteRegistration(id int) error
}

type studentRegistrationService struct {
	repo repository.StudentRegistrationRepository
}

func NewStudentRegistrationService(repo repository.StudentRegistrationRepository) StudentRegistrationService {
	return &studentRegistrationService{repo: repo}
}

func (s *studentRegistrationService) RegisterStudent(registration *model.StudentRegistration) error {
	return s.repo.RegisterStudent(registration)
}

func (s *studentRegistrationService) GetRegistrationsByStudentID(studentID int) ([]model.StudentRegistration, error) {
	return s.repo.GetRegistrationsByStudentID(studentID)
}

func (s *studentRegistrationService) GetRegistrationsByPracticumID(practicumID int) ([]model.StudentRegistration, error) {
	return s.repo.GetRegistrationsByPracticumID(practicumID)
}

func (s *studentRegistrationService) DeleteRegistration(id int) error {
	return s.repo.DeleteRegistration(id)
}
