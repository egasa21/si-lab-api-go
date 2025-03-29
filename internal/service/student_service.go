package service

import (
	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/repository"
)

type StudentService interface {
	GetAllStudents(page, limit int) ([]model.Student, int, error)
	GetStudentByID(id int) (*model.Student, error)
	GetStudentByUserID(id int) (*model.Student, error)
	CreateStudent(student *model.Student) (int, error)
	GetStudentByStudentID(student_id_number string) (*model.Student, error)
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo: repo}
}

func (s *studentService) GetAllStudents(page, limit int) ([]model.Student, int, error) {
	return s.repo.GetAllStudents(page, limit)
}

func (s *studentService) GetStudentByID(id int) (*model.Student, error) {
	return s.repo.GetStudentByID(id)
}

func (s *studentService) GetStudentByStudentID(student_id_number string) (*model.Student, error) {
	return s.repo.GetStudentByStudentID(student_id_number)
}

func (s *studentService) GetStudentByUserID(id int) (*model.Student, error) {
	return s.repo.GetStudentByUserID(id)
}

func (s *studentService) CreateStudent(student *model.Student) (int, error) {
	return s.repo.CreateStudent(student)
}
