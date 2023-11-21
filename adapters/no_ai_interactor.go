package adapters

import "fmt"

type NoAIInteractor struct{}

func (app NoAIInteractor) ProcessPrompt(prompt string) string {
	return fmt.Sprintf("Interaction for prompt: %s", prompt)
}
