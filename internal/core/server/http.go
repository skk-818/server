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
		server      *http.Server
		cfg         *config.HTTPServer
		initManager []InitManager
		EngineProvider
	}

	EngineProvider interface {
		Engine() *gin.Engine
	}

	InitManager interface {
		InitIfNeeded() error
	}
)

func NewHTTPServer(engine EngineProvider, cfg *config.HTTPServer, manager []InitManager) *HTTPServer {
	return &HTTPServer{
		server: &http.Server{
			Addr:         fmt.Sprintf(":%s", cfg.Addr),
			Handler:      engine.Engine(),
			ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		},
		initManager:    manager,
		EngineProvider: engine,
	}
}

func (s *HTTPServer) Start() error {
	if len(s.initManager) > 0 {
		for _, m := range s.initManager {
			if err := m.InitIfNeeded(); err != nil {
				return err
			}
		}
	}
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
