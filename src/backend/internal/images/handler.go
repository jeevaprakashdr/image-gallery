package images

import (
	"net/http"

	json "github.com/jeevaprakashdr/image-gallery/internal"
)

type handler struct {
	service ImageService
}

func NewHandler(service ImageService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListImages(w http.ResponseWriter, r *http.Request) {
	// call service to list images
	// return json in an http response

	images := []string{"image1", "image2"}

	json.Write(w, http.StatusAccepted, images)
}
