package interactor

type Interactor struct {
	aiModel AI
}

func (app *Interactor) Interact(prompt string) string {
	return app.aiModel.ProcessPrompt(prompt)
}

func NewInteractor(model AI) Interactor {
	return Interactor{
		aiModel: model,
	}
}
