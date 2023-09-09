package gorm

import (
	"context"
	"testing"

	"github.com/murasame29/echo-hex-arch-template/pkg/env"
	"github.com/stretchr/testify/require"
)

func TestConnectDB(t *testing.T) {
	env := env.LoadEnvConfig("../../../../")
	db := New(context.Background(), &env)

	sqlDB := db.ConnectDB()

	require.NoError(t, sqlDB.Ping())
	require.NoError(t, sqlDB.Close())
}
