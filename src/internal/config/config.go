package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type Config struct {
	ServerAddress string `envconfig:"SERVER_ADDRESS"`
	PostgresConn  string `envconfig:"POSTGRES_CONN"`
	JdbcUrl       string `envconfig:"POSTGRES_JDBC_URL"`
	Username      string `envconfig:"POSTGRES_USERNAME"`
	Password      string `envconfig:"POSTGRES_PASSWORD"`
	Host          string `envconfig:"POSTGRES_HOST"`
	Port          uint16 `envconfig:"POSTGRES_PORT"`
	Database      string `envconfig:"POSTGRES_DATABASE"`
}

func (c *Config) initPostgresConn() {
	replacer := strings.NewReplacer(
		"{POSTGRES_USERNAME}", c.Username,
		"{POSTGRES_PASSWORD}", c.Password,
		"{POSTGRES_HOST}", c.Host,
		"{POSTGRES_PORT}", strconv.Itoa(int(c.Port)),
		"{POSTGRES_DATABASE}", c.Database)
	c.PostgresConn = replacer.Replace(c.PostgresConn)
}

func (c *Config) initJdbcUrl() {
	replacer := strings.NewReplacer(
		"{POSTGRES_HOST}", c.Host,
		"{POSTGRES_PORT}", strconv.Itoa(int(c.Port)),
		"{POSTGRES_DATABASE}", c.Database)
	c.JdbcUrl = replacer.Replace(c.JdbcUrl)
}

func Load() error {
	if err := godotenv.Load(); err != nil {
		return errors.Wrap(err, "Error loading .env file")
	}
	return nil
}

func FromEnv() (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, errors.Wrap(err, "init config")
	}

	cfg.initPostgresConn()
	cfg.initJdbcUrl()

	return cfg, nil
}
