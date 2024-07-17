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
	CompanyAdminRoutes(e)
}

func AuthRequiredRoutes(e *echo.Echo) {
	// Auth Required Routes:
	e.GET("/user", handlers.GetUser, middlewares.Auth)
	// Company Related Routes:
	e.POST("/company", handlers.CreateCompany, middlewares.Auth)
	// NGO Related Routes:
	e.POST("/ngo", handlers.CreateNGO, middlewares.Auth)
	e.GET("/user/ngo", handlers.GetUserNGO, middlewares.Auth)
}

func CompanyAdminRoutes(e *echo.Echo) {
	e.POST("/company/material", handlers.CreateMaterial, middlewares.CompanyAdminAuth)
	e.GET("/company/material", handlers.GetCompanyMaterial)
	e.GET("/user/company", handlers.GetUserCompany, middlewares.Auth)
}

func AdminRoutes(e *echo.Echo) {
	// In production all of those routes should apply the middleware adminAuth
	e.GET("/users", handlers.GetUsers)
	// Companies related routes:
	e.GET("/companies", handlers.GetCompanies)
	e.GET("/companies/admins", handlers.GetCompaniesAdmins)
	// NGO Related Routes:
	e.GET("/ngos", handlers.GetNGOs)
	e.GET("/ngos/admins", handlers.GetNGOsAdmins)
}
