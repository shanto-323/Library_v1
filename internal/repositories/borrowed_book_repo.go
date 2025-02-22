package repository

import (
	"github.com/shanto-323/Library_v1.git/internal/models"
	"gorm.io/gorm"
)

type BorrowedBookRepository struct {
	Db *gorm.DB
}

func NewBorrowedBookRepository(db *gorm.DB) *BorrowedBookRepository {
	return &BorrowedBookRepository{Db: db}
}

func (r *BorrowedBookRepository) Create(borrowedBook *models.BorrowedBook) error {
	return r.Db.Create(borrowedBook).Error
}

func (r *BorrowedBookRepository) GetByID(id uint) (*models.BorrowedBook, error) {
	var borrowedBook models.BorrowedBook
	if err := r.Db.Preload("Student").Preload("Book").First(&borrowedBook, id).Error; err != nil {
		return nil, err
	}
	return &borrowedBook, nil
}

func (r *BorrowedBookRepository) GetAll() ([]models.BorrowedBook, error) {
	var borrowedBooks []models.BorrowedBook
	if err := r.Db.Preload("Student").Preload("Book").Find(&borrowedBooks).Error; err != nil {
		return nil, err
	}
	return borrowedBooks, nil
}

func (r *BorrowedBookRepository) Update(borrowedBook *models.BorrowedBook) error {
	return r.Db.Save(borrowedBook).Error
}

func (r *BorrowedBookRepository) Delete(id uint) error {
	return r.Db.Delete(&models.BorrowedBook{}, id).Error
}
