package images

import (
	"context"
)

type Service interface {
	ListImages(ctx context.Context) error
}

type imageService struct {
}

func NewService() Service {
	return &imageService{}
}

func (s *imageService) ListImages(ctx context.Context) error {
	return nil
}
