package gorm

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/murasame29/echo-hex-arch-template/cmd/config"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/database"
	"github.com/murasame29/echo-hex-arch-template/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	ctx context.Context
	l   logger.Logger
}

func (g *Gorm) Ping(err error, db *sql.DB) error {
	if err != nil {
		return err
	}

	if err = db.PingContext(g.ctx); err != nil {
		return err
	}

	return nil
}

func (g *Gorm) connect(dsn string) (*sql.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if g.Ping(err, sqlDB) != nil {
		return nil, err
	}

	return sqlDB, nil
}

func (g *Gorm) ConnectDB() *sql.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Config.Database.Host, config.Config.Database.User, config.Config.Database.Password, config.Config.Database.Name, config.Config.Database.Port)

	for i := 0; i < config.Config.Database.ConnectAttempts; i++ {
		db, err := g.connect(dsn)

		if database.Try(err, config.Config.Database.ConnectionTimeout, i) == nil {
			return db
		}

		continue
	}

	return nil
}

func New(ctx context.Context, l logger.Logger) database.Database {
	return &Gorm{
		ctx: ctx,
		l:   l,
	}
}
