package validator

import (
	"errors"

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

func IdValidate(id int) error {
	if id < 0 {
		return errors.New("Id cant be < 0")
	}
	return nil
}

func ValidateContentType(contentType string) error {
	approvedContentType := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"image/webp": true,
	}

	if !approvedContentType[contentType] {
		return errors.New("Content type not approved")
	}
	return nil
}
