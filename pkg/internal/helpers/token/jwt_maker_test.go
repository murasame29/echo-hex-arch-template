package token

import (
	"log"
	"testing"
	"time"

	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers"
	"github.com/stretchr/testify/require"
)

func newToken(t *testing.T) Maker {
	secretKey := helpers.RandomString(32)

	maker, err := NewJWTMaker(secretKey)
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	return maker
}

func TestErrJWTMaker(t *testing.T) {
	secretKey := helpers.RandomString(30)
	_, err := NewJWTMaker(secretKey)
	require.Error(t, err)
}

func TestJWTToken(t *testing.T) {
	maker := newToken(t)

	email := helpers.RandomEmail()
	userID := helpers.RandomString(10)

	testCases := []struct {
		name       string
		duration   time.Duration
		build      func(t *testing.T, duration time.Duration) string
		checkToken func(t *testing.T, token string)
	}{
		{
			name:     "success case",
			duration: time.Duration(time.Minute * 10),
			build: func(t *testing.T, duration time.Duration) string {
				token, err := maker.CreateToken(userID, email, duration)
				require.NoError(t, err)
				require.NotEmpty(t, token)

				return token
			},
			checkToken: func(t *testing.T, token string) {
				payload, err := maker.VerifyToken(token)
				require.NoError(t, err)
				require.NotEmpty(t, payload)

				require.Equal(t, email, payload.Email)
				require.Equal(t, userID, payload.UserID)
			},
		},
		{
			name:     "token has expired",
			duration: time.Duration(time.Minute * -10),
			build: func(t *testing.T, duration time.Duration) string {
				token, err := maker.CreateToken(userID, email, duration)
				require.NoError(t, err)
				require.NotEmpty(t, token)

				return token
			},
			checkToken: func(t *testing.T, token string) {
				_, err := maker.VerifyToken(token)
				require.Error(t, err)
				require.Equal(t, ErrExpiredToken, err)
			},
		},
		{
			name:     "token is invalid",
			duration: time.Duration(time.Minute * 10),
			build: func(t *testing.T, duration time.Duration) string {
				token, err := maker.CreateToken(userID, email, duration)
				require.NoError(t, err)
				require.NotEmpty(t, token)

				return token
			},
			checkToken: func(t *testing.T, token string) {
				_, err := maker.VerifyToken(token[:20])
				log.Println(err)
				require.Error(t, err)
				require.Equal(t, ErrInvalidToken, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checkToken(t, tc.build(t, tc.duration))
		})
	}
}
