package pkg

import (
	"context"
	"log"

	"github.com/murasame29/echo-hex-arch-template/pkg/env"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/database/gorm"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/token"
	"github.com/murasame29/echo-hex-arch-template/pkg/server"
)

func Start() {
	ctx := context.Background()

	env := env.LoadEnvConfig("./")

	db := gorm.New(ctx, &env).ConnectDB()
	if db == nil {
		log.Fatal("DBの接続が出来ませんでした")
	}

	defer db.Close()

	maker, err := token.NewJWTMaker(env.TokenSecret)
	if err != nil {
		log.Fatal("token Maker Error :", err)
	}

	server.NewServer(ctx, db, maker, &env).Run()
}
