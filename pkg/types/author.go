package models

type Author struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	BirthYear   int    `json:"birth_year"`
	Nationality string `json:"nationality"`
	CreatedAt   string `json:"created_at"`
}
