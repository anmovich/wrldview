package service

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"mime"
	"net/http"

	"github.com/anmovich/wrldview/internal/validator"
	"github.com/anmovich/wrldview/models"
	"github.com/google/uuid"
)

func (s Service) UploadCover(body io.ReadSeeker, bookId int) error {
	//Create buffer for reader
	buffer := make([]byte, 512)
	n, err := body.Read(buffer)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	//Detect content type
	contentType := http.DetectContentType(buffer[:n])
	if _, err := body.Seek(0, io.SeekStart); err != nil {
		return err
	}
	slog.Info(contentType)
	//Validate content type
	err = validator.ValidateContentType(contentType)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	//Create file extention
	ext, err := mime.ExtensionsByType(contentType)
	if err != nil || len(ext) == 0 {
		return errors.New("unknown file extension")
	}
	//Create filename
	nameUUID := uuid.New()
	name := nameUUID.String() + ext[0]
	slog.Info(name)
	//Upload to storage
	err = s.storage.Upload(context.Background(), name, body, contentType)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	//Upload to repository
	book := models.UpdateBook{
		Cover: &name,
	}
	s.Repo.PatchBook(book, bookId)
	return nil
}

func (s Service) GetCover(bookId int) (*models.CoverS3, error) {
	//Get cover name from repository
	cover, err := s.Repo.GETCover(bookId)
	if err != nil {
		return nil, err
	}

	//Get item from s3
	object, err := s.storage.Download(context.Background(), cover)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	return object, nil
}

func (s Service) DeleteCover(bookId int) error {
	cover, err := s.Repo.GETCover(bookId)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	err = s.Repo.DeleteCover(context.Background(), bookId)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	err = s.storage.Delete(context.Background(), cover)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	return nil
}
