package repository

import (
	"context"
	"errors"
	"log/slog"
	"os"

	"github.com/anmovich/wrldview/models"
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

func (r Repository) NewBook(book models.Book) (int, error) {
	sqlreq := `INSERT INTO books(title, author, pages) VALUES ($1, $2, $3) RETURNING id;`
	var id int
	row := r.Conn.QueryRow(r.ctx, sqlreq, book.Title, book.Author, book.Pages)
	if err := row.Scan(&id); err == pgx.ErrNoRows {
		slog.Error(err.Error())
		return 0, errors.New("Error while add new book")
	}
	slog.Info("succes add new book")
	return id, nil
}

func (r Repository) CheckID(book models.Book) error {
	sqlreq := `SELECT id FROM books WHERE id = $1`

}
