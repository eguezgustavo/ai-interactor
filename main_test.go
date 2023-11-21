package main

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRestAPI struct {
	mock.Mock
}

func (api *MockRestAPI) Start() {
	api.Called()
}

func TestMainFunctionShouldStartTheRestApi(t *testing.T) {
	api := new(MockRestAPI)
	api.On("Start").Return()

	StartApplication(api)

	api.AssertCalled(t, "Start")
}
