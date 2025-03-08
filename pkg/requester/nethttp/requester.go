package nethttp

import (
	"io"
	"net/http"
)

type NetHTTPRequester struct {
}

func NewNetHTTPRequester() *NetHTTPRequester {
	return &NetHTTPRequester{}
}

func (n *NetHTTPRequester) CreateRequest(method string, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}

func (n *NetHTTPRequester) Do(req *http.Request, headers map[string]string) (*http.Response, error) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	return client.Do(req)
}
