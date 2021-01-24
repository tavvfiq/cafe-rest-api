package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	c "github.com/tavvfiq/cafe-rest-api/controllers"
)

// Start the router
func Start() {
	e := echo.New()
	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}\turi=${uri}\t\tstatus=${status}\n",
	}))
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Routes
	e.POST("/auth/register", c.RegisterHandler)
	e.POST("/auth/login", c.LoginHandler)
	e.GET("/menu", c.GetMenuHandler)

	// Logger
	e.Logger.Fatal(e.Start(":8001"))
}
