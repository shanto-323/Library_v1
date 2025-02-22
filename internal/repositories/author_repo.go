package repository

import (
	"github.com/shanto-323/Library_v1.git/internal/models"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	Db *gorm.DB
}

func NewAuthorRepo(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{Db: db}
}

func (r *AuthorRepository) Create(author *models.Author) error {
	return r.Db.Create(author).Error
}

func (r *AuthorRepository) GetByID(id uint) (*models.Author, error) {
	var author models.Author
	if err := r.Db.Preload("Books").First(&author, id).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *AuthorRepository) GetAll() ([]models.Author, error) {
	var authors []models.Author
	if err := r.Db.Preload("Books").Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *AuthorRepository) Update(author *models.Author) error {
	return r.Db.Save(author).Error
}

func (r *AuthorRepository) Delete(id uint) error {
	return r.Db.Delete(&models.Author{}, id).Error
}
