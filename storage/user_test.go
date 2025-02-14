package storage

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shanto-323/Library_v1.git/models"
)

func TestGetAllData(t *testing.T) {
	db, mock, err := sqlmock.New()
	pgDb := &PostgresDb{Db: db}
	if err != nil {
		t.Fatalf("failed to create new db %v", err)
	}
	defer db.Close()

	createdAt := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "is_active", "created_at"}).
		AddRow(1, "John Doe1", "john.doe@example.com", "1234567890", true, createdAt).
		AddRow(2, "Jane Doe2", "jane.doe@example.com", "0987654321", false, createdAt)
	mock.ExpectQuery(`SELECT id, name, email, phone, is_active, created_at FROM users`).
		WillReturnRows(rows)

	users, err := pgDb.GetAllUser()
	if err != nil {
		t.Fatalf("Error getting data %v", err)
	}
	if len(users) > 2 {
		t.Fatalf("Got all data %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Got all data %v", err)
	}
}

func TestGetAllDataById(t *testing.T) {
	db, mock, err := sqlmock.New()
	pgDb := &PostgresDb{Db: db}
	if err != nil {
		t.Fatalf("failed to create new db %v", err)
	}
	defer db.Close()

	createdAt := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	row := sqlmock.NewRows([]string{"id", "name", "email", "phone", "is_active", "created_at"}).
		AddRow(1, "John Doe1", "john.doe@example.com", "1234567890", true, createdAt)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, phone, is_active, created_at FROM users WHERE id = $1")).
		WithArgs(1).
		WillReturnRows(row)

	user, err := pgDb.GetUserById(1)
	if err != nil {
		t.Fatalf("Error getting data: %v", err)
	}

	if user == nil {
		t.Fatalf("Expected user, got nil")
	}

	if user.Id != 1 || user.Name != "John Doe1" || user.Email != "john.doe@example.com" {
		t.Errorf("Expected user ID 1, got %+v", user)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unmet expectations: %v", err)
	}
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	pgDb := &PostgresDb{Db: db}
	if err != nil {
		t.Fatalf("failed to create new db %v", err)
	}
	defer db.Close()

	createdAt := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	user := &models.User{
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		Phone:     "1234567890",
		IsActive:  true,
		CreatedAt: createdAt,
	}
	mock.ExpectQuery(regexp.QuoteMeta(`
        INSERT INTO users (name, email, phone, is_active, created_at)
        VALUES ($1, $2, $3, $4, $5)
    `)).
		WithArgs(user.Name, user.Email, user.Phone, user.IsActive, user.CreatedAt)

	// write all test
}
