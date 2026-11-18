package repository

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Conn *pgx.Conn
	ctx  context.Context
}

func New(ctx context.Context) *Repository {
	conn, err := pgx.Connect(ctx, os.Getenv("DBURL"))
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	return &Repository{
		ctx:  ctx,
		Conn: conn,
	}
}

func NewTest() *Repository {
	return &Repository{}
}

func (r Repository) NewBook() error {
	_, err := r.Conn.Exec(r.ctx, "DROP TABLE atata;")

	if err != nil {
		return err
	}
	return nil
}
