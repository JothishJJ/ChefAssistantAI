package main

import (
	"html/template"
	"io"

	"github.com/JothishJJ/ChefAssistantAI/routes"
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
	* Serve static files
	 */
	e.Static("/", "public")

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

func main() {

	e := createServer()

	routes.Routes(e)

	// Turns out you to run this after every handler
	e.Logger.Fatal(e.Start(":8080"))
}
