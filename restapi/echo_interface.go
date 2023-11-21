package restapi

import (
	"github.com/labstack/echo/v4"
)

type Echo interface {
	Start(string) error
	POST(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route
}
