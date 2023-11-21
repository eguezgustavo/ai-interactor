package interactor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockAI struct {
	mock.Mock
}

func (ai *mockAI) ProcessPrompt(prompt string) string {
	shouldReturnArguments := ai.Called(prompt)
	return shouldReturnArguments.Get(0).(string)
}

func Test_InteractMethod_UsesTheAIToInteract(t *testing.T) {
	prompt := "some prompt"
	expectedInteraction := "some interaction"
	ai := new(mockAI)
	ai.On("ProcessPrompt", prompt).Return(expectedInteraction)

	app := Interactor{
		aiModel: ai,
	}
	interaction := app.Interact(prompt)

	assert.Equal(t, expectedInteraction, interaction)
}
