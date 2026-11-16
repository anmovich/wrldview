package transport

import (
	"net/http"

	"github.com/anmovich/wrldview/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	serv service.Service
}

func New(serv service.Service) *Handler {
	return &Handler{
		serv: serv,
	}
}

func (h Handler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "ping ok")
}
