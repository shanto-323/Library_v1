package cmd

import (
	"database/sql"
	"log"

	h "github.com/shanto-323/Library_v1.git/internal/handlers"
	s "github.com/shanto-323/Library_v1.git/internal/storage"
)

func main() {
	storage, err := s.NewPostgresDb(sql.Open)
	if err != nil {
		log.Fatal("Error creating database:", err)
	}

	if err = storage.CreateDb(); err != nil {
		log.Fatal("Error creating database table:", err)
	}

	userStore := s.NewUserStorage(storage)
	api := h.NewLibraryApi(":8080", userStore)
	api.Start()
}
