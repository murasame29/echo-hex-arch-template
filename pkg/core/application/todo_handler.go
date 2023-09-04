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

type todoHTTPService interface {
	ListTodo(ctx echo.Context) error
	CreateTodo(ctx echo.Context) error
	DeleteAllTodo(ctx echo.Context) error
	GetTodo(ctx echo.Context) error
	UpdateTodo(ctx echo.Context) error
	DeleteTodo(ctx echo.Context) error
}

type TodoHTTPService struct {
	gw gateway.TodoGateway
}

func NewTodoHTTPService(ctx context.Context, db *sql.DB) todoHTTPService {
	return &TodoHTTPService{
		gw: gateway.NewTodoGateway(ctx, db),
	}
}

func (ts *TodoHTTPService) ListTodo(ctx echo.Context) error {
	var req request.ListTodo
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

	return ctx.JSON(ts.gw.ListUser(req, payload))
}

func (ts *TodoHTTPService) CreateTodo(ctx echo.Context) error {
	var req request.CreateTodo
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

	return ctx.JSON(ts.gw.CreateTodo(req, payload))
}

func (ts *TodoHTTPService) DeleteAllTodo(ctx echo.Context) error {
	var req request.DeleteAllTodo
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

	return ctx.JSON(ts.gw.DeleteAllTodo(req, payload))
}

func (ts *TodoHTTPService) GetTodo(ctx echo.Context) error {
	var req request.TodoURI
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}

	return ctx.JSON(ts.gw.GetTodo(req.TodoID))
}

func (ts *TodoHTTPService) UpdateTodo(ctx echo.Context) error {
	var req request.UpdateTodo
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

	return ctx.JSON(ts.gw.UpdateTodo(req, payload))
}

func (ts *TodoHTTPService) DeleteTodo(ctx echo.Context) error {
	var req request.TodoURI
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}

	return ctx.JSON(ts.gw.DeleteTodo(req.TodoID))
}
