package models

import "time"

type Books struct {
	Id              int       `json:"id"`
	Isbn            int       `json:"isbn"`
	Title           string    `json:"title"`
	Genre           string    `json:"genre"`
	PublishedYear   int       `json:"published_year"`
	CreatedAt       time.Time `json:"created_at"`
	AvailableCopies int       `json:"available_copies"`
}

type BorrowedBooks struct {
	Id       int       `json:"id"`
	UserId   int       `json:"user_id"`
	BookIsbn int       `json:"book_isbn"`
	TakenAt  time.Time `json:"taken_at"`
	ReturnAt time.Time `json:"return_at"`
	Returned bool      `json:"returned"`
}
