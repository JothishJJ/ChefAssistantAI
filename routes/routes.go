package routes

import (
	"log"
	"net/http"

	"github.com/JothishJJ/ChefAssistantAI/gemini"
	"github.com/google/generative-ai-go/genai"
	"github.com/labstack/echo/v4"
)

type Page struct {
	Name string
}

type Message struct {
	Message interface{}
}

func sendMessage(message string) genai.Part {

	cs, ctx := gemini.GenerateChat()
	resp, err := gemini.SendMessage(cs, ctx, message)
	if err != nil {
		log.Fatal(err)
	}

	responseText := resp.Candidates[0].Content.Parts[0]
	return responseText
}

/**
* function for handling Routes
 */
func Routes(e *echo.Echo) {

	// Home Route
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", Page{Name: "The no.1 AI tool for all your needs!"})
	})

	// Dashboard Routes
	e.GET("/dashboard", func(c echo.Context) error {
		return c.Render(http.StatusOK, "dashboard", Page{Name: "Dashboard"})
	})

	e.POST("/generate", func(c echo.Context) error {

		message := c.FormValue("message")
		resp := sendMessage(message)

		return c.Render(http.StatusOK, "messages", Message{Message: resp})
	})
}
