package adapters

import (
	"github.com/sashabaranov/go-openai"

	"github.com/eguezgustavo/ai-interaction/domain"
)

type OpenAIImpl struct {
}

func (openAI *OpenAIImpl) NewClient(authToken string) OpenAIClient {
	return openai.NewClient(authToken)
}

func GetModelWithNoAIBehind() interactor.AI {
	return NoAIInteractor{}
}

func GetOpenAIGPTModel() interactor.AI {

	return ChatGPTInteractor{
		openai: &OpenAIImpl{},
	}
}
