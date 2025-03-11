package nethttp

import (
	"bytes"
	"net/http"
	"testing"
)

func TestCreateRequest(t *testing.T) {
	requester := NewNetHTTPRequester()

	tests := []struct {
		method string
		url    string
		body   []byte
	}{
		{"GET", "http://example.com", nil},
		{"POST", "http://example.com", []byte("test body")},
	}

	for _, test := range tests {
		req, err := requester.CreateRequest(test.method, test.url, test.body)
		if err != nil {
			t.Errorf("CreateRequest returned an error: %v", err)
		}
		if req.Method != test.method {
			t.Errorf("Expected method %s, got %s", test.method, req.Method)
		}
		if req.URL.String() != test.url {
			t.Errorf("Expected URL %s, got %s", test.url, req.URL.String())
		}
		if test.body != nil {
			buf := new(bytes.Buffer)
			buf.ReadFrom(req.Body)
			if buf.String() != string(test.body) {
				t.Errorf("Expected body %s, got %s", string(test.body), buf.String())
			}
		}
	}
}
func TestDo(t *testing.T) {
	requester := NewNetHTTPRequester()

	tests := []struct {
		method  string
		url     string
		body    []byte
		headers map[string]string
	}{
		{"GET", "http://google.com", nil, map[string]string{"Content-Type": "application/json"}},
	}

	for _, test := range tests {
		req, err := requester.CreateRequest(test.method, test.url, test.body)
		if err != nil {
			t.Errorf("CreateRequest returned an error: %v", err)
		}

		resp, err := requester.Do(req, test.headers)
		if err != nil {
			t.Errorf("Do returned an error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		for key, value := range test.headers {
			if req.Header.Get(key) != value {
				t.Errorf("Expected header %s to be %s, got %s", key, value, req.Header.Get(key))
			}
		}
	}
}
