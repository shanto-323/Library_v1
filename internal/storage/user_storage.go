package storage

import (
	"database/sql"
	"fmt"

	m "github.com/shanto-323/Library_v1.git/pkg/types"
)

type UserStorage interface {
	GetAllUser() ([]*m.User, error)
	GetUserById(int) (*m.User, error)
	CreateUser(*m.User) error
	UpdateUser(*m.User) error
	DeleteUser(int) error
}

type UserStore struct {
	*PostgresDb
}

func NewUserStorage(db *PostgresDb) *UserStore {
	return &UserStore{db}
}

func (p *UserStore) GetAllUser() ([]*m.User, error) {
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

func (p *UserStore) GetUserById(id int) (*m.User, error) {
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

func (p *UserStore) CreateUser(user *m.User) error {
	_, err := p.Db.Exec(`
        INSERT INTO users (name, email, phone, is_active, created_at)
        VALUES ($1, $2, $3, $4, NOW())`,
		user.Name, user.Email, user.Phone, user.IsActive)
	if err != nil {
		return err
	}
	return nil
}

func (p *UserStore) UpdateUser(user *m.User) error {
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

func (p *UserStore) DeleteUser(id int) error {
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
