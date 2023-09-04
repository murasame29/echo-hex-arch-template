package gorm

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/murasame29/echo-hex-arch-template/pkg/env"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	ctx    context.Context
	config database.Config
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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", g.config.Host, g.config.User, g.config.Password, g.config.DBname, g.config.Port)

	for i := 0; i < g.config.ConnectAttempts; i++ {
		db, err := g.connect(dsn)

		if database.Try(err, g.config.ConnectTimeout, i) == nil {
			return db
		}

		continue
	}

	return nil
}

func New(ctx context.Context, env *env.Env) database.Database {
	return &Gorm{
		ctx:    ctx,
		config: env.Dbconfig,
	}
}
