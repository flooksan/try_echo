package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/flooksan/try_echo/app/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := handler.Handler{}

	// Routes
	e.GET("/", h.Hello)
	e.GET("/err", h.ThrowErr)

	// Anonymouse Func
	func() {
		fmt.Printf("Kloof start server!")
	}()

	// Graceful shuttdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(":1415"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("// ----- Shutdown Service ----- // ")
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
