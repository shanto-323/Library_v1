package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"username"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
