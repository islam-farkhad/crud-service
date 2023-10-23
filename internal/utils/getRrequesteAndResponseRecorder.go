package utils

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

// GetRequestAndResponseRecorder creates new HTTP request and response recorder
func GetRequestAndResponseRecorder(method, route string, body []byte) (*http.Request, *httptest.ResponseRecorder) {
	req, err := http.NewRequest(method, route, bytes.NewBuffer(body))
	if err != nil {
		panic("could not create request")
	}
	return req, httptest.NewRecorder()
}
