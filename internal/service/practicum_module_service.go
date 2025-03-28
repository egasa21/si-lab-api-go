package service

import (
	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/repository"
)

type PracticumModuleService interface {
	CreateModule(module *model.PracticumModule) error
	GetModuleByID(id int) (*model.PracticumModule, error)
	GetModuleByIDs(ids []int) ([]model.PracticumModule, error)
	GetModulesByPracticumID(practicumID, page, limit int) ([]model.PracticumModule, int, error)
}

type practicumModuleService struct {
	repo repository.PracticumModuleRepository
}

func NewPracticumModuleService(repo repository.PracticumModuleRepository) PracticumModuleService {
	return &practicumModuleService{repo: repo}
}

func (s *practicumModuleService) CreateModule(module *model.PracticumModule) error {
	return s.repo.CreateModule(module)
}

func (s *practicumModuleService) GetModuleByID(id int) (*model.PracticumModule, error) {
	return s.repo.GetModuleByID(id)
}

func (s *practicumModuleService) GetModulesByPracticumID(practicumID, page, limit int) ([]model.PracticumModule, int, error) {
	return s.repo.GetModulesByPracticumID(practicumID, page, limit)
}

func (s *practicumModuleService) GetModuleByIDs(ids []int) ([]model.PracticumModule, error) {
	return s.repo.GetModuleByIDs(ids)
}
