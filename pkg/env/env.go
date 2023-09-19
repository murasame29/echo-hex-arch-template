package env

import (
	"log"

	"github.com/murasame29/echo-hex-arch-template/pkg/internal/database"
	"github.com/murasame29/echo-hex-arch-template/pkg/internal/helpers"
	"github.com/spf13/viper"
)

type EnvStructure struct {
	User            string `mapstructure:"USER"`
	Password        string `mapstructure:"PASSWORD"`
	Host            string `mapstructure:"HOST"`
	Port            string `mapstructure:"PORT"`
	DBname          string `mapstructure:"DB_NAME"`
	ConnectTimeout  string `mapstructure:"CONNECT_TIMEOUT"`
	ConnectAttempts string `mapstructure:"CONNECT_ATTEMPTS"`

	TokenExpired string `mapstructure:"TOKEN_EXPIRED"`
	TokenSecret  string `mapstructure:"TOKEN_SECRET"`

	ServerPort     string `mapstructure:"SERVER_PORT"`
	ContextTimeout string `mapstructure:"CONTEXT_TIMEOUT"`

	ShutdownTimeout string `mapstructure:"SHUTDOWN_TIMEOUT"`
}

type Env struct {
	Dbconfig database.Config

	TokenExpired int
	TokenSecret  string

	ServerPort      string
	ContextTimeout  int
	ShutdownTimeout int
}

func LoadEnvConfig(path string) Env {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	var config EnvStructure

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	port, err := helpers.ParseInt(config.Port)
	if err != nil {
		log.Fatal("Could not parse to int type: ", err)
	}

	connectTimeout, err := helpers.ParseInt(config.ContextTimeout)
	if err != nil {
		log.Fatal("Could not parse to int type: ", err)
	}

	connectAttempts, err := helpers.ParseInt(config.ConnectAttempts)
	if err != nil {
		log.Fatal("Could not parse to int type: ", err)
	}

	contextTimeout, err := helpers.ParseInt(config.ContextTimeout)
	if err != nil {
		log.Fatal("Could not parse to int type: ", err)
	}

	tokenExpired, err := helpers.ParseInt(config.TokenExpired)
	if err != nil {
		log.Fatal("Could not parse to int type: ", err)
	}

	shutdownTimeout, err := helpers.ParseInt(config.ShutdownTimeout)
	if err != nil {
		log.Fatal("Could not parse to int type: ", err)
	}

	return Env{
		Dbconfig: database.Config{
			User:            config.User,
			Password:        config.Password,
			Host:            config.Host,
			Port:            port,
			DBname:          config.DBname,
			ConnectTimeout:  connectTimeout,
			ConnectAttempts: connectAttempts,
		},
		TokenExpired: tokenExpired,
		TokenSecret:  config.TokenSecret,

		ServerPort:      config.ServerPort,
		ContextTimeout:  contextTimeout,
		ShutdownTimeout: shutdownTimeout,
	}
}
