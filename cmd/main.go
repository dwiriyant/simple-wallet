package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"simple-wallet/internal/interfaces/api"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Execute() {
	e := echo.New()

	api.API(e)

	e.Logger.SetLevel(log.INFO)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(":8081"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
