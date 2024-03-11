package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type request struct {
	Name string `json:"name"`
}

func echoServer() {
	e := echo.New()
	e.GET("/small", func(c echo.Context) error {
		var r request
		err := c.Bind(&r)
		if err != nil {
			panic(err)
		}
		return c.NoContent(http.StatusOK)
	})
	e.GET("/medium", func(c echo.Context) error {
		var r request
		err := c.Bind(&r)
		if err != nil {
			panic(err)
		}
		return c.NoContent(http.StatusOK)
	})
	e.GET("/large", func(c echo.Context) error {
		var r request
		err := c.Bind(&r)
		if err != nil {
			panic(err)
		}
		return c.NoContent(http.StatusOK)
	})
	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
