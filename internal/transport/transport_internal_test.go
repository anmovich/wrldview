package transport

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anmovich/wrldview/internal/repository"
	"github.com/anmovich/wrldview/internal/service"
	"github.com/stretchr/testify/assert"
)

var s Server = NewServer(New(*service.NewTest(*repository.NewTest())))

func TestAPI_Ping(t *testing.T) {
	router := s.Setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, "ping ok", w.Body.String())
}
