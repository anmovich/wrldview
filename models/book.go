package models

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}
