package deepl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arturbaccarin/go-subtitle-translator/pkg/requester"
	"github.com/arturbaccarin/go-subtitle-translator/pkg/translator/dto"
)

type APIClient struct {
	authKey     string
	apiHostname string
	requester   requester.Requester
}

func NewAPIClient(authKey string, apiHostname string, requester requester.Requester) *APIClient {
	return &APIClient{
		authKey:     authKey,
		apiHostname: apiHostname,
		requester:   requester,
	}
}

func (a *APIClient) Translate(requestPayload dto.Request) (*dto.Response, error) {
	reqBody, err := json.Marshal(requestPayload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %v", err)
	}

	url := fmt.Sprintf("%s/v2/translate", a.apiHostname)

	req, err := a.requester.CreateRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := a.requester.Do(req, a.buildHeaders())
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response dto.Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil
}

func (a *APIClient) buildHeaders() map[string]string {
	return map[string]string{
		"Authorization": "DeepL-Auth-Key " + a.authKey,
		"Content-Type":  "application/json",
	}
}
