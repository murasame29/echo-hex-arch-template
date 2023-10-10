package pkg

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/murasame29/echo-hex-arch-template/cmd/config"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/database/gorm"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/token"
	"github.com/murasame29/echo-hex-arch-template/pkg/logger"
	"github.com/murasame29/echo-hex-arch-template/pkg/server"
)

func Start(l logger.Logger) {
	ctx := context.Background()
	notifyCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	db := gorm.New(ctx, l).ConnectDB()
	if db == nil {
		log.Fatal("DBの接続が出来ませんでした")
	}

	defer db.Close()

	maker, err := token.NewJWTMaker(config.Config.Token.Secret)
	if err != nil {
		log.Fatal("token Maker Error :", err)
	}

	srv := server.NewServer(ctx, db, maker, l)

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

	ctx, cancel := context.WithTimeout(ctx, time.Duration(config.Config.Server.ShutdownTimeout)*time.Second)
	defer cancel()

	if err := srv.Close(ctx); err != nil {
		log.Printf("シャットダウンエラー:%v\n", err)
	}
}
