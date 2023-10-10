package gorm

import (
	"context"
	"testing"

	"github.com/murasame29/echo-hex-arch-template/cmd/config"
	"github.com/murasame29/echo-hex-arch-template/pkg/logger"
	"github.com/stretchr/testify/require"
)

func TestConnectDB(t *testing.T) {
	l := logger.NewLogger(logger.DEBUG)
	config.LoadEnv(l)
	db := New(context.Background(), l)

	sqlDB := db.ConnectDB()

	require.NoError(t, sqlDB.Ping())
	require.NoError(t, sqlDB.Close())
}
