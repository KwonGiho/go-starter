package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {

	e := echo.New()

	e.GET("/say", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hi")
	})

	e.Logger.Fatal(e.Start(":1331"))
}
