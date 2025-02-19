package main

import (
	"log"

	"github.com/shanto-323/Library_v1.git/config"
	"github.com/shanto-323/Library_v1.git/internal/handlers"
	"github.com/shanto-323/Library_v1.git/internal/models"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer func() {
		sqldb, _ := db.DB()
		sqldb.Close()
	}()

	err = db.AutoMigrate(
		&models.Genre{},
		&models.Author{},
		&models.Book{},
		&models.Student{},
		&models.BorrowedBook{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	api := handlers.NewLibraryApi(":8080", db)
	api.Start()
}
