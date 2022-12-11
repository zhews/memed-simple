package postgres

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/zhews/memed-simple/pkg/domain"
	"github.com/zhews/memed-simple/pkg/repository"
)

type MemeRepositoryPostgres struct {
	Conn Conn
}

const queryGetAllMemes = "SELECT id, title, image, created_by, created_at, updated_at FROM memed_meme"

func (mrp *MemeRepositoryPostgres) GetAll() ([]domain.Meme, error) {
	rows, err := mrp.Conn.Query(context.Background(), queryGetAllMemes)
	if err != nil {
		return nil, err
	}
	var memes []domain.Meme
	for rows.Next() {
		var meme domain.Meme
		err = rows.Scan(&meme.Id, &meme.Title, &meme.Image, &meme.CreatedBy, &meme.CreatedAt, &meme.UpdatedAt)
		if err != nil {
			return nil, err
		}
		memes = append(memes, meme)
	}
	return memes, err
}

const queryGetMemeById = "SELECT id, title, image, created_by, created_at, updated_at FROM memed_meme WHERE id = $1"

func (mrp *MemeRepositoryPostgres) GetById(id uuid.UUID) (domain.Meme, error) {
	row := mrp.Conn.QueryRow(context.Background(), queryGetMemeById, id)
	var meme domain.Meme
	err := row.Scan(&meme.Id, &meme.Title, &meme.Image, &meme.CreatedBy, &meme.CreatedAt, &meme.UpdatedAt)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return meme, repository.ErrorNoRows
	}
	return meme, err
}

const queryInsertMeme = "INSERT INTO memed_meme (id, title, image, created_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"

func (mrp *MemeRepositoryPostgres) Insert(meme domain.Meme) error {
	_, err := mrp.Conn.Exec(context.Background(), queryInsertMeme, meme.Id, meme.Title, meme.Image, meme.CreatedBy, meme.CreatedAt, meme.UpdatedAt)
	return err
}

const queryUpdateMeme = "UPDATE memed_meme SET title = $1, updated_at = $2 WHERE id = $3"

func (mrp *MemeRepositoryPostgres) Update(meme domain.Meme) error {
	_, err := mrp.Conn.Exec(context.Background(), queryUpdateMeme, meme.Title, meme.UpdatedAt, meme.Id)
	return err
}

const queryDeleteMeme = "DELETE FROM memed_meme WHERE id = $1"

func (mrp *MemeRepositoryPostgres) Delete(id uuid.UUID) error {
	_, err := mrp.Conn.Exec(context.Background(), queryDeleteMeme, id)
	return err
}
