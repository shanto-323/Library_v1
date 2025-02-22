package repository

import (
	"github.com/shanto-323/Library_v1.git/internal/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (b *BookRepository) Create(book *models.Book) error {
	return b.db.Create(book).Error
}

func (b *BookRepository) GetAll() ([]*models.Book, error) {
	var books []*models.Book
	err := b.db.Preload("Genres").Preload("Authors").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b *BookRepository) GetById(id uint) (*models.Book, error) {
	var book *models.Book
	if err := b.db.Preload("Genres").Preload("Authors").First(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (b *BookRepository) Update(book *models.Book) error {
	return b.db.Save(book).Error
}

func (b *BookRepository) Delete(id uint) error {
	return b.db.Delete(&models.Book{}, id).Error
}
