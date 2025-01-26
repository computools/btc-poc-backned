package config

import (
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerConfig
	DatabaseConfig
	S3Config
	KeyCloakConfig
}

type ServerConfig struct {
	Addr              string        `env:"SERVER_ADDRESS"`
	CorsMaxAge        int           `env:"CORS_MAX_AGE" envDefault:"300"`
	ReadHeaderTimeout time.Duration `env:"SERVER_READ_HEADER_TIMEOUT"`
}

type DatabaseConfig struct {
	DSN           string `env:"DATABASE_DSN"`
	MigrationPath string `env:"DATABASE_MIGRATION_PATH" envDefault:"migrations"`
}

type S3Config struct {
	AccessKey     string `env:"S3_ACCESS_KEY"`
	SecretKey     string `env:"S3_SECRET_KEY"`
	Region        string `env:"S3_REGION" envDefault:"us-east-1"`
	CompanyBucket string `env:"S3_COMPANY_BUCKET" envDefault:"company-data"`
}

type KeyCloakConfig struct {
	Addr              string `env:"KEYCLOAK_ADDRESS"`
	ClientID          string `env:"KEYCLOAK_CLIENT_ID"`
	ClientCredentials string `env:"KEYCLOAK_CLIENT_CREDENTIALS"`
	Realm             string `env:"KEYCLOAK_REALM" envDefault:"btc-users"`
}

func Get() (*Config, error) {
	_ = godotenv.Load()

	cfg, err := env.ParseAs[Config]()
	return &cfg, err
}
