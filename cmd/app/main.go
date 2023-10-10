package main

import (
	"github.com/murasame29/echo-hex-arch-template/cmd/config"
	"github.com/murasame29/echo-hex-arch-template/pkg"
	"github.com/murasame29/echo-hex-arch-template/pkg/logger"
)

func main() {
	l := logger.NewLogger(logger.DEBUG)
	config.LoadEnv(l)
	pkg.Start(l)
}
