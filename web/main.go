package main

import (
	"example-rest/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// Setup router
	router.Setup(e)

	e.Logger.Fatal(e.Start(":8080"))

}
