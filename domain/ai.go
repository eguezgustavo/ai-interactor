package interactor

type AI interface {
	ProcessPrompt(prompt string) string
}
