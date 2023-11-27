package restapi

import (
	"bytes"
	interactor "github.com/eguezgustavo/ai-interaction/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TestData struct {
	Port     int
	Endpoint string
}

func NewEchoMock(testData TestData) *MockEcho {
	echo := new(MockEcho)
	echo.On("POST", testData.Endpoint, mock.Anything).Return()
	echo.On("Start", fmt.Sprintf(":%d", testData.Port)).Return()
	return echo
}

func NewMyContextAI(expectedPrompt string, expectedInteraction string) interactor.InteractorInterface {
	app := new(MockMyContextAI)
	app.On("Interact", expectedPrompt).Return(expectedInteraction)
	return app
}

func NewTestData(endpoint string) TestData {
	return TestData{
		Port:     8000,
		Endpoint: endpoint,
	}
}

func NewContext(endpoint string, payload string) (echo.Context, httptest.ResponseRecorder) {
	request := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewReader([]byte(payload)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	return echo.New().NewContext(request, recorder), *recorder
}

func Interact(api RestAPI, context echo.Context, recorder httptest.ResponseRecorder) (int, string) {
	var response InteractResponse

	api.interactEndpoint(context)

	json.Unmarshal(recorder.Body.Bytes(), &response)

	return recorder.Code, response.Text
}

func Test_StartMethod_CreatesAndStartsAnEchoServer(t *testing.T) {
	testData := NewTestData("/interact")
	echo := NewEchoMock(testData)

	api := NewRestAPI(echo, new(MockMyContextAI))
	api.Start()

	firstCalledMethodName := echo.Calls[0].Method
	secondCallMethodName := echo.Calls[1].Method
	echo.AssertCalled(t, "POST", testData.Endpoint, mock.Anything)
	echo.AssertCalled(t, "Start", fmt.Sprintf(":%d", 8000))
	assert.Equal(t, "POST", firstCalledMethodName, "POST should have been called first")
	assert.Equal(t, "Start", secondCallMethodName, "Start should have been called after registering the endpoints")
}

func Test_InteractMethod_returnsDesiredInteraction(t *testing.T) {

	prompt := "some prompt"
	expectedInteraction := "some interaction"
	context, recorder := NewContext("/interact", fmt.Sprintf(`{"prompt": "%s"}`, prompt))
	app := NewMyContextAI(prompt, expectedInteraction)

	api := NewRestAPI(new(MockEcho), app)
	responseCode, response := Interact(api, context, recorder)

	assert.Equal(t, 200, responseCode)
	assert.Equal(t, expectedInteraction, response)
}
