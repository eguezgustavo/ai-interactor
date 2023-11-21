package adapters

import (
	"context"
	"os"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var openAIToken string

type MockOpenAILibrary struct {
	mock.Mock
}

func (mockOpenAI *MockOpenAILibrary) NewClient(authToken string) OpenAIClient {
	valuesToReturn := mockOpenAI.Called(authToken)
	return valuesToReturn.Get(0).(OpenAIClient)
}

type MockOpenAIClient struct {
	mock.Mock
}

func (mockclient *MockOpenAIClient) CreateChatCompletion(ctx context.Context, request openai.ChatCompletionRequest) (response openai.ChatCompletionResponse, err error) {
	shouldReturn := mockclient.Called(ctx, request)
	return shouldReturn.Get(0).(openai.ChatCompletionResponse), nil
}

func TestMain(tests *testing.M) {
	openAIToken = faker.UUIDHyphenated()
	os.Setenv("OPENAI_TOKEN", openAIToken)

	result := tests.Run()

	os.Clearenv()
	os.Exit(result)
}

func testSetup() (*MockOpenAILibrary, string, string) {
	prompt := faker.Word()
	expectedResponse := faker.Word()
	expectedContext := mock.AnythingOfType("context.backgroundCtx")
	expectedRequest := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}
	openAIResponse := openai.ChatCompletionResponse{
		Choices: []openai.ChatCompletionChoice{
			{
				Message: openai.ChatCompletionMessage{Content: expectedResponse},
			},
		},
	}
	client := new(MockOpenAIClient)
	client.On("CreateChatCompletion", expectedContext, expectedRequest).Return(openAIResponse)
	mockOpenAI := new(MockOpenAILibrary)
	mockOpenAI.On("NewClient", openAIToken).Return(client)

	return mockOpenAI, prompt, expectedResponse
}

func Test_ProcessPromt_CreatesANewOpenAIChat_WithCredentialsFromAEnvVar(t *testing.T) {
	mockOpenAI, prompt, _ := testSetup()

	interactor := ChatGPTInteractor{openai: mockOpenAI}
	interactor.ProcessPrompt(prompt)

	mockOpenAI.AssertCalled(t, "NewClient", openAIToken)
}

func Test_ProcessPromt_ReturnsTheInteractionFromOpenAI(t *testing.T) {
	mockOpenAI, prompt, expectedResponse := testSetup()

	interactor := ChatGPTInteractor{openai: mockOpenAI}
	response := interactor.ProcessPrompt(prompt)

	assert.Equal(t, expectedResponse, response)
}
