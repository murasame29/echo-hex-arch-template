package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/models"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/request"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/response"
)

// ここではユーザに関するいわゆるリポジトリ層を実装する

type UserStorage interface {
	ListUser(req request.ListUser) ([]response.ListUser, error)
	CreateUser(req request.CreateUser) (response.User, error)
	DeleteAllUser() error
	GetUser(userID string) (response.User, error)
	UpdateUser(arg models.UpdateUserParam) (response.User, error)
	DeleteUser(userID string) error
}

type userStorage struct {
	db  *sql.DB
	ctx context.Context
}

func NewUserStorage(ctx context.Context, db *sql.DB) UserStorage {
	return &userStorage{
		db:  db,
		ctx: ctx,
	}
}

func (us *userStorage) ListUser(req request.ListUser) ([]response.ListUser, error) {
	query := `SELECT user_id,username FROM users LIMIT $1 OFFSET $2`

	rows, err := us.db.QueryContext(us.ctx, query, req.PageSize, (req.PageID-1)*req.PageSize)
	if err != nil {
		return nil, err
	}

	var users []response.ListUser

	for rows.Next() {
		var user response.ListUser
		err = rows.Scan(&user.UserID, &user.Username)

		users = append(users, user)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (us *userStorage) CreateUser(req request.CreateUser) (response.User, error) {
	query := `INSERT INTO users(user_id,username,email,password)VALUES($1,$2,$3,$4)RETURNING user_id,username,email,created_at,updated_at`
	rows := us.db.QueryRowContext(us.ctx, query, uuid.New().String(), req.Username, req.Email, req.Password)

	var user response.User
	err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (us *userStorage) DeleteAllUser() error {
	query := `DELETE FROM users`
	_, err := us.db.ExecContext(us.ctx, query)
	return err
}

func (us *userStorage) GetUser(userID string) (response.User, error) {
	query := `SELECT user_id,username,email,created_at,updated_at FROM users WHERE user_id = $1`
	rows := us.db.QueryRowContext(us.ctx, query, userID)

	var user response.User
	err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (us *userStorage) UpdateUser(arg models.UpdateUserParam) (response.User, error) {
	query := `UPDATE users SET username=$1,email=$2,password=$3,updated_at=$4 WHERE user_id = $5 RETURNING user_id,username,email,created_at,updated_at`
	rows := us.db.QueryRowContext(us.ctx, query,
		arg.Username,
		arg.Email,
		arg.Password,
		time.Now(),
		arg.UserID,
	)

	var user response.User
	err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (us *userStorage) DeleteUser(userID string) error {
	query := `DELETE FROM users WHERE user_id = $1`
	_, err := us.db.ExecContext(us.ctx, query, userID)
	return err
}
