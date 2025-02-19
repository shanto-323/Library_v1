package models

import "time"

type Genre struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique;not null"`
	Books []Book `gorm:"many2many:book_genres;constraint:OnDelete:CASCADE;"`
}

type Author struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique;not null"`
	BirthYear   int    `gorm:"check:birth_year > 0"`
	Nationality string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	Books       []Book    `gorm:"many2many:book_authors;constraint:OnDelete:CASCADE;"`
}

type Book struct {
	ID              uint      `gorm:"primaryKey"`
	ISBN            string    `gorm:"unique;not null;size:30"`
	Title           string    `gorm:"not null;size:200"`
	PublishedYear   int       `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	TotalCopies     int       `gorm:"not null;check:total_copies >= 0"`
	AvailableCopies int       `gorm:"not null;check:available_copies >= 0"`
	Genres          []Genre   `gorm:"many2many:book_genres;constraint:OnDelete:CASCADE;"`
	Authors         []Author  `gorm:"many2many:book_authors;constraint:OnDelete:CASCADE;"`
}

type Student struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique"`
	Phone     string    `gorm:"unique;not null"`
	IsActive  bool      `gorm:"default:true"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type BorrowedBook struct {
	ID        uint      `gorm:"primaryKey"`
	StudentID uint      `gorm:"not null"`
	BookISBN  string    `gorm:"not null"`
	TakenAt   time.Time `gorm:"autoCreateTime"`
	ReturnAt  int64     `gorm:"not null"`
	Returned  bool      `gorm:"default:false"`
	Student   Student   `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`
	Book      Book      `gorm:"foreignKey:BookISBN;references:ISBN;constraint:OnDelete:CASCADE;"`
}
