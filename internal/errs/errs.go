package errs

import "errors"

var (
	ErrBookWrongPages  = errors.New("Pages cant be smaller or equal 0")
	ErrBookWrongTitle  = errors.New("Title cant be null")
	ErrBookWrongAuthor = errors.New("Author cant be null")
	ErrBookWrongID     = errors.New("ID cant be <= 0")
)
