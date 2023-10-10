package server

import (
	"context"
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/murasame29/echo-hex-arch-template/cmd/config"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/token"
	"github.com/murasame29/echo-hex-arch-template/pkg/logger"
)

type EchoServer struct {
	*echo.Echo
	ctx   context.Context
	db    *sql.DB
	maker token.Maker
	l     logger.Logger
}

func (es *EchoServer) configure() {
	// ここではmiddlewareやcorsを設定する
}

func (es *EchoServer) Run() error {
	return es.Start(config.Config.Server.Addr)
}

func (es *EchoServer) Close(ctx context.Context) error {
	return es.Shutdown(ctx)
}

func NewServer(ctx context.Context, db *sql.DB, maker token.Maker, l logger.Logger) Server {

	server := &EchoServer{
		echo.New(),
		ctx,
		db,
		maker,
		l,
	}
	server.configure()
	server.routes()

	return server
}
