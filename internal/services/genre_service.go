package services

import (
	"github.com/shanto-323/Library_v1.git/internal/models"
	"github.com/shanto-323/Library_v1.git/internal/repositories"
)

type GenreService struct {
	genreRepo *repository.GenreRepository
}

func NewGenreService(repo *repository.GenreRepository) *GenreService {
	return &GenreService{genreRepo: repo}
}

func (s *GenreService) CreateGenre(genre *models.Genre) error {
	return s.genreRepo.Create(genre)
}

func (s *GenreService) GetAllGenres() ([]*models.Genre, error) {
	return s.genreRepo.GetAll()
}

func (s *GenreService) UpdateGenre(genre *models.Genre) error {
	return s.genreRepo.Update(genre)
}

func (s *GenreService) DeleteGenre(id uint) error {
	return s.genreRepo.Delete(id)
}
