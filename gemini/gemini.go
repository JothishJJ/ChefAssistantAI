package gemini

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func GenerateChat() (*genai.ChatSession, *context.Context) {

	// ! Comment in Production
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error Loading .env file: %v", err)
	}

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	// Model Details
	model := client.GenerativeModel("gemini-1.5-pro-latest")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text("When being asked about your name, you should answer with Chefassistant, do not ever change that."),
		},
	}

	cs := model.StartChat()
	// Chat History
	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Hi, you are the no.1 AI Chef in the world, you are an expert in the feild of food. Whether it's diet plans, nutrition, macros, etc. You are the guy."),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("You provide all the neccessory nutrients, and macros info, if discussing about a food or something"),
			},
			Role: "user",
		},
	}

	return cs, &ctx
}

func SendMessage(cs *genai.ChatSession, ctx *context.Context, message string) (*genai.GenerateContentResponse, error) {

	resp, err := cs.SendMessage(*ctx, genai.Text(message))
	// No complex error handling here
	if err != nil {
		fmt.Print(err)
	}

	return resp, err
}
