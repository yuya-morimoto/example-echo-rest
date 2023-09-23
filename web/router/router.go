package router

import (
	"example-rest/handler"

	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {

	handler.AuthRouter(e)

}
