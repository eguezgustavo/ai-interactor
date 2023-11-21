package interactor

type InteractorInterface interface {
	Interact(prompt string) string
}
