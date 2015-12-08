package main

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/wangyibin/akita"
)

func accessible(c *echo.Context) error {
	return c.String(http.StatusOK, "No auth required for this route.\n")
}

func restricted(c *echo.Context) error {
	return c.String(http.StatusOK, "Access granted with JWT.\n")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Unauthenticated route
	e.Get("/", accessible)

	// Restricted group
	r := e.Group("/restricted")
	akita.Config.Secret = []byte("Secret cannot tell")
	akita.Config.DataSourceDriven = "mysql"
	akita.Config.DataSourceURL = "root:root@/akita_test"

	r.Use(akita.JWTAuth)
	r.Get("", restricted)

	// Start server
	e.Run(":1323")
}
