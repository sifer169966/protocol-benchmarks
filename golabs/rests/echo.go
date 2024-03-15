package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type request struct {
	Name string `json:"name"`
}

func echoServer() {
	e := echo.New()
	e.POST("/small", func(c echo.Context) error {
		var r request
		err := c.Bind(&r)
		if err != nil {
			panic(err)
		}
		return c.NoContent(http.StatusOK)
	})
	e.POST("/medium", func(c echo.Context) error {
		var r request
		err := c.Bind(&r)
		if err != nil {
			panic(err)
		}
		return c.NoContent(http.StatusOK)
	})
	e.POST("/large", func(c echo.Context) error {
		var r request
		err := c.Bind(&r)
		if err != nil {
			panic(err)
		}
		return c.NoContent(http.StatusOK)
	})
	exposePort := os.Getenv("HTTP_EXPOSE_PORT")
	if exposePort == "" {
		exposePort = "8080"
	}
	err := e.Start(":" + exposePort)
	if err != nil {
		log.Fatal(err)
	}
}
