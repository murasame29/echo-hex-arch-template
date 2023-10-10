package config

var Config *config

type config struct {
	Database struct {
		User              string `env:"DB_USER" envDefault:"postgres"`
		Password          string `env:"DB_PASSWORD" envDefault:"postgres"`
		Host              string `env:"DB_HOST" envDefault:"localhost"`
		Port              int    `env:"DB_PORT" envDefault:"5432"`
		Name              string `env:"DB_NAME" envDefault:"example"`
		ConnectionTimeout int    `env:"DB_CONNECTION_TIMEOUT" envDefault:"5"`
		ConnectAttempts   int    `env:"DB_CONNECT_ATTEMPTS" envDefault:"3"`
	}

	Server struct {
		ServerAddr      string `env:"SERVER_ADDR" envDefault:"0.0.0.0:8080"`
		ContextTimeout  int    `env:"SERVER_CONTEXT_TIMEOUT" envDefault:"5"`
		ShutdownTimeout int    `env:"SERVER_SHUTDOWN_TIMEOUT" envDefault:"10"`
	}

	Token struct {
		TokenExpired int    `env:"TOKEN_EXPIRED" envDefault:"10"`
		TokenSecret  string `env:"TOKEN_SECRET" envDefault:"abcdefghijabcdefghijabcdefghijab"`
	}

	NewRelic struct {
		AppName    string `env:"NEWRELIC_APP_NAME"`
		LicenseKey string `env:"NEWRELIC_LICENSE_KEY"`
	}
}
