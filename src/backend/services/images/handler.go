package images

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	json "github.com/jeevaprakashdr/image-gallery/services"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListImages(w http.ResponseWriter, r *http.Request) {
	images, err := h.service.ListImages(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, images)
}

func (h *handler) SaveImage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("payload")
	title := r.FormValue("title")
	tags := r.FormValue("tags")

	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "Uploaded File: %s\n", handler.Filename)
	fmt.Fprintf(w, "File Size: %d\n", handler.Size)
	fmt.Fprintf(w, "MIME Header: %v\n", handler.Header)
	fmt.Fprintf(w, "Title: %v\n", title)
	fmt.Fprintf(w, "Tags: %v\n", tags)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}

	if !isValidFileType(fileBytes) {
		http.Error(w, "Invalid file type", http.StatusUnsupportedMediaType)
		return
	}

	id := uuid.New().String()
	if err := uploadToObjectStorage(fileBytes, id, w); err != nil {
		http.Error(w, "Error uploading to Object Storage", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully uploaded!")
}

func (h *handler) SearchImages(tag string, w http.ResponseWriter, r *http.Request) {
	images, err := h.service.SearchImages(tag, r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, images)
}

func isValidFileType(file []byte) bool {
	fileType := http.DetectContentType(file)
	return strings.HasPrefix(fileType, "image/")
}

func uploadToObjectStorage(fileBytes []byte, filename string, w http.ResponseWriter) error {
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
	}

	contentType := "image/png"

	info, err := minioClient.PutObject(context.Background(), bucketName, filename+".png", bytes.NewReader(fileBytes), int64(len(fileBytes)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, "Successfully uploaded %s of size %d\n", filename+".png", info.Size)
	return nil
}
