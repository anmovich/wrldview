package repository

import (
	"context"

	"github.com/anmovich/wrldview/models"
)

type Book interface {
	NewBook(book models.Book) (int, error)
	CheckID(id int) error
	PatchBook(book models.UpdateBook, id int) error
	GetBook(ctx context.Context, bookId int) (*models.Book, error)
}

type Cover interface {
	GETCover(bookId int) (string, error)
	DeleteCover(ctx context.Context, bookId int) error
}

type Repository interface {
	Book
	Cover
}
