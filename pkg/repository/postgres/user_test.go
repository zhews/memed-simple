package postgres

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/zhews/memed-simple/pkg/domain"
	"reflect"
	"testing"
)

// These tests require the test environment to be running. Take a look at the "test" folder for more information.

const queryBaseGetUserByUsername = "SELECT id, username, name, password_hash FROM memed_user"

func TestUserRepositoryPostgres_GetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.FailNow()
	}
	user := domain.User{
		Id:           uuid.New(),
		Username:     "zhews",
		Name:         "First Last",
		PasswordHash: "password",
	}
	nonExistingUser := "nonExisting"
	mock.ExpectQuery(queryBaseGetUserByUsername).
		WithArgs(user.Username).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "username", "name", "password_hash"}).AddRow(
				user.Id,
				user.Username,
				user.Name,
				user.PasswordHash,
			),
		)
	mock.ExpectQuery(queryBaseGetUserByUsername).
		WithArgs(nonExistingUser).
		WillReturnError(sql.ErrNoRows)
	type fields struct {
		DB *sql.DB
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
			name: "Get the user that is in the database",
			fields: fields{
				DB: db,
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
				DB: db,
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
				DB: tt.fields.DB,
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
}

const queryBaseInsertUser = "INSERT INTO memed_user"

func TestUserRepositoryPostgres_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.FailNow()
	}
	user := domain.User{
		Id:           uuid.New(),
		Username:     "zhews",
		Name:         "First Last",
		PasswordHash: "password",
	}
	mock.ExpectExec(queryBaseInsertUser).
		WithArgs(user.Id, user.Username, user.Name, user.PasswordHash).
		WillReturnResult(sqlmock.NewResult(0, 0))
	type fields struct {
		DB *sql.DB
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
				DB: db,
			},
			args: args{
				user: user,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urp := &UserRepositoryPostgres{
				DB: tt.fields.DB,
			}
			if err := urp.Insert(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
