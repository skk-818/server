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
