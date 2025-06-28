package server

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"server/internal/core/config"
	"server/internal/core/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	Engine *gin.Engine
	server *http.Server
	cfg    *config.HTTPServer
	logger logger.Logger
}

type EngineProvider interface {
	Engine() *gin.Engine
}

func NewHTTPServer(logger logger.Logger, engine EngineProvider, cfg *config.HTTPServer) *HTTPServer {
	return &HTTPServer{
		Engine: engine.Engine(),
		server: &http.Server{
			Addr:         fmt.Sprintf(":%s", cfg.Addr),
			Handler:      engine.Engine(),
			ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		},
		logger: logger,
	}
}

func (s *HTTPServer) Start() error {
	s.logger.Info("🚀 HTTP server starting...",
		zap.String("addr", s.server.Addr),
	)
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
