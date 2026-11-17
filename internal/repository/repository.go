package repository

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Conn *pgx.Conn
}

func New(ctx context.Context) *Repository {
	conn, err := pgx.Connect(ctx, os.Getenv("DBLINK"))
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	return &Repository{
		Conn: conn,
	}
}

func NewTest() *Repository {
	return &Repository{}
}

func (r Repository) NewBook() {

}
