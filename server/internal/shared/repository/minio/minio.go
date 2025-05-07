package minio

import (
	"context"
	"fmt"
	"kino/internal/shared/config"
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3 struct {
	S3 *minio.Client
}

func MinioConnection(conf *config.Config) *S3 {
	endpoint := conf.Minio.Endpoint
	accessKeyID := conf.Minio.AccessKeyID
	secretAccessKey := conf.Minio.SecretAccessKey
	useSSL := conf.Minio.UseSSL

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
		Region: "ru-central-1",
	})
	if err != nil {
		log.Fatalln(fmt.Errorf("minio client new error: %w", err))
	}

	location := "ru-central-1"

	err = minioClient.MakeBucket(context.Background(), "cinema-media", minio.MakeBucketOptions{Region: location})
	if err != nil {

		exists, errBucketExists := minioClient.BucketExists(context.Background(), "cinema-media")
		if errBucketExists == nil && exists {
			fmt.Printf("We already own %s\n", "cinema-media")
		} else {
			log.Fatalln(fmt.Errorf("minio MakeBucket cinema-media error: %w", err))
		}
	} else {
		fmt.Printf("successfully created %s\n", "cinema-media")
	}

	err = minioClient.MakeBucket(context.Background(), "film-media", minio.MakeBucketOptions{Region: location})
	if err != nil {

		exists, errBucketExists := minioClient.BucketExists(context.Background(), "film-media")
		if errBucketExists == nil && exists {
			fmt.Printf("We already own %s\n", "film-media")
		} else {
			log.Fatalln(fmt.Errorf("minio MakeBucket film-media error: %w", err))
		}
	} else {
		fmt.Printf("successfully created %s\n", "film-media")
	}

	return &S3{minioClient}
}

func (s3 *S3) FPutObject(ctx context.Context, bucketName, fileName, filePath, contentType string) error {
	_, err := s3.S3.FPutObject(ctx, bucketName, fileName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}
	return nil
}

func (s3 *S3) PresignedGetObject(ctx context.Context, bucketName, filePath string, expiration time.Duration) (*url.URL, error) {
	object, err := s3.S3.PresignedGetObject(ctx, bucketName, filePath, expiration, nil)
	if err != nil {
		return nil, err
	}

	publicURL := *object
	publicURL.Host = "localhost:10420"
	publicURL.Scheme = "http"

	return object, nil
}

func (s3 *S3) GetObject(ctx context.Context, bucketName, filePath string) (*minio.Object, error) {
	object, err := s3.S3.GetObject(ctx, bucketName, filePath, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (s3 *S3) RemoveObject(ctx context.Context, bucketName, fileName string) error {
	err := s3.S3.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
