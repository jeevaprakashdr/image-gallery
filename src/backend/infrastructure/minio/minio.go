package minioClient

import (
	"bytes"
	"context"
	"log"
	"os"

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
	endpoint := os.Getenv("MINIO_URL")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("MINIO_ACCESS_KEY_SECRETE")
	bucketName := os.Getenv("MINIO_GALLERY_BUCKET")

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
