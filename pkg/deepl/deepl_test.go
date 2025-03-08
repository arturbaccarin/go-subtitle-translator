package deepl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/arturbaccarin/go-subtitle-translator/pkg/deepl/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRequester struct {
	mock.Mock
}

func (m *MockRequester) CreateRequest(method string, url string, body []byte) (*http.Request, error) {
	args := m.Called(method, url, body)
	return args.Get(0).(*http.Request), args.Error(1)
}

func (m *MockRequester) Do(req *http.Request, headers map[string]string) (*http.Response, error) {
	args := m.Called(req, headers)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestNewAPIClient(t *testing.T) {
	authKey := "test-auth-key"
	apiHostname := "https://api.deepl.com"
	mockRequester := new(MockRequester)

	client := NewAPIClient(authKey, apiHostname, mockRequester)

	assert.NotNil(t, client)
	assert.Equal(t, authKey, client.authKey)
	assert.Equal(t, apiHostname, client.apiHostname)
	assert.Equal(t, mockRequester, client.requester)
}

func TestTranslate(t *testing.T) {
	authKey := "test-auth-key"
	apiHostname := "https://api.deepl.com"
	mockRequester := new(MockRequester)

	requestPayload := dto.Request{
		Text:       []string{"Hello, World!"},
		SourceLang: "EN",
		TargetLang: "PT",
	}

	headers := map[string]string{
		"Authorization": "DeepL-Auth-Key test-auth-key",
		"Content-Type":  "application/json",
	}

	apiClient := NewAPIClient(authKey, apiHostname, mockRequester)

	t.Run("should return an error when create request fails", func(t *testing.T) {
		mockRequester.ExpectedCalls = nil

		mockRequester.On("CreateRequest", http.MethodPost, apiHostname, mock.Anything).Return(&http.Request{}, errors.New("some error"))

		resp, err := apiClient.Translate(requestPayload)

		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "error creating request: some error")
	})

	t.Run("should return an error when do fails", func(t *testing.T) {
		mockRequester.ExpectedCalls = nil

		mockRequester.On("CreateRequest", http.MethodPost, apiHostname, mock.Anything).Return(&http.Request{}, nil)
		mockRequester.On("Do", &http.Request{}, headers).Return(&http.Response{}, errors.New("some error"))

		resp, err := apiClient.Translate(requestPayload)

		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "error making request: some error")
	})

	t.Run("should return an error when status code is not 200", func(t *testing.T) {
		mockRequester.ExpectedCalls = nil

		httpResponse := http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(bytes.NewBufferString("response body")),
		}

		mockRequester.On("CreateRequest", http.MethodPost, apiHostname, mock.Anything).Return(&http.Request{}, nil)
		mockRequester.On("Do", &http.Request{}, headers).Return(&httpResponse, nil)

		resp, err := apiClient.Translate(requestPayload)

		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "unexpected status code: 400")
	})

	t.Run("should return an error when decoding response fails", func(t *testing.T) {
		mockRequester.ExpectedCalls = nil

		httpResponse := http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("response body")),
		}

		mockRequester.On("CreateRequest", http.MethodPost, apiHostname, mock.Anything).Return(&http.Request{}, nil)
		mockRequester.On("Do", &http.Request{}, headers).Return(&httpResponse, nil)

		resp, err := apiClient.Translate(requestPayload)

		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "error decoding response: invalid character 'r' looking for beginning of value")
	})

	t.Run("should return the response", func(t *testing.T) {
		mockRequester.ExpectedCalls = nil

		expectedResponse := dto.Response{
			Translations: []dto.TranslationResult{
				{
					DetectedSourceLanguage: "EN",
					Text:                   "Olá, Mundo!",
				},
			},
		}

		responseBytes, _ := json.Marshal(expectedResponse)

		httpResponse := http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(responseBytes)),
		}

		mockRequester.On("CreateRequest", http.MethodPost, apiHostname, mock.Anything).Return(&http.Request{}, nil)
		mockRequester.On("Do", &http.Request{}, headers).Return(&httpResponse, nil)

		resp, err := apiClient.Translate(requestPayload)

		assert.NotNil(t, resp)
		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, *resp)
	})
}
