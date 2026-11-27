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

	book := router.Group("/book")
	{
		book.POST("/", s.h.NewBook)
		book.PATCH("/", s.h.PATCHBook)
		book.GET("/", s.h.GetBook)
		book.POST("/cover", s.h.UploadCover)
		book.GET("/cover", s.h.GETCover)
		book.DELETE("/cover", s.h.DeleteCover)
	}

	return router
}
