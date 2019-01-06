package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {

	//runWithPointer()
	runWithLocalVariable()
}

func runWithPointer() {
	e := echo.New()

	e.GET("/say", sayHi)

	e.Logger.Fatal(e.Start(":1331"))
}

func runWithLocalVariable() {
	var e echo.Echo
	e = *echo.New()
	e.GET("/say", sayHi)

	e.Logger.Fatal(e.Start(":1332"))
}

func sayHi(e echo.Context) error {
	return e.JSON(http.StatusOK, "hi")
}