package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/caarlos0/env"
	"github.com/update/update_me/internal/config"
	"github.com/update/update_me/internal/handler"

	"go.uber.org/zap"
)

// Secrets
var ()

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		msg := fmt.Sprintf("Can't initialize zap logger Error: %s", err.Error())
		log.Fatalf(msg)
	}

	defer func() {
		if err := logger.Sync(); err != nil {
			log.Printf("Unable to flush buffered log entries: %v", err)
		}
	}()
	sugar := logger.Sugar()

	cfg := config.New()
	if err := env.Parse(cfg); err != nil {
		sugar.Fatalf("parsing config from environment: %s", err.Error())
	}

	handler := handler.New(sugar, cfg)
	addr := net.JoinHostPort(cfg.Host, cfg.Port)

	server := &http.Server{
		Handler:        handler,
		Addr:           addr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	sugar.Info("starting web server")

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				sugar.Fatalf("web server shutdown unexpectedly: %s", err.Error())
			}

			sugar.Errorf("web server shutdown: %s", err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	sugar.Info("stopping web server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		sugar.Fatalf("stopping web server: %s", err.Error())
	}

	sugar.Info("web server shutdown complete")
}
