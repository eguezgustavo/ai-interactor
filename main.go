package main

import (
	"github.com/labstack/echo/v4"

	"github.com/eguezgustavo/ai-interaction/adapters"
	"github.com/eguezgustavo/ai-interaction/domain"
	"github.com/eguezgustavo/ai-interaction/restapi"
)

type RestAPI interface {
	Start()
}

func main() {
	model := adapters.GetOpenAIGPTModel()
	app := interactor.NewInteractor(model)
	echoInstance := echo.New()
	api := restapi.NewRestAPI(echoInstance, &app)
	StartApplication(api)
}

func StartApplication(api RestAPI) {
	api.Start()
	// if err != nil {
	// 	fmt.Printf("ChatCompletion error: %v\n", err)
	// 	return
	// }

}
