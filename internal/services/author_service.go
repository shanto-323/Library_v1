package services

import (
	"github.com/shanto-323/Library_v1.git/internal/models"
	"github.com/shanto-323/Library_v1.git/internal/repositories"
)

type AuthorService struct {
	Repo *repository.AuthorRepository
}

func NewAuthorService(repo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{Repo: repo}
}

func (s *AuthorService) CreateAuthor(author *models.Author) error {
	return s.Repo.Create(author)
}

func (s *AuthorService) GetAuthorByID(id uint) (*models.Author, error) {
	return s.Repo.GetByID(id)
}

func (s *AuthorService) GetAllAuthors() ([]models.Author, error) {
	return s.Repo.GetAll()
}

func (s *AuthorService) UpdateAuthor(author *models.Author) error {
	return s.Repo.Update(author)
}

func (s *AuthorService) DeleteAuthor(id uint) error {
	return s.Repo.Delete(id)
}
