package router

import (
	"example-rest/presentation/handler"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	handler.AuthRouter(e)

}
