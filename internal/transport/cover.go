package transport

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const maxFileSize = 5 * 1024 * 1024

func (h Handler) UploadCover(c *gin.Context) {
	id, err := strconv.Atoi(c.Request.URL.Query().Get("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "No id",
		})
		return
	}
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		slog.Error(err.Error())
		return
	}
	if fileHeader.Size < maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "size need to be < 5MiB",
		})
		return
	}
	if fileHeader == nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		slog.Error(errors.New("File header nil").Error())
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		slog.Error(err.Error())
		return
	}
	err = h.serv.UploadCover(file, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Unsupported type",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func (h Handler) GETCover(c *gin.Context) {
	id, err := strconv.Atoi(c.Request.URL.Query().Get("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "no id",
		})
		return
	}
	cover, err := h.serv.GetCover(id)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.DataFromReader(
		http.StatusAccepted,
		*cover.Size,
		*cover.ContentType,
		cover.Body,
		nil,
	)
}

func (h Handler) DeleteCover(c *gin.Context) {
	id, err := strconv.Atoi(c.Request.URL.Query().Get("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "no id",
		})
		return
	}
	err = h.serv.DeleteCover(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{})
}
