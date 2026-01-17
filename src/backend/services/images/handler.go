package images

import (
	"log"
	"net/http"

	json "github.com/jeevaprakashdr/image-gallery/services"
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

func (h *handler) SearchImages(tag string, w http.ResponseWriter, r *http.Request) {
	images, err := h.service.SearchImages(tag, r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, images)
}
