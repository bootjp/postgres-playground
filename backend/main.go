package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/koyashiro/rdbms-playground/backend/client"
	"github.com/koyashiro/rdbms-playground/backend/handler"
	"github.com/koyashiro/rdbms-playground/backend/service"
)

const port = "1323"

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// TODO: replace DI

	// Clients
	containerClient := client.NewContainerClient()
	rdbmsClient := client.NewRDBMSClient()

	// Services
	playgroundService := service.NewWorkspaceService(containerClient, rdbmsClient)

	// Handlers
	playgroundHandler := handler.NewWorkspacesHandler(playgroundService)

	// Routes
	e.GET("/workspaces", playgroundHandler.Index)
	e.GET("/workspaces/:id", playgroundHandler.Show)
	e.POST("/workspaces", playgroundHandler.Create)
	e.DELETE("/workspaces/:id", playgroundHandler.Delete)
	e.POST("/workspaces/:id/query", playgroundHandler.Query)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
