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
				decodeJWT(ctx, maker, authorizationHeader)
			}

			err := next(ctx)

			return err
		}
	}
}

func decodeJWT(ctx echo.Context, maker token.Maker, authToken string) {
	fields := strings.Fields(authToken)
	if len(fields) < 2 {
		ctx.Set(AuthorizationStatus, INVALID_TOKEN)
		return
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != AuthorizationTypeBearer {
		ctx.Set(AuthorizationStatus, INVALID_TOKEN)
		return
	}

	accessToken := fields[1]
	payload, err := maker.Verifytoken(accessToken)
	if err != nil {
		ctx.Set(AuthorizationStatus, INVALID_TOKEN)
		return
	}

	ctx.Set(AuthorizationStatus, TOKEN_OK)
	ctx.Set(AuthorizationPayloadKey, payload)
}

func getHeader(ctx echo.Context, key string) string {
	return ctx.Request().Header.Get(key)
}
