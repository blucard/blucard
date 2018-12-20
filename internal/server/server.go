package server

import (
	"net/http"

	"github.com/blucard/blucard/internal/backend"
	"github.com/gin-gonic/gin"
)

// Server exposes HTTP endpoints
type Server struct {
	HTTP    *http.Server
	backend backend.Backend
}

// NewServer allocates and creates a new Server
func NewServer(backend backend.Backend) (*Server, error) {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	r.GET("/record/:uuid", backend.GetRecord)
	r.POST("/record/:uuid", backend.SetRecord)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	return &Server{
		HTTP:    srv,
		backend: backend,
	}, nil
}
