package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	Engine *gin.Engine
	server *http.Server
}

type EngineProvider interface {
	Engine() *gin.Engine
}

func NewHTTPServer(engine EngineProvider) *HTTPServer {
	return &HTTPServer{
		Engine: engine.Engine(),
		server: &http.Server{
			Addr:         ":8080",
			Handler:      engine.Engine(),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *HTTPServer) Start() error {
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
