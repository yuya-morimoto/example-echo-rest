package main

import (
	"example-rest/presentation/router"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// Setup router
	router.Init(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
