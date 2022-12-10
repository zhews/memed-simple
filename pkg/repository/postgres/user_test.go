package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pashagolub/pgxmock/v2"
	"github.com/zhews/memed-simple/pkg/domain"
	"reflect"
	"testing"
	"time"
)

const queryBaseGetUserByUsername = "SELECT id, username, name, admin, password_hash, created_at, updated_at FROM memed_user"

func TestUserRepositoryPostgres_GetByUsername(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.FailNow()
	}
	now := time.Now().Unix()
	user := domain.User{
		Id:           uuid.New(),
		Username:     "zhews",
		Name:         "First Last",
		Admin:        false,
		PasswordHash: []byte{},
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	nonExistingUser := "nonExisting"
	mock.ExpectQuery(queryBaseGetUserByUsername).
		WithArgs(user.Username).
		WillReturnRows(
			pgxmock.NewRows([]string{"id", "username", "name", "admin", "password_hash", "created_at", "updated_at"}).AddRow(
				user.Id,
				user.Username,
				user.Name,
				user.Admin,
				user.PasswordHash,
				user.CreatedAt,
				user.UpdatedAt,
			),
		)
	mock.ExpectQuery(queryBaseGetUserByUsername).
		WithArgs(nonExistingUser).
		WillReturnError(sql.ErrNoRows)
	type fields struct {
		DB Conn
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
		wantErr bool
	}{
		{
			name: "Get user that is in the database",
			fields: fields{
				DB: mock,
			},
			args: args{
				username: user.Username,
			},
			want:    user,
			wantErr: false,
		},
		{
			name: "Get user that is not in the database",
			fields: fields{
				DB: mock,
			},
			args: args{
				username: nonExistingUser,
			},
			want:    domain.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urp := &UserRepositoryPostgres{
				Conn: tt.fields.DB,
			}
			got, err := urp.GetByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUsername() got = %v, want %v", got, tt.want)
			}
		})
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %s", err)
	}
}

const queryBaseInsertUser = "INSERT INTO memed_user"

func TestUserRepositoryPostgres_Insert(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.FailNow()
	}
	now := time.Now().Unix()
	user := domain.User{
		Id:           uuid.New(),
		Username:     "zhews",
		Name:         "First Last",
		Admin:        true,
		PasswordHash: []byte{},
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	duplicateUser := domain.User{
		Id:           uuid.New(),
		Username:     "zhews",
		Name:         "First Last",
		Admin:        true,
		PasswordHash: []byte{},
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	mock.ExpectExec(queryBaseInsertUser).
		WithArgs(user.Id, user.Username, user.Name, user.Admin, user.PasswordHash, user.CreatedAt, user.UpdatedAt).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))
	mock.ExpectExec(queryBaseInsertUser).
		WithArgs(duplicateUser.Id, duplicateUser.Username, duplicateUser.Name, duplicateUser.Admin, duplicateUser.PasswordHash, duplicateUser.CreatedAt, duplicateUser.UpdatedAt).
		WillReturnError(&pgconn.PgError{Code: pgerrcode.UniqueViolation})
	type fields struct {
		Conn Conn
	}
	type args struct {
		user domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Insert user",
			fields: fields{
				Conn: mock,
			},
			args: args{
				user: user,
			},
			wantErr: false,
		},
		{
			name: "Insert user with invalid username",
			fields: fields{
				Conn: mock,
			},
			args: args{
				user: duplicateUser,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urp := &UserRepositoryPostgres{
				Conn: tt.fields.Conn,
			}
			if err := urp.Insert(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %s", err)
	}
}
