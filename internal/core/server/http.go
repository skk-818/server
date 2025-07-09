package server

import (
	"context"
	"fmt"
	"net/http"
	"server/internal/core/config"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	HTTPServer struct {
		server *http.Server
		cfg    *config.HTTPServer
		InitManager
		EngineProvider
	}

	EngineProvider interface {
		Engine() *gin.Engine
		Initializer() error
	}

	InitManager interface {
		InitIfNeeded() error
	}
)

func NewHTTPServer(engine EngineProvider, cfg *config.HTTPServer, manager InitManager) *HTTPServer {
	return &HTTPServer{
		server: &http.Server{
			Addr:         fmt.Sprintf(":%s", cfg.Addr),
			Handler:      engine.Engine(),
			ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		},
		InitManager:    manager,
		EngineProvider: engine,
	}
}

func (s *HTTPServer) Start() error {
	if err := s.InitManager.InitIfNeeded(); err != nil {
		return err
	}
	if err := s.EngineProvider.Initializer(); err != nil {
		return err
	}
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
