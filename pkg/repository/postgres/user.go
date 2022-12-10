package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/zhews/memed-simple/pkg/domain"
	"github.com/zhews/memed-simple/pkg/repository"
)

type UserRepositoryPostgres struct {
	Conn Conn
}

const queryGetUserByUsername = "SELECT id, username, name, admin, password_hash, created_at, updated_at FROM memed_user WHERE id = $1"

func (urp *UserRepositoryPostgres) GetByUsername(username string) (domain.User, error) {
	row := urp.Conn.QueryRow(context.Background(), queryGetUserByUsername, username)
	var user domain.User
	err := row.Scan(&user.Id, &user.Username, &user.Name, &user.Admin, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return user, repository.ErrorNoRows
	}
	return user, err
}

const queryInsertUser = "INSERT INTO memed_user (id, username, name, admin, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

func (urp *UserRepositoryPostgres) Insert(user domain.User) error {
	_, err := urp.Conn.Exec(context.Background(), queryInsertUser, user.Id, user.Username, user.Name, user.Admin, user.PasswordHash, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
				return repository.ErrorUsernameAlreadyTaken
			}
		}
		return err
	}
	return nil
}
