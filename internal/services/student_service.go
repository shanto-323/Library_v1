package services

import (
	"github.com/shanto-323/Library_v1.git/internal/models"
	"github.com/shanto-323/Library_v1.git/internal/repository"
)

type StudentService struct {
	Repo *repository.StudentRepo
}

func NewStudentService(repo repository.StudentRepo) *StudentService {
	return &StudentService{
		Repo: &repo,
	}
}

func (s *StudentService) CreateStudent(student *models.Student) error {
	return s.Repo.CreateStudent(student)
}

func (s *StudentService) GetAllStudent() ([]*models.Student, error) {
	return s.Repo.GetAllStudent()
}

func (s *StudentService) GetStudentById(id uint) (*models.Student, error) {
	return s.Repo.GetStudentById(id)
}

func (s *StudentService) UpdateStudent(student *models.Student) error {
	return s.Repo.Update(student)
}

func (s *StudentService) DeleteStudent(id uint) error {
	return s.Repo.Delete(id)
}
