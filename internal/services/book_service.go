package services

import (
	"github.com/shanto-323/Library_v1.git/internal/models"
	"github.com/shanto-323/Library_v1.git/internal/repositories"
)

type BookService struct {
	Repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{Repo: repo}
}

func (s *BookService) CreateBook(book *models.Book) error {
	return s.Repo.Create(book)
}

func (s *BookService) GetBookByID(id uint) (*models.Book, error) {
	return s.Repo.GetById(uint(id))
}

func (s *BookService) GetAllBooks() ([]*models.Book, error) {
	return s.Repo.GetAll()
}

func (s *BookService) UpdateBook(book *models.Book) error {
	return s.Repo.Update(book)
}

func (s *BookService) DeleteBook(id uint) error {
	return s.Repo.Delete(id)
}
