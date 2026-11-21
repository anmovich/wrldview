package service

import (
	"errors"
	"log/slog"

	"github.com/anmovich/wrldview/internal/repository"
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
