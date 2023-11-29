package restapi

import (
	"net/http"

	"github.com/labstack/echo/v4"

	interactor "github.com/eguezgustavo/ai-interactor/domain"
)

type RestAPI struct {
	echo Echo
	app  interactor.InteractorInterface
}

type InteractPayload struct {
	Prompt string `json:"prompt"`
}

type InteractResponse struct {
	Text string `json:"text"`
}

func (api RestAPI) Start() {
	api.echo.POST("/interact", api.interactEndpoint)
	api.echo.Start(":8000")
}

func (api RestAPI) interactEndpoint(context echo.Context) error {
	var payload InteractPayload
	context.Bind(&payload)

	interaction := InteractResponse{Text: api.app.Interact(payload.Prompt)}

	return context.JSON(http.StatusOK, interaction)
}

func NewRestAPI(echo Echo, app interactor.InteractorInterface) RestAPI {
	return RestAPI{echo: echo, app: app}
}
