package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/update_me/update/internal/config"
	"github.com/update_me/update/server"
	"github.com/update_me/update/services/greetings"

	"github.com/caarlos0/env"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.New()
	if err := env.Parse(cfg); err != nil {
		msg := fmt.Sprintf("Error parsing config from environment: %s", err.Error())
		log.Fatalf(msg)
	}

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

	// create a listener on TCP port 8080
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		sugar.Fatalf("Failed to listen: %v", err)
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// create a server instance
	s := server.Server{
		Log: sugar,
		Cfg: cfg,
	}

	greetings.RegisterGreeterServiceServer(grpcServer, &s)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	sugar.Info("Starting gRPC server on port 8080")

	go func() {
		// start the server
		if err := grpcServer.Serve(lis); err != nil {
			msg := fmt.Sprintf("Failed to serve (Error: %s)", err.Error())
			sugar.Fatalf(msg)
		}
	}()

	sugar.Info("Starting http server on port 8000")
	if err := runHTTPServer(); err != nil {
		sugar.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	sugar.Info("Gracefully stopping gRPC server")
	grpcServer.GracefulStop()

	sugar.Info("Shutting down")
	os.Exit(0)
}

func runHTTPServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := greetings.RegisterGreeterServiceHandlerFromEndpoint(ctx, mux, "localhost:8080", opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8000", mux)
}
