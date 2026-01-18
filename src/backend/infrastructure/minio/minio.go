package minioClient

import (
	"bytes"
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient interface {
	Upload(fileBytes []byte, filename string) (string, error)
}

type minioClient struct {
}

func NewMinioClient() MinioClient {
	return &minioClient{}
}

func (c *minioClient) Upload(fileBytes []byte, filename string) (string, error) {
	endpoint := "localhost:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	bucketName := "images"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	contentType := "image/png"

	info, err := minioClient.PutObject(context.Background(), bucketName, filename+".png", bytes.NewReader(fileBytes), int64(len(fileBytes)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return info.Key, nil
}
