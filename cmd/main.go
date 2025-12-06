package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/internal/di"
	"syscall"
	"time"
)

const configFile = "./etc/config.yaml"

// @title 系统管理API
// @version 1.0
// @description 系统管理后台API接口文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9999
// @BasePath /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	app, err := di.InitApp(configFile)
	if err != nil {
		log.Fatalf("init error: %v", err)
	}

	go func() {
		if err := app.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown error: %v", err)
	}
}
