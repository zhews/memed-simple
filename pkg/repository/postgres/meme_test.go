package postgres

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v2"
	"github.com/zhews/memed-simple/pkg/domain"
	"reflect"
	"testing"
	"time"
)

const queryBaseDeleteMeme = "DELETE FROM memed_meme"

func TestMemeRepositoryPostgres_Delete(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.FailNow()
	}
	memeId := uuid.New()
	mock.ExpectExec(queryBaseDeleteMeme).
		WithArgs(memeId).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))
	errorId := uuid.New()
	mock.ExpectExec(queryBaseDeleteMeme).
		WithArgs(errorId).
		WillReturnError(errors.New("some error"))
	type fields struct {
		Conn Conn
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Delete a meme",
			fields: fields{
				Conn: mock,
			},
			args: args{
				id: memeId,
			},
			wantErr: false,
		},
		{
			name: "Database error",
			fields: fields{
				Conn: mock,
			},
			args: args{
				id: errorId,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mrp := &MemeRepositoryPostgres{
				Conn: tt.fields.Conn,
			}
			if err := mrp.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %s", err)
	}
}

const queryBaseGetMeme = "SELECT id, title, image, created_by, created_at, updated_at FROM memed_meme"

func TestMemeRepositoryPostgres_GetAll(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.FailNow()
	}
	now := time.Now().Unix()
	meme := domain.Meme{
		Id:        uuid.New(),
		Title:     "a funny meme",
		Image:     "uuid.ext",
		CreatedBy: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
	}
	memes := []domain.Meme{meme}
	mock.ExpectQuery(queryBaseGetMeme).
		WillReturnRows(pgxmock.NewRows([]string{"id", "title", "image", "created_by", "created_at", "updated_at"}).AddRow(meme.Id, meme.Title, meme.Image, meme.CreatedBy, meme.CreatedAt, meme.UpdatedAt))
	mock.ExpectQuery(queryBaseGetMeme).
		WillReturnRows(pgxmock.NewRows([]string{"id", "title", "image", "created_by", "created_at", "updated_at"}).AddRow(meme.Id, meme.Title, meme.Image, meme.CreatedBy, meme.CreatedAt, meme.UpdatedAt).RowError(0, errors.New("some error")))
	mock.ExpectQuery(queryBaseGetMeme).
		WillReturnError(errors.New("some error"))
	type fields struct {
		Conn Conn
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Meme
		wantErr bool
	}{
		{
			name: "Get memes",
			fields: fields{
				Conn: mock,
			},
			want:    memes,
			wantErr: false,
		},
		{
			name: "Row error",
			fields: fields{
				Conn: mock,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Database error",
			fields: fields{
				Conn: mock,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mrp := &MemeRepositoryPostgres{
				Conn: tt.fields.Conn,
			}
			got, err := mrp.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %s", err)
	}
}

func TestMemeRepositoryPostgres_GetById(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.FailNow()
	}
	now := time.Now().Unix()
	meme := domain.Meme{
		Id:        uuid.New(),
		Title:     "funny meme",
		Image:     "uuid.ext",
		CreatedBy: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
	}
	mock.ExpectQuery(queryBaseGetMeme).
		WithArgs(meme.Id).
		WillReturnRows(pgxmock.NewRows([]string{"id", "title", "image", "created_by", "created_at", "updated_at"}).AddRow(meme.Id, meme.Title, meme.Image, meme.CreatedBy, meme.CreatedAt, meme.UpdatedAt))
	nonExistingMemeId := uuid.New()
	mock.ExpectQuery(queryBaseGetMeme).
		WithArgs(nonExistingMemeId).
		WillReturnError(pgx.ErrNoRows)
	errorMemeId := uuid.New()
	mock.ExpectQuery(queryBaseGetMeme).
		WithArgs(errorMemeId).
		WillReturnError(errors.New("some error"))
	type fields struct {
		Conn Conn
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Meme
		wantErr bool
	}{
		{
			name: "Get meme",
			fields: fields{
				Conn: mock,
			},
			args: args{
				id: meme.Id,
			},
			want:    meme,
			wantErr: false,
		},
		{
			name: "Non existing meme",
			fields: fields{
				Conn: mock,
			},
			args: args{
				id: nonExistingMemeId,
			},
			want:    domain.Meme{},
			wantErr: true,
		},
		{
			name: "Database error",
			fields: fields{
				Conn: mock,
			},
			args: args{
				id: errorMemeId,
			},
			want:    domain.Meme{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mrp := &MemeRepositoryPostgres{
				Conn: tt.fields.Conn,
			}
			got, err := mrp.GetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() got = %v, want %v", got, tt.want)
			}
		})
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %s", err)
	}
}

const queryBaseInsertMeme = "INSERT INTO memed_meme"

func TestMemeRepositoryPostgres_Insert(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.FailNow()
	}
	now := time.Now().Unix()
	meme := domain.Meme{
		Id:        uuid.New(),
		Title:     "funny meme",
		Image:     "uuid.ext",
		CreatedBy: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
	}
	mock.ExpectExec(queryBaseInsertMeme).
		WithArgs(meme.Id, meme.Title, meme.Image, meme.CreatedBy, meme.CreatedAt, meme.UpdatedAt).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))
	errorMeme := domain.Meme{
		Id:        uuid.New(),
		Title:     "funny meme",
		Image:     "uuid.ext",
		CreatedBy: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
	}
	mock.ExpectExec(queryBaseInsertMeme).
		WithArgs(errorMeme.Id, errorMeme.Title, errorMeme.Image, errorMeme.CreatedBy, errorMeme.CreatedAt, errorMeme.UpdatedAt).
		WillReturnError(errors.New("some error"))
	type fields struct {
		Conn Conn
	}
	type args struct {
		meme domain.Meme
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Insert meme",
			fields: fields{
				Conn: mock,
			},
			args: args{
				meme: meme,
			},
			wantErr: false,
		},
		{
			name: "Database error",
			fields: fields{
				Conn: mock,
			},
			args: args{
				meme: errorMeme,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mrp := &MemeRepositoryPostgres{
				Conn: tt.fields.Conn,
			}
			if err := mrp.Insert(tt.args.meme); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %s", err)
	}
}

const queryBaseUpdateMeme = "UPDATE memed_meme"

func TestMemeRepositoryPostgres_Update(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.FailNow()
	}
	now := time.Now().Unix()
	creator := uuid.New()
	meme := domain.Meme{
		Id:        uuid.New(),
		Title:     "funny meme",
		Image:     "uuid.ext",
		CreatedBy: creator,
		CreatedAt: now,
		UpdatedAt: now,
	}
	mock.ExpectExec(queryBaseUpdateMeme).
		WithArgs(meme.Title, meme.UpdatedAt, meme.Id, creator).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))
	errorCreator := uuid.New()
	errorMeme := domain.Meme{
		Id:        uuid.New(),
		Title:     "error meme",
		Image:     "uuid.ext",
		CreatedBy: errorCreator,
		CreatedAt: now,
		UpdatedAt: now,
	}
	mock.ExpectExec(queryBaseUpdateMeme).
		WithArgs(errorMeme.Title, errorMeme.UpdatedAt, errorMeme.Id, errorCreator).
		WillReturnError(errors.New("some error"))
	type fields struct {
		Conn Conn
	}
	type args struct {
		meme   domain.Meme
		userId uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Update meme",
			fields: fields{
				Conn: mock,
			},
			args: args{
				meme:   meme,
				userId: creator,
			},
			wantErr: false,
		},
		{
			name: "Database error",
			fields: fields{
				Conn: mock,
			},
			args: args{
				meme:   errorMeme,
				userId: errorCreator,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mrp := &MemeRepositoryPostgres{
				Conn: tt.fields.Conn,
			}
			if err := mrp.Update(tt.args.meme, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %s", err)
	}
}
