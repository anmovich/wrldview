package storage

import (
	"context"
	"io"

	"github.com/anmovich/wrldview/models"
)

type Storage interface {
	Upload(ctx context.Context, name string, body io.Reader, contentType string) error
	Delete(ctx context.Context, name string) error
	Download(ctx context.Context, name string) (*models.CoverS3, error)
}
