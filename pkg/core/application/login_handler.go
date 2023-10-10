package application

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/request"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/infrastructure/gateway"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/token"
	"github.com/murasame29/echo-hex-arch-template/pkg/logger"
)

type loginHTTPService interface {
	Login(ctx echo.Context) error
}

type LoginHTTPService struct {
	gw gateway.LoginGateway
}

func NewLoginHTTPService(ctx context.Context, db *sql.DB, maker token.Maker, l logger.Logger) loginHTTPService {
	return &LoginHTTPService{
		gw: gateway.NewLoginGateway(ctx, db, maker, l),
	}
}

func (ls *LoginHTTPService) Login(ctx echo.Context) error {
	var req request.Login
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}

	return ctx.JSON(ls.gw.Login(req))
}
