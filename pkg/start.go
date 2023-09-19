package pkg

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/murasame29/echo-hex-arch-template/pkg/env"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/database/gorm"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/token"
	"github.com/murasame29/echo-hex-arch-template/pkg/server"
)

func Start() {
	ctx := context.Background()
	notifyCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

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

	srv := server.NewServer(ctx, db, maker, &env)

	go func() {
		log.Println("starting server ...")
		if err := srv.Run(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				// サーバが閉じられた以外の操作 => logを出力する必要あり
				log.Printf("異常終了:%v\n", err)
			}
		}
	}()

	<-notifyCtx.Done()

	ctx, cancel := context.WithTimeout(ctx, time.Duration(env.ShutdownTimeout)*time.Second)
	defer cancel()

	if err := srv.Close(ctx); err != nil {
		log.Printf("シャットダウンエラー:%v\n", err)
	}
}
