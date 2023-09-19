package server

import (
	"context"
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/murasame29/echo-hex-arch-template/pkg/env"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/token"
)

type EchoServer struct {
	*echo.Echo
	ctx   context.Context
	db    *sql.DB
	maker token.Maker
	env   *env.Env
}

func (es *EchoServer) configure() {
	// ここではmiddlewareやcorsを設定する
}

func (es *EchoServer) Run() error {
	return es.Start(":" + es.env.ServerPort)
}

func NewServer(ctx context.Context, db *sql.DB, maker token.Maker, env *env.Env) Server {
	if env.ServerPort == "" {
		env.ServerPort = "8080"
	}

	server := &EchoServer{
		echo.New(),
		ctx,
		db,
		maker,
		env,
	}
	server.configure()
	server.routes()

	return server
}
