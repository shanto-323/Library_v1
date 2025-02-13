package storage

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMakePostgresDb(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT EXISTS \\(SELECT 1 FROM pg_database WHERE datname = \\$1\\)").
		WithArgs("library").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	mock.ExpectExec("CREATE DATABASE library").
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectPing()

	open := func(driverName, dataSourceName string) (*sql.DB, error) {
		return db, nil
	}

	pgDb, err := MakePostgresDb(open)
	assert.NoError(t, err, "MakePostgresDb should not return an error")
	assert.NotNil(t, pgDb, "PostgresDb instance should not be nil")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestCreateDb(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	for _, query := range []string{
		`CREATE TABLE IF NOT EXISTS genres`,
		`CREATE TABLE IF NOT EXISTS author`,
		`CREATE TABLE IF NOT EXISTS books`,
		`CREATE TABLE IF NOT EXISTS book_genres`,
		`CREATE TABLE IF NOT EXISTS book_author`,
		`CREATE TABLE IF NOT EXISTS students`,
		`CREATE TABLE IF NOT EXISTS borrowed_book`,
	} {
		mock.ExpectExec(query).
			WillReturnResult(sqlmock.NewResult(0, 1))
	}
	pgDb := &PostgresDb{Db: db}

	err = pgDb.CreateDb()
	assert.NoError(t, err, "CreateDb should not return an error")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}

func TestCreateTablesError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS genre`).
		WillReturnError(errors.New("mocked error"))
	pgDb := &PostgresDb{Db: db}

	err = pgDb.CreateDb()
	assert.Error(t, err, "CreateDb should return an error when table creation fails")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "All expectations should be met")
}
