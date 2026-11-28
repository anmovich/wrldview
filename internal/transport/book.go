package transport

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/anmovich/wrldview/internal/validator"
	"github.com/anmovich/wrldview/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) NewBook(c *gin.Context) {
	var book models.Book

	if err := json.NewDecoder(c.Request.Body).Decode(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		slog.Error("handler Error in NewBook, " + err.Error())
		return
	}
	if err := validator.BookValidate(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		slog.Info(err.Error())
		return
	}

	id, err := h.serv.NewBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	c.JSON(http.StatusCreated, gin.H{
		"ID": id,
	})
	fmt.Println(book)
}

func (h Handler) PATCHBook(c *gin.Context) {

	idstring := c.Request.URL.Query().Get("id")
	slog.Info(idstring)
	id, err := strconv.Atoi(idstring)

	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong id type",
		})
		return
	}

	var book models.UpdateBook

	err = json.NewDecoder(c.Request.Body).Decode(&book)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	err = h.serv.PatchBook(book, id)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h Handler) GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Request.URL.Query().Get("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "no id",
		})
		return
	}
	book, err := h.serv.GetBook(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "no such id in base",
		})
		return
	}

	c.JSON(http.StatusAccepted, book)
}
