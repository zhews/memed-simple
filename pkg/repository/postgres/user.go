package postgres

import (
	"context"
	"github.com/zhews/memed-simple/pkg/domain"
)

type UserRepositoryPostgres struct {
	Conn Conn
}

const queryGetUserByUsername = "SELECT id, username, name, admin, password_hash, created_at, updated_at FROM memed_user WHERE id = $1"

func (urp *UserRepositoryPostgres) GetByUsername(username string) (domain.User, error) {
	row := urp.Conn.QueryRow(context.Background(), queryGetUserByUsername, username)
	var user domain.User
	err := row.Scan(&user.Id, &user.Username, &user.Name, &user.Admin, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

const queryInsertUser = "INSERT INTO memed_user (id, username, name, admin, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

func (urp *UserRepositoryPostgres) Insert(user domain.User) error {
	_, err := urp.Conn.Exec(context.Background(), queryInsertUser, user.Id, user.Username, user.Name, user.Admin, user.PasswordHash, user.CreatedAt, user.UpdatedAt)
	return err
}
