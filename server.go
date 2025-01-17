package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/marcArisen/assessment/api"
)

func getPort() string {

	port := os.Getenv("PORT")
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", port)

	return port
}

func gracefulShutdown(e *echo.Echo) {

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	fmt.Println("SEVER IS SHUTTING DOWN...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Println("BYE BYE.")
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

func createRoute(e *echo.Echo) {

	e.POST("/expenses", api.CreateExpenses)
	e.GET("/expenses/:id", api.GetByIdExpenses)
	e.GET("/expenses", api.GetAllExpenses)
	e.PUT("/expenses/:id", api.UpdateExpenses)

}

func main() {

	port := getPort()

	e := echo.New()

	createRoute(e)

	go func() {
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	gracefulShutdown(e)

}
