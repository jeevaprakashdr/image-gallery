package images

import (
	"log"
	"net/http"

	json "github.com/jeevaprakashdr/image-gallery/internal"
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
	err := h.service.ListImages(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	images := []string{"image1", "image2"}

	json.Write(w, http.StatusAccepted, images)
}
