package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 5 * time.Second,
	}))

	e.GET("/", timeoutFn)

	e.Logger.Fatal(e.Start(":8080"))
}

func timeoutFn(c echo.Context) error {
	time.Sleep(3 * time.Second)
	return c.String(http.StatusOK, "timeoutFn")
}
