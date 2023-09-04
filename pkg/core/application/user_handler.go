package application

import (
	"context"
	"database/sql"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/entities/request"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/infrastructure/gateway"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/middleware"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/token"
)

type userHTTPService interface {
	ListUser(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
	DeleteAllUser(ctx echo.Context) error
	GetUser(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
}

type UserHTTPService struct {
	gw gateway.UserGateway
}

func NewUserHTTPService(ctx context.Context, db *sql.DB) userHTTPService {
	return &UserHTTPService{
		gw: gateway.NewUserGateway(ctx, db),
	}
}

func (us *UserHTTPService) ListUser(ctx echo.Context) error {
	var req request.ListUser
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}

	return ctx.JSON(us.gw.ListUser(req))
}

func (us *UserHTTPService) CreateUser(ctx echo.Context) error {
	var req request.CreateUser
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}

	return ctx.JSON(us.gw.CreateUser(req))
}

func (us *UserHTTPService) DeleteAllUser(ctx echo.Context) error {
	return ctx.JSON(us.gw.DeleteAllUser())
}

func (us *UserHTTPService) GetUser(ctx echo.Context) error {
	var req request.UserURI
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}

	switch ctx.Get(middleware.AuthorizationStatus).(int) {

	case middleware.NOT_READY:
		return ctx.JSON(http.StatusUnauthorized, "invalid header")

	case middleware.TOKEN_NOT_FOUNT:
		return ctx.JSON(http.StatusUnauthorized, "authorization is empty")

	case middleware.INVALID_TOKEN:
		return ctx.JSON(http.StatusUnauthorized, token.ErrInvalidToken)
	}

	payload := ctx.Get(middleware.AuthorizationPayloadKey).(*token.Payload)

	return ctx.JSON(us.gw.GetUser(req.UserID, payload))
}

func (us *UserHTTPService) UpdateUser(ctx echo.Context) error {
	var req request.UpdateUser
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request! 1")
	}

	switch ctx.Get(middleware.AuthorizationStatus).(int) {

	case middleware.NOT_READY:
		return ctx.JSON(http.StatusUnauthorized, "invalid header")

	case middleware.TOKEN_NOT_FOUNT:
		return ctx.JSON(http.StatusUnauthorized, "authorization is empty")

	case middleware.INVALID_TOKEN:
		return ctx.JSON(http.StatusUnauthorized, token.ErrInvalidToken)
	}

	payload := ctx.Get(middleware.AuthorizationPayloadKey).(*token.Payload)

	return ctx.JSON(us.gw.UpdateUser(req, payload))
}

func (us *UserHTTPService) DeleteUser(ctx echo.Context) error {
	var req request.UserURI
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}

	switch ctx.Get(middleware.AuthorizationStatus).(int) {

	case middleware.NOT_READY:
		return ctx.JSON(http.StatusUnauthorized, "invalid header")

	case middleware.TOKEN_NOT_FOUNT:
		return ctx.JSON(http.StatusUnauthorized, "authorization is empty")

	case middleware.INVALID_TOKEN:
		return ctx.JSON(http.StatusUnauthorized, token.ErrInvalidToken)
	}

	payload := ctx.Get(middleware.AuthorizationPayloadKey).(*token.Payload)

	return ctx.JSON(us.gw.DeleteUser(req.UserID, payload))
}
