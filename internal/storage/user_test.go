package storage

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shanto-323/Library_v1.git/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	newDb := &PostgresDb{db}
	pgDb := NewUserStorage(newDb)
	if err != nil {
		t.Fatalf("failed to create new db %v", err)
	}
	defer db.Close()

	//Get AllUser
	{
		createdAt := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "is_active", "created_at"}).
			AddRow(1, "John Doe1", "john.doe@example.com", "1234567890", true, createdAt).
			AddRow(2, "Jane Doe2", "jane.doe@example.com", "0987654321", false, createdAt)

		mock.ExpectQuery(`SELECT id, name, email, phone, is_active, created_at FROM users`).
			WillReturnRows(rows)

		users, err := pgDb.GetAllUser()
		assert.NoError(t, err, "Error getting data")

		assert.Equal(t, len(users), 2, "found extra data")

		assert.NoError(t, mock.ExpectationsWereMet(), "Unmet expectations")
	}

	mock.ExpectationsWereMet()

	//Test GetUserById
	{
		createdAt := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		row := sqlmock.NewRows([]string{"id", "name", "email", "phone", "is_active", "created_at"}).
			AddRow(1, "John Doe1", "john.doe@example.com", "1234567890", true, createdAt)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, phone, is_active, created_at FROM users WHERE id = $1")).
			WithArgs(1).
			WillReturnRows(row)

		user, err := pgDb.GetUserById(1)
		assert.NoError(t, err, "Error getting data")
		assert.NotNil(t, user, "No user found")

		assert.Equal(t, user.Id, 1, "Id not matched")
		assert.Equal(t, user.Name, "John Doe1", "Name not matched")
		assert.Equal(t, user.Email, "john.doe@example.com", "Gmail not matched")

		assert.NoError(t, mock.ExpectationsWereMet(), "Unmet expectations")
	}

	mock.ExpectationsWereMet()

	//Test CreateUser
	{
		user := &models.User{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "1234567890",
			IsActive: true,
		}

		mock.ExpectExec(regexp.QuoteMeta(`
        INSERT INTO users (name, email, phone, is_active, created_at)
        VALUES ($1, $2, $3, $4, NOW())
    `)).
			WithArgs(user.Name, user.Email, user.Phone, user.IsActive).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err = pgDb.CreateUser(user)
		assert.NoError(t, err, "user creating error")

		assert.NoError(t, mock.ExpectationsWereMet(), "Unmet expectations")
	}

	mock.ExpectationsWereMet()

	//Test TestUpdateUserById
	{
		user := &models.User{
			Id:       1,
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "1234567890",
			IsActive: false,
		}

		mock.ExpectExec(regexp.QuoteMeta(`
        UPDATE users
        SET name = $1, email = $2, phone = $3, is_active = $4
        WHERE id = $5
    `)).
			WithArgs(user.Name, user.Email, user.Phone, user.IsActive, user.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err = pgDb.UpdateUser(user)
		assert.NoError(t, err, "user creating error")

		assert.NoError(t, mock.ExpectationsWereMet(), "Unmet expectations")
	}
	mock.ExpectationsWereMet()

	//Test DeteleById
	{
		mock.ExpectExec(regexp.QuoteMeta(`
       DELETE FROM users WHERE id = $1
    `)).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err = pgDb.DeleteUser(1)
		assert.NoError(t, err, "user creating error")

		assert.NoError(t, mock.ExpectationsWereMet(), "Unmet expectations")
	}
}
