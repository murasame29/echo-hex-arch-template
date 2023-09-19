package middleware

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

type AccessLog struct {
	Timestamp time.Time `json:"timestamp"`
	Address   string    `json:"address"`
	Method    string    `json:"method"`
	Status    int       `json:"status"`
	Latency   int64     `json:"latency(ms)"`
	Path      string    `json:"path"`
}

func AccessLogger() echo.MiddlewareFunc {
	return accessLogger
}

func accessLogger(next echo.HandlerFunc) echo.HandlerFunc {
	var logger AccessLog
	logger.Timestamp = time.Now()
	return func(ctx echo.Context) error {
		logger.Path = ctx.Path()
		logger.Method = ctx.Request().Method

		err := next(ctx)

		logger.Latency = getLatency(logger.Timestamp)
		logger.Status = ctx.Response().Status
		log.Println(logger)
		return err
	}
}

func getLatency(start time.Time) int64 {
	return int64(time.Since(start))
}
