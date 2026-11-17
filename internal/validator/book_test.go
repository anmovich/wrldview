package validator

import (
	"log/slog"
	"testing"

	"github.com/anmovich/wrldview/internal/errs"
	"github.com/anmovich/wrldview/models"
	"github.com/stretchr/testify/require"
)

type CaseBook struct {
	Name   string
	In     *models.Book
	ExpErr error
}

func TestValidateBookErr(t *testing.T) {
	cases := []CaseBook{
		{
			Name:   "Test pages",
			In:     models.NewBook("dkfj", "dfksfj", -13131),
			ExpErr: errs.ErrBookWrongPages,
		},
		{
			Name:   "Test author",
			In:     models.NewBook("dkfjs", "", 343),
			ExpErr: errs.ErrBookWrongAuthor,
		},
		{
			Name:   "Test title",
			In:     models.NewBook("", "dkfksaf", 338),
			ExpErr: errs.ErrBookWrongTitle,
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.Name, func(t *testing.T) {
			err := BookValidate(*tCase.In)
			slog.Info(err.Error())
			require.Error(t, err)
			require.EqualError(t, err, tCase.ExpErr.Error())
		})
	}
}
