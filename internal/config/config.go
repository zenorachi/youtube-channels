package config

import (
	"fmt"
	"os"
	"time"

	"github.com/zenorachi/youtube-task/pkg/database/postgres"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const envFile = ".env"

type Config struct {
	DB              postgres.DBConfig
	DBIntegration   bool          `envconfig:"DB_INTEGRATION" default:"false"`
	Host            string        `envconfig:"HOST" default:"0.0.0.0"`
	Port            int           `envconfig:"PORT" default:"8080"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"5s"`
	APIKey          string        `envconfig:"API_KEY" required:"true"`
}

func New() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(envFile); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if err := envconfig.Process(postgres.ConfigPrefix, &cfg.DB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) HTTPAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
