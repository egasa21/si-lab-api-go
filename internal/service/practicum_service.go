package service

import (
	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/repository"
)

type PracticumService interface {
	CreatePracticum(practicum *model.Practicum) error
	GetPracticumByID(id int) (*model.Practicum, error)
	GetAllPracticums(page, limit int) ([]model.Practicum, int, error)
}

type practicumService struct {
	repo repository.PracticumRepository
}

func NewPracticumService(repo repository.PracticumRepository) PracticumService {
	return &practicumService{repo: repo}
}

func (s *practicumService) CreatePracticum(practicum *model.Practicum) error {
	return s.repo.CreatePracticum(practicum)
}

func (s *practicumService) GetPracticumByID(id int) (*model.Practicum, error) {
	return s.repo.GetPracticumByID(id)
}

func (s *practicumService) GetAllPracticums(page, limit int) ([]model.Practicum, int, error) {
	return s.repo.GetAllPracticums(page, limit)
}
