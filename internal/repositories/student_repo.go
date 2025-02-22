package repository

import (
	"github.com/shanto-323/Library_v1.git/internal/models"
	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (s *StudentRepository) Create(student *models.Student) error {
	return s.db.Create(student).Error
}

func (s *StudentRepository) GetAll() ([]*models.Student, error) {
	var students []*models.Student
	if err := s.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (s *StudentRepository) GetById(id uint) (*models.Student, error) {
	var student *models.Student
	if err := s.db.First(&student, id).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentRepository) Update(student *models.Student) error {
	return s.db.Save(student).Error
}

func (s *StudentRepository) Delete(id uint) error {
	return s.db.Delete(&models.Student{}, id).Error
}
