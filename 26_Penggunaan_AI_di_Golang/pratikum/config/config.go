package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

var client *openai.Client
var apiKey string
var messages []openai.ChatCompletionMessage
var model string

// Bot is an interface for the chatbot.
type Bot interface {
	GetClient() *openai.Client
	GetPrompt() (string, []openai.ChatCompletionMessage)
}

type bot struct{}

// NewBot creates a new instance of the chatbot.
func NewBot() Bot {
	return &bot{}
}

// GetClient returns the chatbot's OpenAI client.
func (b *bot) GetClient() *openai.Client {
	if client != nil {
		return client
	}

	if apiKey == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Cannot load env file. Err: %s", err)
		}

		apiKey = os.Getenv("API_KEY")
	}

	client = openai.NewClient(apiKey)
	return client
}

// GetPrompt returns the chatbot's model and initial prompt.
func (b *bot) GetPrompt() (string, []openai.ChatCompletionMessage) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	if messages == nil {
		messages = []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: `
				You are an assistant that provides laptop recommendations based on the following user-defined constraints:

				- Budget: [User's Budget in Indonesian Rupiah in ones]
				- Preferred Brand: [User's Preferred Brand]
				- Purpose: [User's Intended Use]
				- RAM: [User's Preferred RAM in Gigabytes]
				- CPU Cores: [User's Preferred Number of CPU Cores]
				- Display Dimension: [User's Preferred Display Size in inches]

				All user defined constraints must be met with no exception, and I also need multiple laptop recommendations, each with comprehensive information including brand, price, and a short description of the laptop's features, capabilities, and suitability for the specified constraints.
				`,
			},
		}
	}
	return model, messages
}