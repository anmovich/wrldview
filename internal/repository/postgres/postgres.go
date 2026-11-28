package postgres

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/anmovich/wrldview/models"
	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	Conn *pgx.Conn
	ctx  context.Context
}

func New(ctx context.Context) *Postgres {
	conn, err := pgx.Connect(ctx, os.Getenv("DBURL"))
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	return &Postgres{
		ctx:  ctx,
		Conn: conn,
	}
}

func NewTest() *Postgres {
	return &Postgres{}
}

func (r Postgres) NewBook(book models.Book) (int, error) {
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

func (r Postgres) CheckID(id int) error {
	sqlreq := `SELECT id FROM books WHERE id = $1`
	row := r.Conn.QueryRow(r.ctx, sqlreq, id)

	var queryId int
	err := row.Scan(&queryId)
	if err != nil {
		return err
	}

	fmt.Println(queryId)

	if queryId != id {
		return errors.New("Smth wrong")
	}
	return nil

}

func (r Postgres) PatchBook(book models.UpdateBook, id int) error {
	query := `UPDATE books SET `

	counter := 1
	args := []any{}
	if book.Author != nil {
		query += fmt.Sprintf("author = $%d, ", counter)
		args = append(args, book.Author)
		counter++
	}

	if book.Title != nil {
		query += fmt.Sprintf("title = $%d, ", counter)
		args = append(args, book.Title)
		counter++
	}

	if book.Pages != nil {
		query += fmt.Sprintf("pages = $%d, ", counter)
		args = append(args, book.Pages)
		counter++
	}

	if book.Cover != nil {
		query += fmt.Sprintf("cover = $%d, ", counter)
		args = append(args, book.Cover)
		counter++
	}

	if len(args) == 0 {
		return errors.New("No any fields to update")
	}

	query = strings.TrimSuffix(query, ", ")

	query += fmt.Sprintf(" WHERE id = $%d", counter)

	slog.Info(query)
	args = append(args, id)
	tag, err := r.Conn.Exec(r.ctx, query, args...)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return errors.New("No rows find")
	}

	return nil
}

func (r Postgres) GetBook(ctx context.Context, bookId int) (*models.Book, error) {
	query := `SELECT id, title, author, pages, cover FROM books WHERE id=$1`
	var book models.Book
	if err := r.Conn.QueryRow(ctx, query, bookId).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Pages,
		&book.Cover,
	); err != nil {
		return nil, err
	}
	return &book, nil

}

func (r Postgres) GETCover(bookId int) (string, error) {
	query := `SELECT cover FROM public.books WHERE id=$1`
	var cover string
	link := r.Conn.QueryRow(r.ctx, query, bookId)
	if err := link.Scan(&cover); err != nil {
		slog.Error(err.Error())
		return "", err
	}
	return cover, nil
}

func (r Postgres) DeleteCover(ctx context.Context, bookId int) error {

	query := `UPDATE books SET cover = NULL WHERE id=$1`
	if _, err := r.Conn.Exec(ctx, query, bookId); err != nil {
		slog.Error(err.Error())
		return err
	}
	return nil
}
