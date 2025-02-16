package storage

import (
	m "github.com/shanto-323/Library_v1.git/pkg/types"
)

type Books interface {
	GetAllBooks() ([]*m.Books, error)
	GetBooksByIsbm(int) (*m.Books, error)
	CreateBooks(*m.Books) (*m.Books, error)
	UpdateBooks(*m.Books) error
	DeleteBooks(int) error
}
