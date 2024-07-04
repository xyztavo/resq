package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/xyztavo/resq/configs"
	"github.com/xyztavo/resq/internal/database"
	"github.com/xyztavo/resq/internal/routes"
)

func main() {
	// Migrates and instanciate a db connection
	if err := database.Migrate(); err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	// Setups the routes
	routes.SetupRoutes(e)
	e.Logger.Fatal(e.Start(configs.GetPort()))
}
