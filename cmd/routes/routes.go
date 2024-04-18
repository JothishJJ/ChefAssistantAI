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
		return c.Render(http.StatusOK, "index", Page{Name: "Chef's Assistant AI"})
	})

	e.GET("/dashboard", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Development is still underway")
	})

}
