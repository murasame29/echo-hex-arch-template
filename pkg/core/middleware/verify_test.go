package middleware

import (
	"fmt"
	"testing"
	"time"

	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/token"
	"github.com/stretchr/testify/require"
)

func TestDecodeJWT(t *testing.T) {
	maker, err := token.NewJWTMaker(helpers.RandomString(32))
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	userID := helpers.RandomString(32)
	email := helpers.RandomEmail()

	testCases := []struct {
		name            string
		createAuthToken func(t *testing.T) string
		checkResult     func(t *testing.T, payload *token.Payload, stauts int)
	}{
		{
			name: "success",
			createAuthToken: func(t *testing.T) string {
				token, err := maker.CreateToken(userID, email, time.Duration(time.Minute*10))
				require.NoError(t, err)
				require.NotEmpty(t, token)
				return fmt.Sprintf("%s %s", AuthorizationTypeBearer, token)
			},
			checkResult: func(t *testing.T, payload *token.Payload, status int) {
				require.Equal(t, TOKEN_OK, status)

				require.Equal(t, email, payload.Email)
				require.Equal(t, userID, payload.UserID)
			},
		},
		{
			name: "fail-token-type",
			createAuthToken: func(t *testing.T) string {
				token, err := maker.CreateToken(userID, email, time.Duration(time.Minute*10))
				require.NoError(t, err)
				require.NotEmpty(t, token)
				return token
			},
			checkResult: func(t *testing.T, payload *token.Payload, status int) {
				require.Equal(t, INVALID_TOKEN, status)
				require.Nil(t, payload)
			},
		},
		{
			name: "fail-not-bearer",
			createAuthToken: func(t *testing.T) string {
				token, err := maker.CreateToken(userID, email, time.Duration(time.Minute*10))
				require.NoError(t, err)
				require.NotEmpty(t, token)
				return fmt.Sprintf("test %s", token)
			},
			checkResult: func(t *testing.T, payload *token.Payload, status int) {
				require.Equal(t, INVALID_TOKEN, status)
				require.Nil(t, payload)
			},
		},
		{
			name: "fail-invalid-token",
			createAuthToken: func(t *testing.T) string {
				token, err := maker.CreateToken(userID, email, time.Duration(time.Minute*10))
				require.NoError(t, err)
				require.NotEmpty(t, token)
				return fmt.Sprintf("%s %s", AuthorizationTypeBearer, token[:30])
			},
			checkResult: func(t *testing.T, payload *token.Payload, status int) {
				require.Equal(t, INVALID_TOKEN, status)
				require.Nil(t, payload)
			},
		},
		{
			name: "fail-expired-token",
			createAuthToken: func(t *testing.T) string {
				token, err := maker.CreateToken(userID, email, time.Duration(time.Minute*-10))
				require.NoError(t, err)
				require.NotEmpty(t, token)
				return fmt.Sprintf("%s %s", AuthorizationTypeBearer, token)
			},
			checkResult: func(t *testing.T, payload *token.Payload, status int) {
				require.Equal(t, INVALID_TOKEN, status)
				require.Nil(t, payload)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token := tc.createAuthToken(t)

			payload, statsu := decodeJWT(maker, token)
			tc.checkResult(t, payload, statsu)
		})
	}
}
