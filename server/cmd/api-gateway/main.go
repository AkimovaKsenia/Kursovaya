package main

import (
	"fmt"
	"kino/internal/api-gateway/handler"
	"kino/internal/shared/config"
	"kino/internal/shared/log"
	"kino/internal/shared/repository"
	"kino/internal/shared/repository/minio"
	"kino/internal/shared/repository/postgres"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}

	conf := config.NewEnvConfig()
	log := log.InitLogger(conf)

	db := postgres.NewDatabase(conf)
	s3 := minio.MinioConnection(conf)

	repo := repository.Repository{DB: db, S3: s3}
	repo.CreateMocks()

	handlers := handler.NewHandler(&repo, log, conf)
	handlers.InitRouter()
}
