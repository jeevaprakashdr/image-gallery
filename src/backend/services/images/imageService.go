package images

import (
	"context"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	repository "github.com/jeevaprakashdr/image-gallery/infrastructure/postgres/sqlc"
)

type Service interface {
	SaveImageDetails(title string, tags string, id uuid.UUID, ctx context.Context) error
	ListImages(ctx context.Context) ([]repository.Image, error)
	SearchImages(tag string, ctx context.Context) ([]repository.Image, error)
}

type imageService struct {
	repository repository.Querier
}

func NewService(repo repository.Querier) Service {
	return &imageService{repo}
}

func (s *imageService) SaveImageDetails(title string, tags string, id uuid.UUID, ctx context.Context) error {
	var image = repository.SaveImageParams{
		ID:    pgtype.UUID{Bytes: id, Valid: true},
		Tags:  ToText(tags),
		Title: ToText(title),
	}

	_, err := s.repository.SaveImage(ctx, image)
	return err
}

func ToText(value string) pgtype.Text {
	var val pgtype.Text
	val.Scan(value)
	return val
}

func (s *imageService) ListImages(ctx context.Context) ([]repository.Image, error) {
	return s.repository.ListImages(ctx)
}

func (s *imageService) SearchImages(tag string, ctx context.Context) ([]repository.Image, error) {
	log.Printf("im in search %s ", tag)
	images, err := s.repository.ListImages(ctx)

	if err != nil {
		return nil, err
	}

	var filtered []repository.Image
	for _, img := range images {
		if strings.Contains(img.Tags.String, tag) {
			filtered = append(filtered, img)
		}
	}

	return filtered, nil
}
