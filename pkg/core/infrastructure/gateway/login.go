package gateway

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/murasame29/echo-hex-arch-template/cmd/config"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/request"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/response"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/infrastructure/storage"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/password"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/token"
	"github.com/murasame29/echo-hex-arch-template/pkg/logger"
)

type LoginGateway interface {
	Login(body request.Login) (int, response.Login)
}

type LoginLogic struct {
	storage storage.LoginStorage
	maker   token.Maker
	l       logger.Logger
}

func NewLoginGateway(ctx context.Context, db *sql.DB, maker token.Maker, l logger.Logger) LoginGateway {
	return &LoginLogic{storage.NewLoginStorage(ctx, db), maker, l}
}

func (ll *LoginLogic) Login(body request.Login) (int, response.Login) {
	user, err := ll.storage.GetUser(body.UserID)
	if err != nil {
		return http.StatusInternalServerError, response.Login{}
	}

	if err := password.CheckPassword(body.Password, user.Password); err != nil {
		return http.StatusUnauthorized, response.Login{}
	}

	token, err := ll.maker.CreateToken(body.UserID, user.Email, time.Duration(config.Config.Token.Expired)*24*time.Hour)
	if err != nil {
		return http.StatusInternalServerError, response.Login{}
	}

	return http.StatusOK, response.Login{
		UserID: body.UserID,
		Token:  token,
	}
}
