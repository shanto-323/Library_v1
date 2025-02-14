package storage

import (
	m "github.com/shanto-323/Library_v1.git/models"
)

type UserData interface {
	GetAllUser() ([]*m.User, error)
	GetUserById(int) (*m.User, error)
	CreateUser(*m.User) (*m.User, error)
	UpdateUser(*m.User) error
	DeleteUser(int) error
}

type Books interface {
	GetAllBooks() ([]*m.Books, error)
	GetBooksByIsbm(int) (*m.Books, error)
	CreateBooks(*m.Books) (*m.Books, error)
	UpdateBooks(*m.Books) error
	DeleteBooks(int) error
}

//More req ..
