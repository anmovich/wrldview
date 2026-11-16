package transport

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port string
	h    *Handler
}

func NewServer(h *Handler) Server {
	return Server{
		Port: fmt.Sprintf(":%s", os.Getenv("PORT")),
		h:    h,
	}
}

func (s Server) Run() {
	router := s.Setup()
	router.Run(s.Port)
}

func (s Server) Setup() *gin.Engine {
	router := gin.Default()

	router.POST("/ping", s.h.Ping)

	return router
}
