package service

import (
	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/repository"
)

type PracticumModuleContentService interface {
	CreateContent(content *model.PracticumModuleContent) (*model.PracticumModuleContent, error)
	GetContentByID(id int) (*model.PracticumModuleContent, error)
	GetContentByIDs(ids []int) ([]model.PracticumModuleContent, error)
	GetContentsByModuleID(moduleID, page, limit int) ([]model.PracticumModuleContent, int, error)
	UpdateContentByID(id int, updatedContent *model.PracticumModuleContent) error
	DeleteContentByID(id int) error
}

type practicumModuleContentService struct {
	repo repository.PracticumModuleContentRepository
}

func NewPracticumModuleContentService(repo repository.PracticumModuleContentRepository) PracticumModuleContentService {
	return &practicumModuleContentService{repo: repo}
}

func (s *practicumModuleContentService) CreateContent(content *model.PracticumModuleContent) (*model.PracticumModuleContent, error) {
	return s.repo.CreateContent(content)
}

func (s *practicumModuleContentService) GetContentByID(id int) (*model.PracticumModuleContent, error) {
	return s.repo.GetContentByID(id)
}

func (s *practicumModuleContentService) GetContentsByModuleID(moduleID, page, limit int) ([]model.PracticumModuleContent, int, error) {
	return s.repo.GetContentsByModuleID(moduleID, page, limit)
}

func (s *practicumModuleContentService) GetContentByIDs(ids []int) ([]model.PracticumModuleContent, error) {
	return s.repo.GetContentByIDs(ids)
}

func (s *practicumModuleContentService) UpdateContentByID(id int, updatedContent *model.PracticumModuleContent) error {
	return s.repo.UpdateContentByID(id, updatedContent)
}

func (s *practicumModuleContentService) DeleteContentByID(id int) error {
	return s.repo.DeleteContentByID(id)
}
