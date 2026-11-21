package validator

import (
	"github.com/anmovich/wrldview/internal/errs"
	"github.com/anmovich/wrldview/models"
)

func BookValidate(book models.Book) error {
	if book.Pages <= 0 {
		return errs.ErrBookWrongPages
	} else if book.Title == "" {
		return errs.ErrBookWrongTitle
	} else if book.Author == "" {
		return errs.ErrBookWrongAuthor
	}
	return nil
}

func BookValidatePatch(book models.Book) error {
	if book.ID <= 0 {
		return errs.ErrBookWrongID
	}

	return nil
}
