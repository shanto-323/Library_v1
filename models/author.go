package models

import "time"

type Author struct {
	Id          int       `json:"id"`
	Username    string    `json:"username"`
	BirthYear   int       `json:"birth_year"`
	Nationality string    `json:"nationality"`
	CreatedAt   time.Time `json:"created_at"`
}
