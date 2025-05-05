package config

import (
	"os"
)

type Application struct {
	SigningKey string
}

type Db struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Minio struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

type Config struct {
	Application Application
	Db          Db
	Minio       Minio
}

func NewEnvConfig() *Config {
	return &Config{
		Application: Application{
			SigningKey: os.Getenv("SIGNING_KEY"),
		},
		Db: Db{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DB"),
		},
		Minio: Minio{
			Endpoint:        os.Getenv("MINIO_ENDPOINT"),
			AccessKeyID:     os.Getenv("MINIO_ROOT_USER"),
			SecretAccessKey: os.Getenv("MINIO_ROOT_PASSWORD"),
			UseSSL:          os.Getenv("MINIO_USE_SSL") == "true",
		},
	}
}
