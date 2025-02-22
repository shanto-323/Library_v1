package services

import (
	"github.com/shanto-323/Library_v1.git/internal/models"
	"github.com/shanto-323/Library_v1.git/internal/repositories"
)

type StudentService struct {
	Repo *repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) *StudentService {
	return &StudentService{
		Repo: &repo,
	}
}

func (s *StudentService) CreateStudent(student *models.Student) error {
	return s.Repo.Create(student)
}

func (s *StudentService) GetAllStudent() ([]*models.Student, error) {
	return s.Repo.GetAll()
}

func (s *StudentService) GetStudentById(id uint) (*models.Student, error) {
	return s.Repo.GetById(id)
}

func (s *StudentService) UpdateStudent(student *models.Student) error {
	return s.Repo.Update(student)
}

func (s *StudentService) DeleteStudent(id uint) error {
	return s.Repo.Delete(id)
}
