package images

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
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

	title := r.FormValue("title")
	tags := r.FormValue("tags")
	file, _, err := r.FormFile("payload")

	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

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

	img, err := imaging.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		http.Error(w, "Failed to decode image", http.StatusInternalServerError)
		return
	}

	resizedImg := imaging.Resize(img, 150, 150, imaging.Lanczos)

	var buf bytes.Buffer
	err = imaging.Encode(&buf, resizedImg, imaging.PNG)
	if err != nil {
		http.Error(w, "Failed to encode resized image", http.StatusInternalServerError)
		return
	}
	resizedImgBytes := buf.Bytes()

	id := uuid.New()
	if err := uploadToObjectStorage(resizedImgBytes, "scaled-"+id.String(), w); err != nil {
		http.Error(w, "Error uploading to save image to gallery", http.StatusInternalServerError)
		return
	}

	if err := h.service.SaveImageDetails(title, tags, id, r.Context()); err != nil {
		http.Error(w, "Error uploading to save image to gallery", http.StatusInternalServerError)
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
