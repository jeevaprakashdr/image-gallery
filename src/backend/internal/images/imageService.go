package images

import (
	"context"

	repository "github.com/jeevaprakashdr/image-gallery/postgres/sqlc"
)

type Service interface {
	ListImages(ctx context.Context) ([]repository.Image, error)
}

type imageService struct {
	repository repository.Querier
}

func NewService(repo repository.Querier) Service {
	return &imageService{repo}
}

func (s *imageService) ListImages(ctx context.Context) ([]repository.Image, error) {
	return s.repository.ListImages(ctx)
}
