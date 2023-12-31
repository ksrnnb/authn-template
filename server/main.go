package main

import (
	"github.com/ksrnnb/authn-template/handler"
	"github.com/ksrnnb/authn-template/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())
	e.Use(middleware.RepositoryMiddleware())

	// Unauthenticated Routes
	e.POST("/signin", handler.SignIn)

	// Authenticated Routes
	e.POST("/authenticated", handler.Authenticated, middleware.AuthMiddleware())
	e.POST("/signout", handler.SignOut, middleware.AuthMiddleware())

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}
