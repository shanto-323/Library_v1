package services

import (
	"github.com/shanto-323/Library_v1.git/internal/models"
	"github.com/shanto-323/Library_v1.git/internal/repositories"
)

type BorrowedBookService struct {
	Repo *repository.BorrowedBookRepository
}

func NewBorrowedBookService(repo *repository.BorrowedBookRepository) *BorrowedBookService {
	return &BorrowedBookService{Repo: repo}
}

func (s *BorrowedBookService) BorrowBook(borrowedBook *models.BorrowedBook) error {
	return s.Repo.Create(borrowedBook)
}

func (s *BorrowedBookService) ReturnBook(id uint) error {
	return s.Repo.Delete(uint(id))
}

func (s *BorrowedBookService) GetAllBorrowedBooks() ([]models.BorrowedBook, error) {
	return s.Repo.GetAll()
}

func (s *BorrowedBookService) GetBorrowedBookByID(id uint) (*models.BorrowedBook, error) {
	return s.Repo.GetByID(id)
}

func (s *BorrowedBookService) UpdateBorrowedBookByID(borrowedBook *models.BorrowedBook) error {
	return s.Repo.Update(borrowedBook)
}
