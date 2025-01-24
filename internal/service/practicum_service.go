package service

import (
	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/repository"
)

type PracticumService interface {
	CreatePracticum(practicum *model.Practicum) error
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
