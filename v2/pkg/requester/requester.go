package requester

import (
	"net/http"
)

type Requester interface {
	CreateRequest(method, url string, body []byte) (*http.Request, error)
	Do(req *http.Request, headers map[string]string) (*http.Response, error)
}
