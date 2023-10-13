package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang-ai/config"
	"golang-ai/dto"

	"github.com/sashabaranov/go-openai"
)

func getCompletionFromMessages(
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
	model string,
) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo // see another model options: https://platform.openai.com/docs/models
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}


type PromptUseCase interface {
	GetLaptopRecommendation(data *dto.LaptopSpecRequest) (string, error)
}

type promptUseCase struct {
	bot config.Bot
}


func NewPromptUseCase(bot config.Bot) *promptUseCase {
	return &promptUseCase{bot}
}

func (p *promptUseCase) GetLaptopRecommendation(data *dto.LaptopSpecRequest) (string, error) {
	userInput, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	
	ctx := context.Background()
	model, messages := p.bot.GetPrompt()
	client := p.bot.GetClient()

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: string(userInput),
	})

	resp, err := getCompletionFromMessages(ctx, client, messages, model)

	if err != nil {
		return "", errors.New(fmt.Sprintf("error when sending request to API: %s", err))
	}

	answer := openai.ChatCompletionMessage{
		Role:    resp.Choices[0].Message.Role,
		Content: resp.Choices[0].Message.Content,
	}
	messages = append(messages, answer)

	return answer.Content, nil
}