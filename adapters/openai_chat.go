package adapters

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

const tokenEnvVarName = "OPENAI_TOKEN"

type OpenAIClient interface {
	CreateChatCompletion(ctx context.Context, request openai.ChatCompletionRequest) (response openai.ChatCompletionResponse, err error)
}

type OpenAI interface {
	NewClient(authToken string) OpenAIClient
}

type ChatGPTInteractor struct {
	openai OpenAI
}

func (interactor ChatGPTInteractor) ProcessPrompt(prompt string) string {
	token := os.Getenv(tokenEnvVarName)
	client := interactor.openai.NewClient(token)
	request := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}
	response, _ := client.CreateChatCompletion(context.Background(), request)
	return response.Choices[0].Message.Content
}
