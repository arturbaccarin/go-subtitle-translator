package requester

import (
	"io"
	"net/http"
)

type Requester interface {
	CreateRequest(method, url string, body io.Reader) (*http.Request, error)
	Do(req *http.Request, headers map[string]string) (*http.Response, error)
}
