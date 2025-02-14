package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	m "github.com/shanto-323/Library_v1.git/models"
)

type PostgresDb struct {
	Db *sql.DB
}

type sqlOpener func(driverName, dataSourceName string) (*sql.DB, error)

func MakePostgresDb(open sqlOpener) (*PostgresDb, error) {
	defaultSrt := "user=postgres password=1234 host=localhost port=5432 sslmode=disable"
	dbName := "library"

	defaultDB, err := open("postgres", defaultSrt)
	if err != nil {
		return nil, err
	}
	defer defaultDB.Close()

	var exists bool
	err = defaultDB.QueryRow("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)", dbName).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if !exists {
		_, err = defaultDB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			return nil, err
		}
	}

	connStr := fmt.Sprintf("user=postgres dbname=%s password=1234 sslmode=disable", dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresDb{Db: db}, nil
}

func (p *PostgresDb) CreateDb() error {
	return p.createTables()
}

func (p *PostgresDb) createTables() error {
	queries := []string{
		`
        CREATE TABLE IF NOT EXISTS genres (
          id SERIAL PRIMARY KEY,
          name VARCHAR(150) NOT NULL UNIQUE
        );`,
		`
        CREATE TABLE IF NOT EXISTS author (
          id SERIAL PRIMARY KEY,
          name VARCHAR(150) NOT NULL UNIQUE,
          birth_year INT CHECK (birth_year > 0),
          nationality VARCHAR(100),
          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );`,
		`
        CREATE TABLE IF NOT EXISTS books (
          id SERIAL PRIMARY KEY,
          isbn VARCHAR(30) NOT NULL UNIQUE,
          title VARCHAR(200) NOT NULL,
          published_year INT NOT NULL,
          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
          total_copies INT NOT NULL CHECK (total_copies >= 0),
          available_copies INT NOT NULL CHECK (available_copies >= 0),
          CONSTRAINT fk_genre FOREIGN KEY (genre) REFERENCES genre(id) ON DELETE CASCADE
        );`,
		`
        CREATE TABLE IF NOT EXISTS book_genres (
          book_id INT REFERENCES books(id) ON DELETE CASCADE,
          genre_id INT REFERENCES genres(id) ON DELETE CASCADE,
          PRIMARY KEY (book_id, genre_id)
        );`,
		`
        CREATE TABLE IF NOT EXISTS book_author (
          book_isbn VARCHAR(30),
          author_id INT,
          PRIMARY KEY (book_isbn, author_id),
          CONSTRAINT fk_book FOREIGN KEY (book_isbn) REFERENCES books(isbn) ON DELETE CASCADE,
          CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES author(id) ON DELETE CASCADE
        );`,
		`
        CREATE TABLE IF NOT EXISTS students (
          id SERIAL PRIMARY KEY,
          name VARCHAR(100) NOT NULL UNIQUE,
          email VARCHAR(150) UNIQUE,
          phone VARCHAR(15) NOT NULL UNIQUE,
          is_active BOOL DEFAULT TRUE,
          created_at DATE DEFAULT CURRENT_DATE
        );`,
		`
        CREATE TABLE IF NOT EXISTS borrowed_book (
          id SERIAL PRIMARY KEY,
          student_id INT NOT NULL,
          book_isbn VARCHAR(20) NOT NULL,
          taken_at DATE DEFAULT CURRENT_DATE,
          return_at DATE NOT NULL,
          returned BOOL DEFAULT FALSE,
          CONSTRAINT fk_book FOREIGN KEY (book_isbn) REFERENCES books(isbn) ON DELETE CASCADE,
          CONSTRAINT fk_student FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE
        );`,
	}

	for _, query := range queries {
		_, err := p.Db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *PostgresDb) GetAllUser() ([]*m.User, error) {
	users := []*m.User{}
	rows, err := p.Db.Query("SELECT id, name, email, phone, is_active, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user, err := getUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (p *PostgresDb) GetUserById(id int) (*m.User, error) {
	row := p.Db.QueryRow("SELECT id, name, email, phone, is_active, created_at FROM users WHERE id = $1", id)

	var user m.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.IsActive, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (p *PostgresDb) CreateUser(user *m.User) error {
	_, err := p.Db.Exec(`
        INSERT INTO users (name, email, phone, is_active, created_at)
        VALUES ($1, $2, $3, $4, NOW())`,
		user.Name, user.Email, user.Phone, user.IsActive)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresDb) UpdateUser(user *m.User) error {
	_, err := p.Db.Exec(`
        UPDATE users
        SET name = $1, email = $2, phone = $3, is_active = $4
        WHERE id = $5`,
		user.Name, user.Email, user.Phone, user.IsActive, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresDb) DeleteUser(id int) error {
	_, err := p.Db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func getUser(row *sql.Rows) (*m.User, error) {
	user := &m.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.IsActive, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
