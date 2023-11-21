package restapi

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockEcho struct {
	mock.Mock
}

func (echo *MockEcho) Start(port string) error {
	echo.Called(port)
	return nil
}

func (echo *MockEcho) POST(endpoint string, handler echo.HandlerFunc, args ...echo.MiddlewareFunc) *echo.Route {
	echo.Called(endpoint, handler)
	return nil
}

type MockMyContextAI struct {
	mock.Mock
}

func (app *MockMyContextAI) Interact(prompt string) string {
	valuesToReturn := app.Called(prompt)
	return valuesToReturn.String(0)
}
