package repository

import (
	"github.com/shanto-323/Library_v1.git/internal/models"
	"gorm.io/gorm"
)

type GenreRepository struct {
	db *gorm.DB
}

func NewGenreRepo(db *gorm.DB) *GenreRepository {
	return &GenreRepository{
		db: db,
	}
}

func (g *GenreRepository) Create(genre *models.Genre) error {
	return g.db.Create(genre).Error
}

func (g *GenreRepository) GetAll() ([]*models.Genre, error) {
	var genres []*models.Genre
	if err := g.db.Find(&genres).Error; err != nil {
		return nil, err
	}
	return genres, nil
}

func (g *GenreRepository) Update(genre *models.Genre) error {
	return g.db.Save(genre).Error
}

func (g *GenreRepository) Delete(id uint) error {
	return g.db.Delete(&models.Genre{}, id).Error
}
