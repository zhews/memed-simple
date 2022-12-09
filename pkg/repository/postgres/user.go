package postgres

import (
	"database/sql"
	"github.com/zhews/memed-simple/pkg/domain"
)

type UserRepositoryPostgres struct {
	DB *sql.DB
}

const queryGetUserByUsername = "SELECT id, username, name, password_hash, created_at, updated_at FROM memed_user WHERE id = $1"

func (urp *UserRepositoryPostgres) GetByUsername(username string) (domain.User, error) {
	row := urp.DB.QueryRow(queryGetUserByUsername, username)
	var user domain.User
	err := row.Scan(&user.Id, &user.Username, &user.Name, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

const queryInsertUser = "INSERT INTO memed_user (id, username, name, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"

func (urp *UserRepositoryPostgres) Insert(user domain.User) error {
	_, err := urp.DB.Exec(queryInsertUser, user.Id, user.Username, user.Name, user.PasswordHash, user.CreatedAt, user.UpdatedAt)
	return err
}
