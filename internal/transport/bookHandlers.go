package transport

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/anmovich/wrldview/internal/validator"
	"github.com/anmovich/wrldview/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) NewBook(c *gin.Context) {
	var book models.Book

	if err := json.NewDecoder(c.Request.Body).Decode(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		slog.Error("Error in NewBook, " + err.Error())
		return
	}
	if err := validator.BookValidate(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		slog.Info(err.Error())
		return
	}
	fmt.Println(book)
}
