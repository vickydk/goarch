package http

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"goarch/pkg/interface/container"

	"github.com/labstack/echo/v4"
)

func StartHttpService(cont *container.Container) {
	server := echo.New()
	server.HideBanner = true

	setupMiddleware(server, cont.Config)
	SetupRouter(server, SetupHandlers(cont))

	// start server
	go func() {
		if err := server.Start(cont.Config.AppAddress()); err != nil {
			server.Logger.Print(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}
