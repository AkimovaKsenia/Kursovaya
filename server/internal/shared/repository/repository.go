package repository

import (
	"kino/internal/shared/repository/minio"
	"kino/internal/shared/repository/postgres"
)

type Repository struct {
	DB *postgres.DB
	S3 *minio.S3
}
