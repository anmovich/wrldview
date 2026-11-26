package service

import (
	"github.com/anmovich/wrldview/internal/repository"
	"github.com/anmovich/wrldview/pkg/storage"
)

type Service struct {
	Repo    repository.Repository
	storage storage.Storage
}

func New(repo repository.Repository, storage storage.Storage) *Service {
	return &Service{
		Repo:    repo,
		storage: storage,
	}
}

func NewTest(repo repository.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}
