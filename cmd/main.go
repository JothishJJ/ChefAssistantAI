package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func createTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func createServer() *echo.Echo {
	e := echo.New()

	/**
	* Middleware Logger for debugging
	*
	* ! Disable this in production
	 */
	e.Use(middleware.Logger())

	// RateLimitter
	rateLimitter := middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20))
	e.Use(rateLimitter)

	// Templating Renderer
	e.Renderer = createTemplate()
	return e
}

type Page struct {
	Name string
}

func main() {

	e := createServer()

	// TODO: Make Routes on seperate folders
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", Page{Name: "Chef's Assistant AI"})
	})

	e.GET("/dashboard", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Development is still underway")

	})

	// Turns out you to run this after every handler
	e.Logger.Fatal(e.Start(":8080"))
}
