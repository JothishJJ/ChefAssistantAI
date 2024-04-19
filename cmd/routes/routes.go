package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Page struct {
	Name string
}

/**
* function for handling Routes
 */
func Routes(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", Page{Name: "The no.1 AI tool for all your needs!"})
	})

	e.GET("/dashboard", func(c echo.Context) error {
		return c.Render(http.StatusOK, "dashboard", Page{Name: "Welcome to Dashboard"})
	})
}
