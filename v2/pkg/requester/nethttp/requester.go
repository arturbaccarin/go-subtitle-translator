package nethttp

import (
	"bytes"
	"net/http"
)

type NetHTTPRequester struct {
}

func NewNetHTTPRequester() *NetHTTPRequester {
	return &NetHTTPRequester{}
}

func (n *NetHTTPRequester) CreateRequest(method string, url string, body []byte) (*http.Request, error) {
	return http.NewRequest(method, url, bytes.NewReader(body))
}

func (n *NetHTTPRequester) Do(req *http.Request, headers map[string]string) (*http.Response, error) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	return client.Do(req)
}
