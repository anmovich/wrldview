package storage

import (
	"context"
	"io"
)

type Storage interface {
	Upload(ctx context.Context, name string, body io.Reader, contentType string) error
	Delete(ctx context.Context, name string) error
}
