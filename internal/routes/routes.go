package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/xyztavo/resq/internal/handlers"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", handlers.HelloWorld)
}
