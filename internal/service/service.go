package service

import (
	"errors"
	"log/slog"

	"github.com/anmovich/wrldview/internal/errs"
	"github.com/anmovich/wrldview/internal/repository"
	"github.com/anmovich/wrldview/internal/validator"
	"github.com/anmovich/wrldview/models"
)

type Service struct {
	Repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func NewTest(repo repository.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

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
