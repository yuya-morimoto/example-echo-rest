package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthRouter(e *echo.Echo) {
	group := e.Group("/auth")
	group.GET("/", index)

}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
