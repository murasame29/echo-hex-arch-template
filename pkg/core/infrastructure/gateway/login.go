package gateway

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/request"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/response"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/infrastructure/storage"
	"github.com/murasame29/echo-hex-arch-template/pkg/env"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/password"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/token"
)

type LoginGateway interface {
	Login(body request.Login) (int, response.Login)
}

type LoginLogic struct {
	storage storage.LoginStorage
	maker   token.Maker
	env     env.Env
}

func NewLoginGateway(ctx context.Context, db *sql.DB, maker token.Maker, env env.Env) LoginGateway {
	return &LoginLogic{storage.NewLoginStorage(ctx, db), maker, env}
}

func (ll *LoginLogic) Login(body request.Login) (int, response.Login) {
	user, err := ll.storage.GetUser(body.UserID)
	if err != nil {
		return http.StatusInternalServerError, response.Login{}
	}

	if err := password.CheckPassword(body.Password, user.Password); err != nil {
		return http.StatusUnauthorized, response.Login{}
	}

	token, err := ll.maker.CreateToken(body.UserID, user.Email, time.Duration(ll.env.TokenExpired)*24*time.Hour)
	if err != nil {
		return http.StatusInternalServerError, response.Login{}
	}

	return http.StatusOK, response.Login{
		UserID: body.UserID,
		Token:  token,
	}
}
