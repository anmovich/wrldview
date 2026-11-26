package service

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"mime"
	"net/http"

	"github.com/anmovich/wrldview/internal/errs"
	"github.com/anmovich/wrldview/internal/validator"
	"github.com/anmovich/wrldview/models"
	"github.com/google/uuid"
)

func (s Service) NewBook(book models.Book) (int, error) {
	id, err := s.Repo.NewBook(book)
	if err != nil {
		slog.Error(err.Error())
		return 0, errors.New("error while add new book")
	}
	return id, nil
}

func (s Service) PatchBook(book models.UpdateBook, id int) error {
	err := validator.IdValidate(id)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	err = s.Repo.CheckID(id)
	if err != nil {
		slog.Error(err.Error())
		return errs.ErrNoID
	}

	err = s.Repo.PatchBookAuthor(book, id)
	if err != nil {
		slog.Error(err.Error())
		return errs.ErrNoRowsFind
	}
	return nil
}

func (s Service) UploadCover(body io.ReadSeeker) error {

	buffer := make([]byte, 512)
	n, err := body.Read(buffer)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	contentType := http.DetectContentType(buffer[:n])
	if _, err := body.Seek(0, io.SeekStart); err != nil {
		return err
	}
	slog.Info(contentType)
	err = validator.ValidateContentType(contentType)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	ext, err := mime.ExtensionsByType(contentType)
	if err != nil || len(ext) == 0 {
		return errors.New("unknown file extension")
	}
	nameUUID := uuid.New()
	name := nameUUID.String() + ext[0]
	slog.Info(name)
	err = s.storage.Upload(context.Background(), name, body, contentType)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	return nil
}
