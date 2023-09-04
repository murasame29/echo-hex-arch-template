package database

import (
	"database/sql"
	"log"
	"time"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     int
	DBname   string

	// タイムアウトするまでの時間
	ConnectTimeout int
	// 接続する試行回数
	ConnectAttempts int
}

type Database interface {
	ConnectDB() *sql.DB
	Ping(err error, db *sql.DB) error
}

func Try(err error, timeout int, count int) error {
	if err != nil {
		// TODO:logをとったほうがいいかも
		log.Println("DBの接続に失敗しました: ", err)
		time.Sleep(time.Duration(timeout) * time.Second)
	}
	return err
}
