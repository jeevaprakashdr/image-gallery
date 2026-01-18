package images

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	minioClient "github.com/jeevaprakashdr/image-gallery/infrastructure/minio"
	json "github.com/jeevaprakashdr/image-gallery/services"
	"github.com/jeevaprakashdr/image-gallery/services/imageProcessors"
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

	processor := imageProcessors.NewImageProcessor()
	resizedImgBytes, err := processor.ResizeImage(fileBytes)
	if err != nil {
		http.Error(w, "Error failed to resize image", http.StatusInternalServerError)
		return
	}

	id := uuid.New()
	client := minioClient.NewMinioClient()
	key, err := client.Upload(resizedImgBytes, "scaled-"+id.String())
	if err != nil {
		http.Error(w, "Error uploading to save image to gallery", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Successfully uploaded with key %s\n", key)

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
