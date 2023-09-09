package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers/token"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationStatus     = "authorization_status"
	AuthorizationPayloadKey = "authorization_payload"
)

const (
	NOT_READY = iota
	TOKEN_NOT_FOUNT
	TOKEN_OK
	INVALID_TOKEN
)

func Verify(maker token.Maker) echo.MiddlewareFunc {
	return verifyToken(maker)
}

func verifyToken(maker token.Maker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authorizationHeader := getHeader(ctx, AuthorizationHeaderKey)
			ctx.Set(AuthorizationStatus, NOT_READY)

			if len(authorizationHeader) == 0 {
				ctx.Set(AuthorizationStatus, TOKEN_NOT_FOUNT)
			} else {
				payload, status := decodeJWT(maker, authorizationHeader)

				ctx.Set(AuthorizationStatus, status)
				ctx.Set(AuthorizationPayloadKey, payload)
			}

			err := next(ctx)

			return err
		}
	}
}

func decodeJWT(maker token.Maker, authToken string) (*token.Payload, int) {
	fields := strings.Fields(authToken)
	if len(fields) < 2 {
		return nil, INVALID_TOKEN
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != AuthorizationTypeBearer {
		return nil, INVALID_TOKEN
	}

	accessToken := fields[1]
	payload, err := maker.VerifyToken(accessToken)
	if err != nil {
		return nil, INVALID_TOKEN
	}
	return payload, TOKEN_OK
}

func getHeader(ctx echo.Context, key string) string {
	return ctx.Request().Header.Get(key)
}
