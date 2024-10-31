package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/natserract/toktik/pkg/config"
	echoHttp "github.com/natserract/toktik/pkg/http"
	echoHttpOptions "github.com/natserract/toktik/pkg/http/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load .env %v\n", err)
		os.Exit(1)
	}

	cfg := config.GetConfig()

	// Shared ctx
	e := echoHttp.NewEchoHttpServer(&echoHttpOptions.EchoHttpOptions{
		Port: ":" + cfg.Port,
		Host: cfg.Host,
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start server
	go func() {
		// https://dev.to/mokiat/proper-http-shutdown-in-go-3fji
		// https://echo.labstack.com/docs/cookbook/graceful-shutdown
		if err := e.RunHttpServer(); !errors.Is(
			err,
			http.ErrServerClosed,
		) {
			// do a fatal for going to OnStop process
			log.Fatalf(
				"(EchoHttpServer.RunHttpServer) error in running server: {%v}",
				err,
			)
		}
		log.Println("Stopped serving new connections.")
	}()
	e.GetEchoInstance().Logger.Infof(
		"Serving on Host:{%s} Http PORT: {%s}",
		e.Cfg().Host,
		e.Cfg().Port,
	)

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.GracefulShutdown(ctx); err != nil {
		e.GetEchoInstance().Logger.Errorf("error shutting down echo server: %v", err)
	} else {
		e.GetEchoInstance().Logger.Info("echo server shutdown gracefully")
	}
}
