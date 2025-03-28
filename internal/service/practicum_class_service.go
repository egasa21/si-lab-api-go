package service

import (
	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/repository"
)

type PracticumClassService interface {
	CreateClass(class *model.PracticumClass) error
	GetClassByID(id int) (*model.PracticumClass, error)
	GetClassByIDs(ids []int) ([]model.PracticumClass, error)
	GetClassesByPracticumID(practicumID int) ([]model.PracticumClass, error)
	UpdateClass(class *model.PracticumClass) error
	DeleteClass(id int) error
}

type practicumClassService struct {
	repo repository.PracticumClassRepository
}

func NewPracticumClassService(repo repository.PracticumClassRepository) PracticumClassService {
	return &practicumClassService{repo: repo}
}

func (s *practicumClassService) CreateClass(class *model.PracticumClass) error {
	return s.repo.CreateClass(class)
}

func (s *practicumClassService) GetClassByID(id int) (*model.PracticumClass, error) {
	return s.repo.GetClassByID(id)
}

func (s *practicumClassService) GetClassesByPracticumID(practicumID int) ([]model.PracticumClass, error) {
	return s.repo.GetClassesByPracticumID(practicumID)
}

func (s *practicumClassService) UpdateClass(class *model.PracticumClass) error {
	return s.repo.UpdateClass(class)
}

func (s *practicumClassService) DeleteClass(id int) error {
	return s.repo.DeleteClass(id)
}

func (s *practicumClassService) GetClassByIDs(ids []int) ([]model.PracticumClass, error){
	return s.repo.GetClassByIDs(ids)
}
