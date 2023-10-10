package config

import (
	"github.com/caarlos0/env"
	"github.com/murasame29/echo-hex-arch-template/pkg/logger"
)

func LoadEnv(l logger.Logger) {
	Config = &config{}

	if err := env.Parse(&Config.Database); err != nil {
		l.Panic(err)
	}

	if err := env.Parse(&Config.Server); err != nil {
		l.Panic(err)
	}

	if err := env.Parse(&Config.Token); err != nil {
		l.Panic(err)
	}

	if err := env.Parse(&Config.NewRelic); err != nil {
		l.Panic(err)
	}
}
