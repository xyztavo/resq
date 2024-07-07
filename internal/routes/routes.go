package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/xyztavo/resq/internal/handlers"
	"github.com/xyztavo/resq/internal/middlewares"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", handlers.HelloWorld)
	e.POST("/user", handlers.CreateUser)
	e.POST("/auth", handlers.Auth)

	AdminRoutes(e)
	AuthRequiredRoutes(e)
}

func AuthRequiredRoutes(e *echo.Echo) {
	// Auth Required Routes
	e.GET("/user", handlers.GetUser, middlewares.Auth)
	e.PATCH("/user/ngo/admin", handlers.UpdateUserNGOAdmin)
	e.PATCH("/user/company/admin", handlers.UpdateUserCompanyAdmin)
	// Company Related Routes:
	e.POST("/company", handlers.CreateCompany)
	e.GET("/user/company", handlers.GetUserCompany)
	// Company admin routes
	e.GET("/user/company/admin", handlers.GetUserCompanyAdmin)
}

func AdminRoutes(e *echo.Echo) {
	// In production all of those routes should apply the middleware adminAuth
	e.GET("/companies", handlers.GetCompanies)
	e.GET("/users", handlers.GetUsers)
}
