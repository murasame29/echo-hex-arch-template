package storage

import (
	"context"
	"database/sql"

	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/models"
)

// ここではユーザに関するいわゆるリポジトリ層を実装する

type LoginStorage interface {
	GetUser(userID string) (models.User, error)
}

type loginStorage struct {
	db  *sql.DB
	ctx context.Context
}

func NewLoginStorage(ctx context.Context, db *sql.DB) LoginStorage {
	return &loginStorage{
		db:  db,
		ctx: ctx,
	}
}

func (ls *loginStorage) GetUser(userID string) (models.User, error) {
	query := `SELECT user_id,username,email,password,created_at,updated_at FROM users WHERE user_id = $1`
	rows := ls.db.QueryRowContext(ls.ctx, query, userID)

	var user models.User
	err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}
